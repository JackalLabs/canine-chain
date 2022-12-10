package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/oracle/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Feed(c context.Context, req *types.QueryFeedRequest) (*types.QueryFeedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	ctx.Logger().Error(req.Name)

	feed, found := k.GetFeed(ctx, req.Name)
	ctx.Logger().Error(feed.Name)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryFeedResponse{Feed: feed}, nil
}

func (k Keeper) AllFeeds(c context.Context, req *types.QueryAllFeedsRequest) (*types.QueryAllFeedsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	f := k.GetAllFeeds(ctx)

	return &types.QueryAllFeedsResponse{Feed: f, Pagination: nil}, nil
}
