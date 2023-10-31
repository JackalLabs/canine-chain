package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jackalLabs/canine-chain/v3/x/filetree/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// To remove
func (k Keeper) FilesAll(c context.Context, req *types.QueryAllFilesRequest) (*types.QueryAllFilesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var filess []types.Files
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	filesStore := prefix.NewStore(store, types.KeyPrefix(types.FilesKeyPrefix))

	pageRes, err := query.Paginate(filesStore, req.Pagination, func(key []byte, value []byte) error {
		var files types.Files
		if err := k.cdc.Unmarshal(value, &files); err != nil {
			return err
		}

		filess = append(filess, files)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFilesResponse{Files: filess, Pagination: pageRes}, nil
}

func (k Keeper) FilesAllByOwner(c context.Context, req *types.QueryAllFilesByOwnerRequest) (*types.QueryAllFilesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var filess []types.Files
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	filesStore := prefix.NewStore(store, types.KeyPrefix(types.FilesKeyPrefix))

	pageRes, err := query.Paginate(filesStore, req.Pagination, func(key []byte, value []byte) error {
		var files types.Files
		if err := k.cdc.Unmarshal(value, &files); err != nil {
			return err
		}

		if files.Owner == req.Owner {
			filess = append(filess, files)
		}

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFilesResponse{Files: filess, Pagination: pageRes}, nil
}

func (k Keeper) Files(c context.Context, req *types.QueryFileRequest) (*types.QueryFileResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetFiles(
		ctx,
		req.Address,
		req.OwnerAddress)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryFileResponse{Files: val}, nil
}
