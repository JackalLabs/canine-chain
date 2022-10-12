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

func (k Keeper) ProvidersAll(c context.Context, req *types.QueryAllProvidersRequest) (*types.QueryAllProvidersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var providerss []types.Providers
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	providersStore := prefix.NewStore(store, types.KeyPrefix(types.ProvidersKeyPrefix))

	pageRes, err := query.Paginate(providersStore, req.Pagination, func(key []byte, value []byte) error {
		var providers types.Providers
		if err := k.cdc.Unmarshal(value, &providers); err != nil {
			return err
		}

		providerss = append(providerss, providers)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllProvidersResponse{Providers: providerss, Pagination: pageRes}, nil
}

func (k Keeper) Providers(c context.Context, req *types.QueryGetProvidersRequest) (*types.QueryGetProvidersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetProviders(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetProvidersResponse{Providers: val}, nil
}
