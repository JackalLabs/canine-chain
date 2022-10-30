package keeper

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	notiTypes "github.com/jackal-dao/canine/x/notifications/types"
)


//viewers will have their own message from editors, so should send in a general notification, and a string of viewers or editors

func notify(k msgServer, ctx sdk.Context, recipients string, notification string, sender string, hashPath string, hashPathOwner string) (bool, error) {
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

		if isFound {
			return false, notiTypes.ErrNotificationAlreadySet
		}


		// if !isSender(notiCounter, sender) {
		// 	return false, notiTypes.ErrCannotAddSenders
		// }

		notifications := notiTypes.Notifications{
			Sender:       sender, // delete this for security?
			Count:        notiCounter.Counter,
			Notification: notification,


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



// func isSender(notiCounter notiTypes.NotiCounter, user string) bool {

// 	currentSenders := notiCounter.PermittedSenders

// 	placeholderMap := make([]string, 0, 1000)
// 	json.Unmarshal([]byte(currentSenders), &placeholderMap)

// 	for _, v := range placeholderMap {

// 		if string(v) == user {
// 			return true
// 		}
// 	}
// 	return false

// }
