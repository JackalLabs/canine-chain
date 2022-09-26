package keeper

import (
	"context"

    "github.com/jackal-dao/canine/x/lp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListRecordsFromPool(goCtx context.Context,  req *types.QueryListRecordsFromPoolRequest) (*types.QueryListRecordsFromPoolResponse, error) {
	if req == nil {
        return nil, status.Error(codes.InvalidArgument, "invalid request")
    }

	ctx := sdk.UnwrapSDKContext(goCtx)

	records := k.GetAllLProviderRecord(ctx)

	return &types.QueryListRecordsFromPoolResponse{Records: records}, nil
}
