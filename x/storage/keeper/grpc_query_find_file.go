package keeper

import (
	"context"
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/storage/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListFiles(ctx sdk.Context, fid string) []string {
	allDeals := k.GetAllActiveDeals(ctx)

	var providers []string

	for i := 0; i < len(allDeals); i++ {
		if allDeals[i].Fid == fid {
			provider, ok := k.GetProviders(ctx, allDeals[i].Provider)
			if !ok {
				continue
			}

			providers = append(providers, provider.Ip)
		}
	}

	return providers
}

func (k Keeper) FindFile(goCtx context.Context, req *types.QueryFindFileRequest) (*types.QueryFindFileResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	ls := k.ListFiles(ctx, req.Fid)

	ProviderIps, err := json.Marshal(ls)
	if err != nil {
		return nil, err
	}

	return &types.QueryFindFileResponse{ProviderIps: string(ProviderIps)}, nil
}
