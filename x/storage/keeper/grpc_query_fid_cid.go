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

func (k Keeper) FidCidAll(c context.Context, req *types.QueryAllFidCidRequest) (*types.QueryAllFidCidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var fidCids []types.FidCid
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	fidCidStore := prefix.NewStore(store, types.KeyPrefix(types.FidCidKeyPrefix))

	pageRes, err := query.Paginate(fidCidStore, req.Pagination, func(key []byte, value []byte) error {
		var fidCid types.FidCid
		if err := k.cdc.Unmarshal(value, &fidCid); err != nil {
			return err
		}

		fidCids = append(fidCids, fidCid)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFidCidResponse{FidCid: fidCids, Pagination: pageRes}, nil
}

func (k Keeper) FidCid(c context.Context, req *types.QueryGetFidCidRequest) (*types.QueryGetFidCidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetFidCid(
		ctx,
		req.Fid,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetFidCidResponse{FidCid: val}, nil
}
