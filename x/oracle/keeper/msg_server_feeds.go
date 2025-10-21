package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v5/x/oracle/types"
)

func (k msgServer) CreateFeed(goCtx context.Context, msg *types.MsgCreateFeed) (*types.MsgCreateFeedResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, found := k.GetFeed(ctx, msg.Name)
	if found {
		return nil, fmt.Errorf("cannot overwrite feed, name exists")
	}

	feed := types.Feed{
		Owner:      msg.Creator,
		Data:       "",
		LastUpdate: ctx.BlockTime(),
		Name:       msg.Name,
	}

	add, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	c := sdk.NewInt64Coin("ujkl", 100*1000000)
	cs := sdk.NewCoins(c)

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, add, types.ModuleName, cs)
	if err != nil {
		return nil, err
	}

	depo, err := sdk.AccAddressFromBech32(k.GetParams(ctx).Deposit)
	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, depo, cs)
	if err != nil {
		return nil, err
	}

	k.SetFeed(ctx, feed)

	return &types.MsgCreateFeedResponse{}, nil
}

func (k msgServer) UpdateFeed(goCtx context.Context, msg *types.MsgUpdateFeed) (*types.MsgUpdateFeedResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	feed, found := k.GetFeed(ctx, msg.Name)
	if !found {
		return nil, fmt.Errorf("cannot find feed")
	}

	if feed.Owner != msg.Creator {
		return nil, fmt.Errorf("you do not own this feed")
	}

	feed.Data = msg.Data
	feed.LastUpdate = ctx.BlockTime()

	k.SetFeed(ctx, feed)

	return &types.MsgUpdateFeedResponse{}, nil
}
