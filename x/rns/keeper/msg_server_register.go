package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/jackalLabs/canine-chain/x/rns/types"
)

func (k Keeper) RegisterName(ctx sdk.Context, sender string, nm string, data string, years string) error {
	name, tld, err := getNameAndTLD(nm)
	if err != nil {
		return err
	}

	if types.IsReserved[tld] {
		return types.ErrReserved
	}

	whois, isFound := k.GetNames(ctx, name, tld)
	// Set the price at which the name has to be bought if it didn't have an owner before

	chars := strings.Count(name, "")

	baseCost := getCost(tld)

	var cost int64

	switch chars {
	case 0:
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Must be 1 or more characters.")
	case 1:
		cost = baseCost * 32
	case 2:
		cost = baseCost * 16
	case 3:
		cost = baseCost * 8
	case 4:
		cost = baseCost * 4
	case 5:
		cost = baseCost * 2
	default:
		cost = baseCost
	}

	price := sdk.Coins{sdk.NewInt64Coin("ujkl", cost)}

	numYears, _ := sdk.NewIntFromString(years)

	blockHeight := ctx.BlockHeight()

	time := numYears.Int64() * 6311520

	owner, _ := sdk.AccAddressFromBech32(sender)
	// If a name is found in store
	if isFound {
		if whois.Value == owner.String() {
			time = whois.Expires + time
		} else if blockHeight < whois.Expires {
			return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Name already registered")
		}
	} else {
		time += blockHeight
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, owner, types.ModuleName, price)
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
