package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/amm/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListRecordsFromPool(goCtx context.Context, req *types.QueryListRecordsFromPoolRequest) (*types.QueryListRecordsFromPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	records := k.GetAllRecordOfPool(ctx, req.PoolId)

	return &types.QueryListRecordsFromPoolResponse{Records: records}, nil
}
