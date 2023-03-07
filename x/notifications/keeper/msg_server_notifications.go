package keeper

import (
	"context"
	"encoding/json"
	"fmt"

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

	// Check if sender is permitted to notify

	if !isSender(notiCounter, msg.Creator) {
		return nil, types.ErrCannotAddSenders
	}

	notifications := types.Notifications{
		Sender:       msg.Creator,
		Count:        notiCounter.Counter,
		Notification: msg.Notification,
		Address:      msg.Address,
		// hashPath and hashPathOwner not needed in this module. Will be used in filetree
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
		msg.Creator, // this needs to be fleshed out with permissions checking
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
		msg.Creator, // this needs to be fleshed out with permissions checking
	)

	return &types.MsgDeleteNotificationsResponse{}, nil
}
