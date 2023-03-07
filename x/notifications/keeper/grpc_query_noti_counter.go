package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jackal-dao/canine/x/notifications/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NotiCounterAll(c context.Context, req *types.QueryAllNotiCounterRequest) (*types.QueryAllNotiCounterResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var notiCounters []types.NotiCounter
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	notiCounterStore := prefix.NewStore(store, types.KeyPrefix(types.NotiCounterKeyPrefix))

	pageRes, err := query.Paginate(notiCounterStore, req.Pagination, func(key []byte, value []byte) error {
		var notiCounter types.NotiCounter
		if err := k.cdc.Unmarshal(value, &notiCounter); err != nil {
			return err
		}

		notiCounters = append(notiCounters, notiCounter)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNotiCounterResponse{NotiCounter: notiCounters, Pagination: pageRes}, nil
}

func (k Keeper) NotiCounter(c context.Context, req *types.QueryGetNotiCounterRequest) (*types.QueryGetNotiCounterResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNotiCounter(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNotiCounterResponse{NotiCounter: val}, nil
}
