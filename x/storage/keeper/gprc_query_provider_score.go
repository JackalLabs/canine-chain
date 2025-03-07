package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ProviderScore(c context.Context, req *types.QueryProviderScore) (*types.QueryProviderScoreResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetProviderScore(ctx, req.Provider)

	if !found {
		return nil, status.Error(codes.NotFound, "provider not found")
	}

	return &types.QueryProviderScoreResponse{TotalSize: val.TotalSize}, nil
}
