package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StorageStats(c context.Context, req *types.QueryStorageStats) (*types.QueryStorageStatsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	payment := k.GetAllStoragePaymentInfo(ctx)

	var spacePurchased int64
	var spaceUsed int64
	var activeUsers uint64
	var totalUsers uint64

	usersByPlan := make(map[int64]int64)

	for _, info := range payment {
		totalUsers++ // counting in total users
		if info.End.Before(ctx.BlockTime()) {
			continue
		}
		usersByPlan[info.SpaceAvailable]++
		spacePurchased += info.SpaceAvailable
		spaceUsed += info.SpaceUsed
		activeUsers++
	}

	decSpent := sdk.NewDec(spacePurchased)
	decUsed := sdk.NewDec(spaceUsed)

	ratio := decUsed.Quo(decSpent).MulInt64(100)

	return &types.QueryStorageStatsResponse{
		Purchased:   uint64(spacePurchased),
		Used:        uint64(spaceUsed),
		UsedRatio:   ratio,
		ActiveUsers: activeUsers,
		UniqueUsers: totalUsers,
		UsersByPlan: usersByPlan,
	}, nil
}
