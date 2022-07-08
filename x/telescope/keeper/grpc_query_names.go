package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/themarstonconnell/telescope/x/telescope/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NamesAll(c context.Context, req *types.QueryAllNamesRequest) (*types.QueryAllNamesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var namess []types.Names
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	namesStore := prefix.NewStore(store, types.KeyPrefix(types.NamesKeyPrefix))

	pageRes, err := query.Paginate(namesStore, req.Pagination, func(key []byte, value []byte) error {
		var names types.Names
		if err := k.cdc.Unmarshal(value, &names); err != nil {
			return err
		}

		namess = append(namess, names)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNamesResponse{Names: namess, Pagination: pageRes}, nil
}

func (k Keeper) Names(c context.Context, req *types.QueryGetNamesRequest) (*types.QueryGetNamesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNames(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNamesResponse{Names: val}, nil
}
