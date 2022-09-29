package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/notifications/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) FilteredNotifications(c context.Context, req *types.QueryFilteredNotificationsRequest) (*types.QueryFilteredNotificationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	i := uint64(0)
	var notifications []types.Notifications
	for i < req.MaxCount {
		val, found := k.GetNotifications(
			ctx,
			i,
			req.Address,
		)
		if !found {
			return nil, status.Error(codes.NotFound, "not found")
		}
		notifications = append(notifications, val)
		i += 1
	}

	return &types.QueryFilteredNotificationsResponse{Notifications: notifications}, nil
}
