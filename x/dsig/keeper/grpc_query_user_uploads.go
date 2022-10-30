package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jackalLabs/canine-chain/x/dsig/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) UserUploadsAll(c context.Context, req *types.QueryAllUserUploadsRequest) (*types.QueryAllUserUploadsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var userUploadss []types.UserUploads
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	userUploadsStore := prefix.NewStore(store, types.KeyPrefix(types.UserUploadsKeyPrefix))

	pageRes, err := query.Paginate(userUploadsStore, req.Pagination, func(key []byte, value []byte) error {
		var userUploads types.UserUploads
		if err := k.cdc.Unmarshal(value, &userUploads); err != nil {
			return err
		}

		userUploadss = append(userUploadss, userUploads)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUserUploadsResponse{UserUploads: userUploadss, Pagination: pageRes}, nil
}

func (k Keeper) UserUploads(c context.Context, req *types.QueryGetUserUploadsRequest) (*types.QueryGetUserUploadsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetUserUploads(
		ctx,
		req.Fid,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetUserUploadsResponse{UserUploads: val}, nil
}
