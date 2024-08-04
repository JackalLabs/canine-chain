package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AllProviders(c context.Context, req *types.QueryAllProviders) (*types.QueryAllProvidersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var providerss []types.Providers
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	providersStore := prefix.NewStore(store, types.KeyPrefix(types.ProvidersKeyPrefix))

	pageRes, err := query.Paginate(providersStore, req.Pagination, func(_ []byte, value []byte) error {
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

func (k Keeper) Provider(c context.Context, req *types.QueryProvider) (*types.QueryProviderResponse, error) {
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

	return &types.QueryProviderResponse{Provider: val}, nil
}

func (k Keeper) ActiveProviders(c context.Context, req *types.QueryActiveProviders) (*types.QueryActiveProvidersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	providers := k.GetAllActiveProviders(
		ctx,
	)

	return &types.QueryActiveProvidersResponse{Providers: providers}, nil
}
