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

func (k Keeper) StoragePaymentInfoAll(c context.Context, req *types.QueryAllStoragePaymentInfoRequest) (*types.QueryAllStoragePaymentInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var storagePaymentInfos []types.StoragePaymentInfo
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	storagePaymentStore := prefix.NewStore(store, types.KeyPrefix(types.StoragePaymentInfoKeyPrefix))

	pageRes, err := query.Paginate(storagePaymentStore, req.Pagination, func(key []byte, value []byte) error {
		var storagePaymentInfo types.StoragePaymentInfo
		if err := k.cdc.Unmarshal(value, &storagePaymentInfo); err != nil {
			return err
		}

		storagePaymentInfos = append(storagePaymentInfos, storagePaymentInfo)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStoragePaymentInfoResponse{StoragePaymentInfo: storagePaymentInfos, Pagination: pageRes}, nil
}

func (k Keeper) StoragePaymentInfo(c context.Context, req *types.QueryStoragePaymentInfoRequest) (*types.QueryStoragePaymentInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetStoragePaymentInfo(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryStoragePaymentInfoResponse{StoragePaymentInfo: val}, nil
}
