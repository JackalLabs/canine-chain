package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/jackalLabs/canine-chain/v4/x/rns/types"
)

func (k Keeper) RegisterName(ctx sdk.Context, sender string, nm string, data string, years int64) error {
	nm = strings.ToLower(nm)
	name, tld, err := GetNameAndTLD(nm)
	if err != nil {
		return err
	}

	if types.IsReserved[tld] {
		return types.ErrReserved
	}

	whois, isFound := k.GetNames(ctx, name, tld)
	// Set the price at which the name has to be bought if it didn't have an owner before

	cost, err := GetCostOfName(name, tld)
	if err != nil {
		return sdkerrors.Wrap(err, "failed to get cost")
	}

	price := sdk.Coins{sdk.NewInt64Coin("ujkl", cost*years)}

	blockHeight := ctx.BlockHeight()

	time := years * 5733818

	owner, err := sdk.AccAddressFromBech32(sender)
	if err != nil {
		return sdkerrors.Wrap(err, "cannot parse sender")
	}

	// If a name is found in store
	if isFound {
		if whois.Value == owner.String() {
			time = whois.Expires + time
		} else if blockHeight < whois.Expires {
			return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "name already registered")
		}
	} else {
		time += blockHeight
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, owner, types.ModuleName, price)
	if err != nil {
		return err
	}

	deposit, err := sdk.AccAddressFromBech32(k.GetParams(ctx).DepositAccount)
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, deposit, price)
	if err != nil {
		return err
	}

	emptySubdomains := []*types.Names{}

	// Create an updated whois record
	newWhois := types.Names{
		Name:       name,
		Expires:    time,
		Value:      owner.String(),
		Data:       data,
		Subdomains: emptySubdomains,
		Tld:        tld,
		Locked:     0,
	}
	// Write whois information to the store
	k.SetNames(ctx, newWhois)

	return nil
}

func (k msgServer) Register(goCtx context.Context, msg *types.MsgRegister) (*types.MsgRegisterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Try getting a name from the store

	err := k.RegisterName(ctx, msg.Creator, msg.Name, msg.Data, msg.Years)

	return &types.MsgRegisterResponse{}, err
}
