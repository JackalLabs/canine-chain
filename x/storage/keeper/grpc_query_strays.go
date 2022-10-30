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

func (k Keeper) StraysAll(c context.Context, req *types.QueryAllStraysRequest) (*types.QueryAllStraysResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var strayss []types.Strays
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	straysStore := prefix.NewStore(store, types.KeyPrefix(types.StraysKeyPrefix))

	pageRes, err := query.Paginate(straysStore, req.Pagination, func(key []byte, value []byte) error {
		var strays types.Strays
		if err := k.cdc.Unmarshal(value, &strays); err != nil {
			return err
		}

		strayss = append(strayss, strays)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStraysResponse{Strays: strayss, Pagination: pageRes}, nil
}

func (k Keeper) Strays(c context.Context, req *types.QueryGetStraysRequest) (*types.QueryGetStraysResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetStrays(
		ctx,
		req.Cid,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetStraysResponse{Strays: val}, nil
}
