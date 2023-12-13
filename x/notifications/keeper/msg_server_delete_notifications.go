package keeper

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/notifications/types"
)

// DeleteNotification deletes a given message
func (k msgServer) DeleteNotification(goCtx context.Context, msg *types.MsgDeleteNotification) (*types.MsgDeleteNotificationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	noti, found := k.GetNotification(ctx, msg.Creator, msg.From, msg.Time)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "cannot find notification in store")
	}

	if msg.Creator != noti.To {
		return nil, sdkerrors.Wrapf(types.ErrNotNotificationOwner, "you do not control this notification")
	}

	k.RemoveNotification(ctx, msg.Creator, msg.From, msg.Time)

	return &types.MsgDeleteNotificationResponse{}, nil
}
