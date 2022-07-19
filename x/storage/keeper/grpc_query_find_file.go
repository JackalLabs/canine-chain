package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/storage/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) FindFile(goCtx context.Context, req *types.QueryFindFileRequest) (*types.QueryFindFileResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	allDeals := k.GetAllActiveDeals(ctx)

	var miners []string

	for i := 0; i < len(allDeals); i++ {
		if allDeals[i].Fid == req.Fid {
			miner, ok := k.GetMiners(ctx, allDeals[i].Miner)
			if !ok {
				continue
			}

			miners = append(miners, miner.Ip)
		}
	}

	return &types.QueryFindFileResponse{MinerIps: fmt.Sprintf("%v", miners)}, nil
}
