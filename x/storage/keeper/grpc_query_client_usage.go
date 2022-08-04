package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jackal-dao/canine/x/storage/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ClientUsageAll(c context.Context, req *types.QueryAllClientUsageRequest) (*types.QueryAllClientUsageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var clientUsages []types.ClientUsage
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	clientUsageStore := prefix.NewStore(store, types.KeyPrefix(types.ClientUsageKeyPrefix))

	pageRes, err := query.Paginate(clientUsageStore, req.Pagination, func(key []byte, value []byte) error {
		var clientUsage types.ClientUsage
		if err := k.cdc.Unmarshal(value, &clientUsage); err != nil {
			return err
		}

		clientUsages = append(clientUsages, clientUsage)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllClientUsageResponse{ClientUsage: clientUsages, Pagination: pageRes}, nil
}

func (k Keeper) ClientUsage(c context.Context, req *types.QueryGetClientUsageRequest) (*types.QueryGetClientUsageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetClientUsage(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetClientUsageResponse{ClientUsage: val}, nil
}
