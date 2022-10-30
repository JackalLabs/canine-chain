package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jackal-dao/canine/x/lp/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) LPoolAll(c context.Context, req *types.QueryAllLPoolRequest) (*types.QueryAllLPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var lPools []types.LPool
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	lPoolStore := prefix.NewStore(store, types.KeyPrefix(types.LPoolKeyPrefix))

	pageRes, err := query.Paginate(lPoolStore, req.Pagination, func(key []byte, value []byte) error {
		var lPool types.LPool
		if err := k.cdc.Unmarshal(value, &lPool); err != nil {
			return err
		}

		lPools = append(lPools, lPool)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLPoolResponse{LPool: lPools, Pagination: pageRes}, nil
}

func (k Keeper) LPool(c context.Context, req *types.QueryGetLPoolRequest) (*types.QueryGetLPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetLPool(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetLPoolResponse{LPool: val}, nil
}
