package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/rns/types"
)

func (k Keeper) AddBid(ctx sdk.Context, sender string, name string, bid string) error {
	bidder, err := sdk.AccAddressFromBech32(sender)
	if err != nil {
		return err
	}

	price, err := sdk.ParseCoinsNormalized(bid)
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, bidder, types.ModuleName, price)
	if err != nil {
		return err
	}

	newBid := types.Bids{
		Index:  fmt.Sprintf("%s%s", bidder.String(), name),
		Name:   name,
		Bidder: bidder.String(),
		Price:  name,
	}
	k.SetBids(ctx, newBid)

	return nil
}

func (k msgServer) Bid(goCtx context.Context, msg *types.MsgBid) (*types.MsgBidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.AddBid(ctx, msg.Creator, msg.Name, msg.Bid)

	return &types.MsgBidResponse{}, err
}
