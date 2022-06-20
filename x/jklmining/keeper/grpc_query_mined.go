package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jackal-dao/canine/x/jklmining/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) MinedAll(c context.Context, req *types.QueryAllMinedRequest) (*types.QueryAllMinedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var mineds []types.Mined
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	minedStore := prefix.NewStore(store, types.KeyPrefix(types.MinedKey))

	pageRes, err := query.Paginate(minedStore, req.Pagination, func(key []byte, value []byte) error {
		var mined types.Mined
		if err := k.cdc.Unmarshal(value, &mined); err != nil {
			return err
		}

		mineds = append(mineds, mined)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMinedResponse{Mined: mineds, Pagination: pageRes}, nil
}

func (k Keeper) Mined(c context.Context, req *types.QueryGetMinedRequest) (*types.QueryGetMinedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	mined, found := k.GetMined(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetMinedResponse{Mined: mined}, nil
}
