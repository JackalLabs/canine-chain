package keeper

import (
	"context"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/x/rns/types"
)

func (k Keeper) AddBid(ctx sdk.Context, sender string, name string, bid string) error {
	name = strings.ToLower(name)

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
		Price:  bid,
	}
	k.SetBids(ctx, newBid)

	return nil
}

func (k msgServer) Bid(goCtx context.Context, msg *types.MsgBid) (*types.MsgBidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.AddBid(ctx, msg.Creator, msg.Name, msg.Bid.String())

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
			types.EventSetBid,
			sdk.NewAttribute(types.AttributeKeySigner, msg.Creator),
		),
	)

	return &types.MsgBidResponse{}, err
}
