package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jackalLabs/canine-chain/x/storage/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PayBlocksAll(c context.Context, req *types.QueryAllPayBlocksRequest) (*types.QueryAllPayBlocksResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var payBlockss []types.PayBlocks
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	payBlocksStore := prefix.NewStore(store, types.KeyPrefix(types.PayBlocksKeyPrefix))

	pageRes, err := query.Paginate(payBlocksStore, req.Pagination, func(key []byte, value []byte) error {
		var payBlocks types.PayBlocks
		if err := k.cdc.Unmarshal(value, &payBlocks); err != nil {
			return err
		}

		payBlockss = append(payBlockss, payBlocks)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPayBlocksResponse{PayBlocks: payBlockss, Pagination: pageRes}, nil
}

func (k Keeper) PayBlocks(c context.Context, req *types.QueryGetPayBlocksRequest) (*types.QueryGetPayBlocksResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetPayBlocks(
		ctx,
		req.Blockid,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPayBlocksResponse{PayBlocks: val}, nil
}
