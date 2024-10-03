package keeper

import (
	"context"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/x/rns/types"
)

func (k Keeper) BuyName(ctx sdk.Context, sender string, nm string) error {
	nm = strings.ToLower(nm)

	buyer, err := sdk.AccAddressFromBech32(sender)
	if err != nil {
		return err
	}

	sale, found := k.GetForsale(ctx, nm)

	if !found {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Name not for sale.")
	}

	n, tld, err := GetNameAndTLD(nm)
	if err != nil {
		return err
	}
	name, nfound := k.GetNames(ctx, n, tld)

	if !nfound {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Name does not exist or has expired.")
	}

	blockHeight := ctx.BlockHeight()

	if blockHeight > name.Expires {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "Name does not exist or has expired.")
	}

	if name.Value == sender {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "You cannot buy your own name.")
	}

	seller, _ := sdk.AccAddressFromBech32(sale.Owner)

	price, err := sdk.ParseCoinNormalized(sale.Price)
	if err != nil {
		return sdkerrors.Wrap(err, "Price is not a valid coin.")
	}

	coins := sdk.NewCoins(price)

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, buyer, types.ModuleName, coins)
	if err != nil {
		return err
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, seller, coins)
	if err != nil {
		return err
	}

	k.RemoveForsale(ctx, sale.Name)
	name.Value = sender
	name.Data = "{}"
	k.SetNames(ctx, name)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventBuyName,
			sdk.NewAttribute(types.AttributeName, nm),
			sdk.NewAttribute(types.AttributeOwner, sender),
		),
	)

	return nil
}

func (k msgServer) Buy(goCtx context.Context, msg *types.MsgBuy) (*types.MsgBuyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.BuyName(ctx, msg.Creator, msg.Name)

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

	return &types.MsgBuyResponse{}, err
}
