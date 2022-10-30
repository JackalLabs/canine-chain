package keeper

import (
	"context"
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackal-dao/canine/x/notifications/types"
)

func (k msgServer) CreateNotifications(goCtx context.Context, msg *types.MsgCreateNotifications) (*types.MsgCreateNotificationsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Find the notiCounter
	notiCounter, found := k.GetNotiCounter(
		ctx,
		msg.Address,
	)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Recipients notiCounter not set")
	}

	// Check if the value already exists
	_, isFound := k.GetNotifications(
		ctx,
		notiCounter.Counter,
		msg.Address,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "notification already set")
	}

	if !isSender(notiCounter, msg.Creator) {
		return nil, types.ErrCannotAddSenders
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
	)

	notiCounter.Counter += 1

	k.SetNotiCounter(
		ctx,
		notiCounter,
	)

	return &types.MsgCreateNotificationsResponse{}, nil
}

func isSender(notiCounter types.NotiCounter, user string) bool {
	currentSenders := notiCounter.PermittedSenders

	placeholderMap := make([]string, 0, 1000)
	json.Unmarshal([]byte(currentSenders), &placeholderMap)
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@ PLACEHOLDERMAP IS", placeholderMap)

	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@ USER IS", user)

	for _, v := range placeholderMap {
		if string(v) == user {
			return true
		}
	}
	return false
}

// DOES NOT WORK
func (k msgServer) DeleteNotifications(goCtx context.Context, msg *types.MsgDeleteNotifications) (*types.MsgDeleteNotificationsResponse, error) {
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

	k.RemoveNotifications(
		ctx,
		msg.Count,
		msg.Creator, // TODO: is this correct? -  NOT CERTAIN THAT "CREATOR IS ADDRESS."
	)

	return &types.MsgDeleteNotificationsResponse{}, nil
}

// DOES NOT WORK
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

	k.SetNotifications(ctx, notifications)

	return &types.MsgUpdateNotificationsResponse{}, nil
}
