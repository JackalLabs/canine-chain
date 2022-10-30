package keeper

import (
	"context"
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/rns/types"
)

func (k msgServer) Buy(goCtx context.Context, msg *types.MsgBuy) (*types.MsgBuyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	buyer, _ := sdk.AccAddressFromBech32(msg.Creator)

	sale, found := k.GetForsale(ctx, msg.Name)

	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Name not for sale.")
	}

	n, tld, err := getNameAndTLD(msg.Name)
	if err != nil {
		return nil, err
	}
	name, nfound := k.GetNames(ctx, n, tld)

	if !nfound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Name does not exist or has expired.")
	}

	blockHeight := ctx.BlockHeight()

	if blockHeight > name.Expires {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Name does not exist or has expired.")
	}

	if name.Value == msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "You cannot buy your own name.")
	}

	seller, _ := sdk.AccAddressFromBech32(sale.Owner)

	price, ok := sdk.NewIntFromString(sale.Price)

	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Price is not a valid number.")
	}

	coin := sdk.NewCoin("ujkl", price)
	coins := sdk.NewCoins(coin)

	ctx.Logger().Error(fmt.Sprintf("%s %s", "coins available: ", k.bankKeeper.SpendableCoins(ctx, buyer).String()))

	err = k.bankKeeper.SendCoins(ctx, buyer, seller, coins)
	if err != nil {
		return nil, err
	}

	k.RemoveForsale(ctx, sale.Name)
	name.Value = msg.Creator
	name.Data = "{}"
	k.SetNames(ctx, name)

	return &types.MsgBuyResponse{}, nil
}
