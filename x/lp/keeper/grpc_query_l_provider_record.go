package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/lp/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) LProviderRecord(c context.Context, req *types.QueryGetLProviderRecordRequest) (*types.QueryGetLProviderRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	recordKey := types.LProviderRecordKey(req.PoolName, req.Provider)
	val, found := k.GetLProviderRecord(
		ctx,
		recordKey,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found ")
	}

	return &types.QueryGetLProviderRecordResponse{LProviderRecord: val}, nil
}
