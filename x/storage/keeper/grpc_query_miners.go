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

func (k Keeper) MinersAll(c context.Context, req *types.QueryAllMinersRequest) (*types.QueryAllMinersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var minerss []types.Miners
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	minersStore := prefix.NewStore(store, types.KeyPrefix(types.MinersKeyPrefix))

	pageRes, err := query.Paginate(minersStore, req.Pagination, func(key []byte, value []byte) error {
		var miners types.Miners
		if err := k.cdc.Unmarshal(value, &miners); err != nil {
			return err
		}

		minerss = append(minerss, miners)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMinersResponse{Miners: minerss, Pagination: pageRes}, nil
}

func (k Keeper) Miners(c context.Context, req *types.QueryGetMinersRequest) (*types.QueryGetMinersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetMiners(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetMinersResponse{Miners: val}, nil
}
