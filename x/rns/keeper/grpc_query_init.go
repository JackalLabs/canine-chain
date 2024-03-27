package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jackalLabs/canine-chain/v3/x/rns/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AllInits(c context.Context, req *types.QueryAllInits) (*types.QueryAllInitsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var inits []types.Init
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	initStore := prefix.NewStore(store, types.KeyPrefix(types.InitKeyPrefix))

	pageRes, err := query.Paginate(initStore, req.Pagination, func(_ []byte, value []byte) error {
		var init types.Init
		if err := k.cdc.Unmarshal(value, &init); err != nil {
			return err
		}

		inits = append(inits, init)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllInitsResponse{Init: inits, Pagination: pageRes}, nil
}

func (k Keeper) Init(c context.Context, req *types.QueryInit) (*types.QueryInitResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	_, found := k.GetInit(
		ctx,
		req.Address,
	)

	return &types.QueryInitResponse{Init: found}, nil
}
