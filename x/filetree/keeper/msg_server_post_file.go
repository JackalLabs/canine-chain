package keeper

import (
	"context"
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/filetree/types"
	notiTypes "github.com/jackal-dao/canine/x/notifications/types"
)

func (k msgServer) PostFile(goCtx context.Context, msg *types.MsgPostFile) (*types.MsgPostFileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	parentOwnerString := MakeOwnerAddress(msg.HashParent, msg.Account)

	parentFile, found := k.GetFiles(ctx, msg.HashParent, parentOwnerString)
	if !found {
		return nil, types.ErrParentFileNotFound
	}

	hasEdit := HasEditAccess(parentFile, msg.Creator)
	if !hasEdit {
		return nil, types.ErrCannotWrite
	}

	//Make the full path
	fullMerklePath := types.AddToMerkle(msg.HashParent, msg.HashChild)

	owner := MakeOwnerAddress(fullMerklePath, msg.Account)

	file := types.Files{
		Address:        fullMerklePath,
		Contents:       msg.Contents,
		Owner:          owner,
		ViewingAccess:  msg.Viewers,
		EditAccess:     msg.Editors,
		TrackingNumber: msg.TrackingNumber,
	}

	incrementTracker(k, ctx, msg)
	k.SetFiles(ctx, file)

	//notify viewers
	bool, error := notify(k, ctx, msg.NotifyViewers, string("you have viewer access"), msg.Creator)
	if !bool {
		return nil, error
	}

	//notify editors
	ok, err := notify(k, ctx, msg.NotifyEditors, string("you have editor access"), msg.Creator)
	if !ok {
		return nil, err
	}

	return &types.MsgPostFileResponse{Path: fullMerklePath}, nil
}

//if bool returns 'true', we successfully notified everyone, otherwise if it's false we return the error
//viewers will have their own message from editors, so should send in a general notification, and a string of viewers or editors

func notify(k msgServer, ctx sdk.Context, recipients string, notification string, sender string) (bool, error) {

	placeholderMap := make([]string, 0, 1000)
	json.Unmarshal([]byte(recipients), &placeholderMap)

	for _, v := range placeholderMap {
		// Find the notiCounter
		notiCounter, found := k.notiKeeper.GetNotiCounter(
			ctx,
			v,
		)

		if !found {
			return false, notiTypes.ErrNotiCounterNotFound
		}

		// Check if the notification already exists. Should always come back false because recipient's notiCounter is incremented everytime someone sends them a msg
		_, isFound := k.notiKeeper.GetNotifications(
			ctx,
			notiCounter.Counter,
			v,
		)
		//If it exists, we return false to return the error
		if isFound {
			return false, notiTypes.ErrNotificationAlreadySet
		}

		//Check if sender is permitted to notify
		if !isSender(notiCounter, sender) {
			return false, notiTypes.ErrCannotAddSenders
		}

		var notifications = notiTypes.Notifications{
			Sender:       sender, //sender of the notification--who in this case is the poster of the file
			Count:        notiCounter.Counter,
			Notification: notification, //need extra param in MsgPostFile
			Address:      v,            //The address of the recipient--their list of notifications
			//merklePath of file
			//hashPathOwner // this is here because the sender of the file won't always be the owner
		}

		k.notiKeeper.SetNotifications(
			ctx,
			notifications,
		)

		notiCounter.Counter += 1

		k.notiKeeper.SetNotiCounter(
			ctx,
			notiCounter,
		)
	}

	return true, nil

}

func isSender(notiCounter notiTypes.NotiCounter, user string) bool {

	currentSenders := notiCounter.PermittedSenders

	placeholderMap := make([]string, 0, 1000)
	json.Unmarshal([]byte(currentSenders), &placeholderMap)

	for _, v := range placeholderMap {

		if string(v) == user {
			return true
		}
	}
	return false

}
