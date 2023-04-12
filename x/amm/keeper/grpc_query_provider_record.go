package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/amm/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ProviderRecord(c context.Context, req *types.QueryGetProviderRecordRequest) (*types.QueryGetProviderRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	recordKey := types.ProviderRecordKey(req.PoolName, req.Provider)
	val, found := k.GetProviderRecord(
		ctx,
		recordKey,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found ")
	}

	return &types.QueryGetProviderRecordResponse{lProviderRecord: val}, nil
}
