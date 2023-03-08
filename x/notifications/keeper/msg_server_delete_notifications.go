package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/notifications/types"
)

// Delete the latest message
func (k msgServer) DeleteNotifications(goCtx context.Context, msg *types.MsgDeleteNotifications) (*types.MsgDeleteNotificationsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	notiCounter, found := k.GetNotiCounter(
		ctx,
		msg.Creator,
	)
	if !found {
		return nil, types.ErrNotiCounterNotFound
	}

	notification, found := k.GetNotifications(
		ctx,
		notiCounter.Counter-1,
		msg.Creator,
	)
	if !found {
		return nil, types.ErrNotificationNotFound
	}

	// Checks if the the msg creator is the same as the current owner of this notification
	if msg.Creator != notification.Address {
		return nil, types.ErrNotNotificationOwner
	}

	k.RemoveNotifications(
		ctx,
		notiCounter.Counter,
		msg.Creator,
	)

	notiCounter.Counter--

	k.SetNotiCounter(
		ctx,
		notiCounter,
	)

	return &types.MsgDeleteNotificationsResponse{}, nil
}
