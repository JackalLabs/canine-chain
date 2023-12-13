package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListFileLocations(ctx sdk.Context, merkle []byte) []string {
	allDeals := k.GetAllFilesWithMerkle(ctx, merkle)

	providers := make([]string, 0)

	for _, deal := range allDeals {
		proofs := deal.GetProofs()
		for _, proof := range proofs {
			ss := strings.Split(proof, "/")
			prov := ss[0]

			provider, ok := k.GetProviders(ctx, prov)
			if !ok {
				continue
			}

			providers = append(providers, provider.Ip)
		}

	}

	return providers
}

func (k Keeper) FindFile(goCtx context.Context, req *types.QueryFindFile) (*types.QueryFindFileResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	ls := k.ListFileLocations(ctx, req.Merkle)

	return &types.QueryFindFileResponse{ProviderIps: ls}, nil
}
