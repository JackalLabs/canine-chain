package keeper

import (
	"context"
	"strings"

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

	// Query proofs for files with matching merkle
	rows, err := k.filebase.Query(`
        SELECT p.proof
        FROM unified_files f
        JOIN proofs p ON f.merkle = p.file_merkle AND f.owner = p.file_owner AND f.start = p.file_start
        WHERE f.merkle = ?
    `, req.Merkle)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer rows.Close()

	// Process each proof
	for rows.Next() {
		var proofKey string
		if err := rows.Scan(&proofKey); err != nil {
			continue
		}

		// Use the KV store to get proof details
		p, found := k.GetProofWithBuiltKey(ctx, []byte(proofKey))
		if !found {
			continue
		}

		prover := p.Prover

		// Use the KV store to get provider details
		provider, found := k.GetProviders(ctx, prover)
		if !found {
			continue
		}

		ips = append(ips, provider.Ip)
	}

	return &types.QueryFindFileResponse{ProviderIps: ips}, nil
}
