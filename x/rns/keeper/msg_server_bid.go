package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/rns/types"
)

func (k msgServer) Bid(goCtx context.Context, msg *types.MsgBid) (*types.MsgBidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	bidder, _ := sdk.AccAddressFromBech32(msg.Creator)

	cost, _ := sdk.NewIntFromString(msg.Bid)
	price := sdk.Coins{sdk.NewInt64Coin("ujkl", cost.Int64())}
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, bidder, types.ModuleName, price)
	if err != nil {
		return nil, err
	}

	newBid := types.Bids{
		Index:  bidder.String() + msg.Name,
		Name:   msg.Name,
		Bidder: bidder.String(),
		Price:  msg.Bid,
	}
	k.SetBids(ctx, newBid)

	return &types.MsgBidResponse{}, nil
}
