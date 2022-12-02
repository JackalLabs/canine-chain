package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jackalLabs/canine-chain/x/rns/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ForsaleAll(c context.Context, req *types.QueryAllForsalesRequest) (*types.QueryAllForsalesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var forsales []types.Forsale
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	forsaleStore := prefix.NewStore(store, types.KeyPrefix(types.ForsaleKeyPrefix))

	pageRes, err := query.Paginate(forsaleStore, req.Pagination, func(key []byte, value []byte) error {
		var forsale types.Forsale
		if err := k.cdc.Unmarshal(value, &forsale); err != nil {
			return err
		}

		forsales = append(forsales, forsale)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllForsalesResponse{Forsale: forsales, Pagination: pageRes}, nil
}

func (k Keeper) Forsale(c context.Context, req *types.QueryForsaleRequest) (*types.QueryForsaleResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetForsale(
		ctx,
		req.Name,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryForsaleResponse{Forsale: val}, nil
}
