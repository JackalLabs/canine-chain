package keeper

import (
	"context"
	"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
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

	if req.Merkle == nil {
		return nil, status.Error(codes.InvalidArgument, "no merkle hash provider")
	}

	var ips []string

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.FilesMerklePrefix(req.Merkle))

	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {

		var file types.UnifiedFile
		if err := k.cdc.Unmarshal(iterator.Value(), &file); err != nil {
			continue
		}

		for _, proof := range file.Proofs {

			p, found := k.GetProofWithBuiltKey(ctx, []byte(proof))
			if !found {
				continue
			}

			prover := p.Prover

			provider, found := k.GetProviders(ctx, prover)
			if !found {
				continue
			}

			ips = append(ips, provider.Ip)
		}
	}

	return &types.QueryFindFileResponse{ProviderIps: ips}, nil
}
