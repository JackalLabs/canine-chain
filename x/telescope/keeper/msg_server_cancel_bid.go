package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackal-dao/canine/x/telescope/types"
)

func (k msgServer) CancelBid(goCtx context.Context, msg *types.MsgCancelBid) (*types.MsgCancelBidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	bidder, _ := sdk.AccAddressFromBech32(msg.Creator)

	bid, bidFound := k.GetBids(ctx, msg.Creator+msg.Name)

	if bidFound {

		cost, _ := sdk.NewIntFromString(bid.Price)
		price := sdk.Coins{sdk.NewInt64Coin("ujkl", cost.Int64())}
		k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, bidder, price)
		k.RemoveBids(ctx, msg.Creator+msg.Name)
	} else {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Bid does not exist or has expired.")
	}

	return &types.MsgCancelBidResponse{}, nil
}
