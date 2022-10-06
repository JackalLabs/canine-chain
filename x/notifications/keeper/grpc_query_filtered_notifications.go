package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackal-dao/canine/x/notifications/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Query all the notifications that belong to an address
func (k Keeper) FilteredNotifications(c context.Context, req *types.QueryFilteredNotificationsRequest) (*types.QueryFilteredNotificationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	// Find the notiCounter
	notiCounter, found := k.GetNotiCounter(
		ctx,
		req.Address,
	)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Recipients notiCounter not set")
	}

	i := uint64(0)
	var notifications []string

	for i < notiCounter.Counter {
		val, found := k.GetNotifications(
			ctx,
			i,
			req.Address,
		)
		if !found {
			return nil, status.Error(codes.NotFound, "not found")
		}
		notifications = append(notifications, val.Notification)
		i += 1
	}

	return &types.QueryFilteredNotificationsResponse{Notifications: notifications}, nil
}
