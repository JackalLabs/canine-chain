package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jackal-dao/canine/x/jklmining/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) MinerClaimsAll(c context.Context, req *types.QueryAllMinerClaimsRequest) (*types.QueryAllMinerClaimsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var minerClaimss []types.MinerClaims
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	minerClaimsStore := prefix.NewStore(store, types.KeyPrefix(types.MinerClaimsKeyPrefix))

	pageRes, err := query.Paginate(minerClaimsStore, req.Pagination, func(key []byte, value []byte) error {
		var minerClaims types.MinerClaims
		if err := k.cdc.Unmarshal(value, &minerClaims); err != nil {
			return err
		}

		minerClaimss = append(minerClaimss, minerClaims)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMinerClaimsResponse{MinerClaims: minerClaimss, Pagination: pageRes}, nil
}

func (k Keeper) MinerClaims(c context.Context, req *types.QueryGetMinerClaimsRequest) (*types.QueryGetMinerClaimsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetMinerClaims(
		ctx,
		req.Hash,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetMinerClaimsResponse{MinerClaims: val}, nil
}
