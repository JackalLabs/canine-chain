package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jackal-dao/canine/x/rns/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) BidsAll(c context.Context, req *types.QueryAllBidsRequest) (*types.QueryAllBidsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var bidss []types.Bids
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	bidsStore := prefix.NewStore(store, types.KeyPrefix(types.BidsKeyPrefix))

	pageRes, err := query.Paginate(bidsStore, req.Pagination, func(key []byte, value []byte) error {
		var bids types.Bids
		if err := k.cdc.Unmarshal(value, &bids); err != nil {
			return err
		}

		bidss = append(bidss, bids)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllBidsResponse{Bids: bidss, Pagination: pageRes}, nil
}

func (k Keeper) Bids(c context.Context, req *types.QueryGetBidsRequest) (*types.QueryGetBidsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetBids(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetBidsResponse{Bids: val}, nil
}
