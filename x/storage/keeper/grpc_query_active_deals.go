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

func (k Keeper) ActiveDealsAll(c context.Context, req *types.QueryAllActiveDealsRequest) (*types.QueryAllActiveDealsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var activeDealss []types.ActiveDeals
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	activeDealsStore := prefix.NewStore(store, types.KeyPrefix(types.ActiveDealsKeyPrefix))

	pageRes, err := query.Paginate(activeDealsStore, req.Pagination, func(key []byte, value []byte) error {
		var activeDeals types.ActiveDeals
		if err := k.cdc.Unmarshal(value, &activeDeals); err != nil {
			return err
		}

		activeDealss = append(activeDealss, activeDeals)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllActiveDealsResponse{ActiveDeals: activeDealss, Pagination: pageRes}, nil
}

func (k Keeper) ActiveDeals(c context.Context, req *types.QueryGetActiveDealsRequest) (*types.QueryGetActiveDealsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetActiveDeals(
		ctx,
		req.Cid,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetActiveDealsResponse{ActiveDeals: val}, nil
}
