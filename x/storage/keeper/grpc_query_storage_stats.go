package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StorageStats(c context.Context, req *types.QueryStorageStatsRequest) (*types.QueryStorageStatsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	payment := k.GetAllStoragePaymentInfo(ctx)

	var spacePurchased int64
	var spaceUsed int64
	var activeUsers uint64

	for _, info := range payment {
		if info.End.Before(time.Now()) {
			continue
		}
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
	}, nil
}
