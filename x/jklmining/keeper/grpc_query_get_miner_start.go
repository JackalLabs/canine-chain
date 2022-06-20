package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/jklmining/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetMinerStart(goCtx context.Context, req *types.QueryGetMinerStartRequest) (*types.QueryGetMinerStartResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	res := k.GetMinedStarting(ctx)

	return &types.QueryGetMinerStartResponse{Index: fmt.Sprintf("%d", res)}, nil
}
