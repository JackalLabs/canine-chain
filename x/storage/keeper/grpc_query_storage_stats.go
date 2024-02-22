package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
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

	totalUsers := make(map[string]bool)

	for _, info := range payment {
		totalUsers[info.Address] = true // counting in total users
		if info.End.Before(ctx.BlockTime()) {
			continue
		}
		spacePurchased += info.SpaceAvailable
		spaceUsed += info.SpaceUsed
	}

	decSpent := sdk.NewDec(spacePurchased)
	decUsed := sdk.NewDec(spaceUsed)

	ratio := decUsed.Quo(decSpent).MulInt64(100)

	users := make(map[string]bool)

	var permSize int64
	k.IterateActiveDeals(ctx, func(deal types.ActiveDeals) bool {
		users[deal.Creator] = true
		totalUsers[deal.Creator] = true

		if deal.Endblock == "0" {
			return false
		}

		s := deal.Filesize
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			ctx.Logger().Debug("cannot parse active deal")
			return false
		}

		permSize += i

		return false
	})

	return &types.QueryStorageStatsResponse{
		Purchased:   spacePurchased,
		Used:        spaceUsed,
		UsedRatio:   ratio,
		ActiveUsers: int64(len(users)),
		UniqueUsers: int64(len(totalUsers)),
		PermUsed:    permSize,
	}, nil
}
