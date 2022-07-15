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

func (k Keeper) ProofsAll(c context.Context, req *types.QueryAllProofsRequest) (*types.QueryAllProofsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var proofss []types.Proofs
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	proofsStore := prefix.NewStore(store, types.KeyPrefix(types.ProofsKeyPrefix))

	pageRes, err := query.Paginate(proofsStore, req.Pagination, func(key []byte, value []byte) error {
		var proofs types.Proofs
		if err := k.cdc.Unmarshal(value, &proofs); err != nil {
			return err
		}

		proofss = append(proofss, proofs)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllProofsResponse{Proofs: proofss, Pagination: pageRes}, nil
}

func (k Keeper) Proofs(c context.Context, req *types.QueryGetProofsRequest) (*types.QueryGetProofsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetProofs(
		ctx,
		req.Cid,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetProofsResponse{Proofs: val}, nil
}
