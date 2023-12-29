package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/query"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/oracle/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Feed(c context.Context, req *types.QueryFeed) (*types.QueryFeedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	ctx.Logger().Debug(req.Name)

	feed, found := k.GetFeed(ctx, req.Name)
	ctx.Logger().Debug(feed.Name)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryFeedResponse{Feed: feed}, nil
}

func (k Keeper) AllFeeds(c context.Context, req *types.QueryAllFeeds) (*types.QueryAllFeedsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	feeds := make([]types.Feed, 0)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeedKeyPrefix))
	pg, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var feed types.Feed
		if err := k.cdc.Unmarshal(value, &feed); err != nil {
			return err
		}

		feeds = append(feeds, feed)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFeedsResponse{Feed: feeds, Pagination: pg}, nil
}
