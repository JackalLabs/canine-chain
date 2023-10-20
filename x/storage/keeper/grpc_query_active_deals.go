package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) FilesAll(c context.Context, req *types.QueryAllFilesRequest) (*types.QueryAllFilesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var activeDealss []types.UnifiedFile
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	activeDealsStore := prefix.NewStore(store, types.KeyPrefix(types.FilePrimaryKeyPrefix))

	pageRes, err := query.Paginate(activeDealsStore, req.Pagination, func(key []byte, value []byte) error {
		var activeDeals types.UnifiedFile
		if err := k.cdc.Unmarshal(value, &activeDeals); err != nil {
			return err
		}

		activeDealss = append(activeDealss, activeDeals)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFilesResponse{Files: activeDealss, Pagination: pageRes}, nil
}

func (k Keeper) File(c context.Context, req *types.QueryFileRequest) (*types.QueryFileResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetFile(
		ctx,
		req.Merkle,
		req.Owner,
		req.Start,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryFileResponse{File: val}, nil
}
