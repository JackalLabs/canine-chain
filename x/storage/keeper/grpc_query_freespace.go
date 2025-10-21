package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v5/x/storage/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) FreeSpace(goCtx context.Context, req *types.QueryFreeSpace) (*types.QueryFreeSpaceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	num := k.GetProviderUsing(ctx, req.Address)

	provider, found := k.GetProviders(ctx, req.Address)

	if !found {
		return nil, fmt.Errorf("can't find provider")
	}

	space, ok := sdk.NewIntFromString(provider.Totalspace)

	if !ok {
		return nil, fmt.Errorf("can't parse total space")
	}

	return &types.QueryFreeSpaceResponse{
		Space: space.Int64() - num,
	}, nil
}

func (k Keeper) StoreCount(goCtx context.Context, req *types.QueryStoreCount) (*types.QueryStoreCountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	num := k.GetProviderDeals(ctx, req.Address)

	return &types.QueryStoreCountResponse{
		Count: num,
	}, nil
}
