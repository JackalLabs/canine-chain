package keeper

import (
	"context"
	"fmt"

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

	feed, found := k.GetFeed(ctx, req.Name)
	if !found {
		return nil, fmt.Errorf("cannot find feed")
	}

	return &types.QueryFeedResponse{Feed: feed}, nil
}

func (k Keeper) AllFeeds(c context.Context, req *types.QueryAllFeedsRequest) (*types.QueryAllFeedsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryAllFeedsResponse{Feed: k.GetAllFeeds(ctx)}, nil
}
