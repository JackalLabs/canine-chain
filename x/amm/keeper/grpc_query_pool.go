package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jackalLabs/canine-chain/x/amm/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PoolAll(c context.Context, req *types.QueryAllPoolRequest) (*types.QueryAllPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var lPools []types.Pool
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	lPoolStore := prefix.NewStore(store, types.KeyPrefix(types.PoolKeyPrefix))

	pageRes, err := query.Paginate(lPoolStore, req.Pagination, func(key []byte, value []byte) error {
		var lPool types.Pool
		if err := k.cdc.Unmarshal(value, &lPool); err != nil {
			return err
		}

		lPools = append(lPools, lPool)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPoolResponse{Pool: lPools, Pagination: pageRes}, nil
}

func (k Keeper) Pool(c context.Context, req *types.QueryGetPoolRequest) (*types.QueryGetPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetPool(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPoolResponse{Pool: val}, nil
}
