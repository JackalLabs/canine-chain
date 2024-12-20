package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NetworkSize(c context.Context, req *types.QueryNetworkSize) (*types.QueryNetworkSizeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	var s uint64

	k.IterateAndParseFilesByMerkle(ctx, false, func(_ []byte, val types.UnifiedFile) bool {
		s += uint64(val.FileSize * val.MaxProofs)

		return false
	})

	return &types.QueryNetworkSizeResponse{Size_: s}, nil
}

func (k Keeper) AvailableSpace(c context.Context, req *types.QueryAvailableSpace) (*types.QueryAvailableSpaceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	var s uint64

	providers := k.GetAllActiveProviders(ctx)

	for _, provider := range providers {
		providerEntry, found := k.GetProviders(ctx, provider.Address)
		if !found {
			continue
		}
		space, err := strconv.ParseInt(providerEntry.Totalspace, 10, 64)
		if err != nil {
			continue
		}

		s += uint64(space)
	}

	return &types.QueryAvailableSpaceResponse{Size_: s}, nil
}

func (k Keeper) StorageStats(c context.Context, req *types.QueryStorageStats) (*types.QueryStorageStatsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	payment := k.GetAllStoragePaymentInfo(ctx)

	var spacePurchased int64
	var spaceUsed int64

	activeUsers := make(map[string]bool)
	allUsers := make(map[string]bool)

	usersByPlan := make(map[int64]int64)

	for _, info := range payment {
		allUsers[info.Address] = true

		if info.End.Before(ctx.BlockTime()) {
			continue
		}
		usersByPlan[info.SpaceAvailable]++
		activeUsers[info.Address] = true
		spacePurchased += info.SpaceAvailable
	}

	k.IterateAndParseFilesByMerkle(ctx, false, func(_ []byte, val types.UnifiedFile) bool {
		allUsers[val.Owner] = true
		activeUsers[val.Owner] = true

		m := val.FileSize * val.MaxProofs

		spaceUsed += m

		return false
	})

	decSpent := sdk.NewDec(spacePurchased)
	decUsed := sdk.NewDec(spaceUsed)

	ratio := decUsed.Quo(decSpent).MulInt64(100)

	return &types.QueryStorageStatsResponse{
		Purchased:   uint64(spacePurchased),
		Used:        uint64(spaceUsed),
		UsedRatio:   ratio,
		ActiveUsers: uint64(len(activeUsers)),
		UniqueUsers: uint64(len(allUsers)),
		UsersByPlan: usersByPlan,
	}, nil
}
