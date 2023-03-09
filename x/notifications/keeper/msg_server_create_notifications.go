package keeper

import (
	"context"
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/x/notifications/types"
)

func (k msgServer) CreateNotifications(goCtx context.Context, msg *types.MsgCreateNotifications) (*types.MsgCreateNotificationsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Find the notiCounter
	notiCounter, found := k.GetNotiCounter(
		ctx,
		msg.Address,
	)
	if !found {
		return nil, types.ErrNotiCounterNotFound
	}

	// Check if the value already exists
	_, isFound := k.GetNotifications(
		ctx,
		notiCounter.Counter,
		msg.Address,
	)
	if isFound {
		return nil, types.ErrNotificationAlreadySet
	}

	// Check if sender is permitted to notify

	blocked, err := isBlocked(notiCounter, msg.Creator)
	if err != nil {
		return nil, types.ErrCantUnmarshall
	}

	if blocked {
		return nil, types.ErrBlockedSender
	}

	notifications := types.Notifications{
		Sender:       msg.Creator,
		Count:        notiCounter.Counter,
		Notification: msg.Notification,
		Address:      msg.Address,
	}

	k.SetNotifications(
		ctx,
		notifications,
		msg.Address,
	)

	notiCounter.Counter++

	k.SetNotiCounter(
		ctx,
		notiCounter,
	)

	return &types.MsgCreateNotificationsResponse{}, nil
}

func isBlocked(notiCounter types.NotiCounter, user string) (bool, error) {
	BlockedSenders := notiCounter.BlockedSenders

	placeholderMap := make([]string, 0, 1000)
	err := json.Unmarshal([]byte(BlockedSenders), &placeholderMap)
	if err != nil {
		return false, types.ErrCantUnmarshall
	}

	for _, v := range placeholderMap {
		if v == user {
			return true, nil
		}
	}
	return false, nil
}

// DOES NOT WORK - stub for now
// I don't think update is needed. Seems pointless to overwrite a current notification--just append to the end
func (k msgServer) UpdateNotifications(goCtx context.Context, msg *types.MsgUpdateNotifications) (*types.MsgUpdateNotificationsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetNotifications(
		ctx,
		msg.Count,
		msg.Address,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Address {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	notifications := types.Notifications{
		Sender:       msg.Creator,
		Count:        msg.Count,
		Notification: msg.Notification,
		Address:      msg.Address,
	}

	k.SetNotifications(ctx, notifications, msg.Address)

	return &types.MsgUpdateNotificationsResponse{}, nil
}
