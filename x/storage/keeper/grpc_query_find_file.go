package keeper

import (
	"context"
	"encoding/hex"
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListFiles(ctx sdk.Context, fid string) []string {
	providers := []string{}

	success := true

	res, found := k.GetFidCid(ctx, fid)
	if found {
		var cids []string

		err := json.Unmarshal([]byte(res.Cids), &cids)
		if err != nil {
			success = false
		} else {
			for _, cid := range cids {
				deal, found := k.GetActiveDeals(ctx, cid)
				if !found {
					continue
				}

				providers = append(providers, deal.Provider)
			}
		}

	}

	if !found || !success {
		allDeals := k.GetAllActiveDeals(ctx)

		for i := 0; i < len(allDeals); i++ {
			if allDeals[i].Fid == fid {
				provider, ok := k.GetProviders(ctx, allDeals[i].Provider)
				if !ok {
					continue
				}

				providers = append(providers, provider.Ip)
			}
		}
	}

	return providers
}

func (k Keeper) ListFilesByMerkle(ctx sdk.Context, merkle []byte) []string {
	providers := make([]string, 0)
	m := hex.EncodeToString(merkle)

	k.IterateActiveDeals(ctx, func(deal types.ActiveDeals) bool {
		if deal.Merkle == m {
			providers = append(providers, deal.Provider)
		}
		return false
	})

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

func (k Keeper) FindSomeFile(goCtx context.Context, req *types.QueryFindFile) (*types.QueryFindSomeFileResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	ls := k.ListFilesByMerkle(ctx, req.Merkle)

	return &types.QueryFindSomeFileResponse{ProviderIps: ls}, nil
}
