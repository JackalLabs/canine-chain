package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/storage/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Freespace(goCtx context.Context, req *types.QueryFreespaceRequest) (*types.QueryFreespaceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	num := k.GetMinerUsing(ctx, req.Address)

	miner, found := k.GetMiners(ctx, req.Address)

	if !found {
		return nil, fmt.Errorf("can't find miner")
	}

	space, ok := sdk.NewIntFromString(miner.Totalspace)

	if !ok {
		return nil, fmt.Errorf("can't parse total space")
	}

	return &types.QueryFreespaceResponse{
		Space: fmt.Sprintf("%d", space.Int64()-num),
	}, nil
}
