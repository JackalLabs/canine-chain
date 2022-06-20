package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jackal-dao/canine/x/jklmining/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SaveRequestsAll(c context.Context, req *types.QueryAllSaveRequestsRequest) (*types.QueryAllSaveRequestsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var saveRequestss []types.SaveRequests
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	saveRequestsStore := prefix.NewStore(store, types.KeyPrefix(types.SaveRequestsKeyPrefix))

	pageRes, err := query.Paginate(saveRequestsStore, req.Pagination, func(key []byte, value []byte) error {
		var saveRequests types.SaveRequests
		if err := k.cdc.Unmarshal(value, &saveRequests); err != nil {
			return err
		}

		saveRequestss = append(saveRequestss, saveRequests)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSaveRequestsResponse{SaveRequests: saveRequestss, Pagination: pageRes}, nil
}

func (k Keeper) SaveRequests(c context.Context, req *types.QueryGetSaveRequestsRequest) (*types.QueryGetSaveRequestsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetSaveRequests(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetSaveRequestsResponse{SaveRequests: val}, nil
}
