package keeper

import (
	"context"
	"fmt"
	"strings"

	allTypes "github.com/jackalLabs/canine-chain/v4/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/jackalLabs/canine-chain/v4/x/rns/types"
)

func (k Keeper) RegisterRNSName(ctx sdk.Context, sender string, nm string, data string, years int64, primary bool) error {
	nm = strings.ToLower(nm)
	nm = strings.ReplaceAll(nm, " ", "")
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

	time := years * 5484530

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

	deposit, err := allTypes.GetPOLAccount()
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

	_, hasPrimary := k.GetPrimaryName(ctx, newWhois.Value)

	if primary || !hasPrimary {
		k.SetPrimaryName(ctx, newWhois.Value, newWhois.Name, newWhois.Tld)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventRegister,
			sdk.NewAttribute(types.AttributeName, fmt.Sprintf("%s.%s", newWhois.Name, newWhois.Tld)),
			sdk.NewAttribute(types.AttributeOwner, sender),
		),
	)

	return nil
}

func (k msgServer) MakePrimary(goCtx context.Context, msg *types.MsgMakePrimary) (*types.MsgMakePrimaryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	nm := strings.ToLower(msg.Name)
	name, tld, err := GetNameAndTLD(nm)
	if err != nil {
		return nil, err
	}

	k.SetPrimaryName(ctx, msg.Creator, name, tld)

	return &types.MsgMakePrimaryResponse{}, err
}

// Register is Deprecated! Use RegisterName instead.
func (k msgServer) Register(goCtx context.Context, msg *types.MsgRegister) (*types.MsgRegisterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Try getting a name from the store

	err := k.RegisterRNSName(ctx, msg.Creator, msg.Name, msg.Data, msg.Years, false)

	return &types.MsgRegisterResponse{}, err
}

func (k msgServer) RegisterName(goCtx context.Context, msg *types.MsgRegisterName) (*types.MsgRegisterNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.RegisterRNSName(ctx, msg.Creator, msg.Name, msg.Data, msg.Years, msg.SetPrimary)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeJackalMessage,
			sdk.NewAttribute(types.AttributeKeySigner, msg.Creator),
		),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventRegister,
			sdk.NewAttribute(types.AttributeKeySigner, msg.Creator),
		),
	)

	return &types.MsgRegisterNameResponse{}, err
}
