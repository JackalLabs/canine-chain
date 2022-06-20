package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/jklmining/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CheckMinerIndex(goCtx context.Context, req *types.QueryCheckMinerIndexRequest) (*types.QueryCheckMinerIndexResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	k.GetMinedStarting(ctx)

	return &types.QueryCheckMinerIndexResponse{}, nil
}
