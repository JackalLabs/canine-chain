package keeper

import (
	"context"
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/notifications/types"
)

func (k msgServer) AddSenders(goCtx context.Context, msg *types.MsgAddSenders) (*types.MsgAddSendersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	notiCounter, found := k.GetNotiCounter(ctx, msg.Creator)
	if !found {
		return nil, types.ErrNotiCounterNotFound
	}

	//This message is already set to only allow the msg.Creator to add to their own notiCounter, but add this in just in case
	if !(notiCounter.Address == msg.Creator) {
		return nil, types.ErrCannotAddSenders
	}

	currentSenders := notiCounter.PermittedSenders

	fmt.Println("@@@@@@@@@@@@@@CURRENT SENDERS ARE", currentSenders)

	placeholderMap := make([]string, 0, 1000) //Perhaps I could just use an array
	json.Unmarshal([]byte(currentSenders), &placeholderMap)

	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@ PLACEHOLDER MAP CONTAINS", placeholderMap)
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@ PASSED IN JSON SENDERS CONTAINS", msg.SenderIds)

	temporaryMap := make([]string, 0, 1000) //Perhaps I could just use an array
	json.Unmarshal([]byte(msg.SenderIds), &temporaryMap)

	//SenderIds := strings.Split(temporaryMap, ",")

	//fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@ SenderIds CONTAINS", SenderIds)

	placeholderMap = append(placeholderMap, temporaryMap...)

	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@ PLACEHOLDER MAP CONTAINS", placeholderMap)

	marshalledSenders, err := json.Marshal(placeholderMap)
	if err != nil {
		return nil, types.ErrCantUnmarshall
	}

	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@ MARSHALED SENDERS CONTAINS", marshalledSenders)

	updatedSenders := string(marshalledSenders)

	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@ UPDATED SENDERS CONTAINS", updatedSenders)

	notiCounter.PermittedSenders = updatedSenders

	k.SetNotiCounter(ctx, notiCounter)

	return &types.MsgAddSendersResponse{}, nil
}

// func isSender(currentSenders []string, user string) bool {

// 	for _, v := range currentSenders {
// 		if v == user {
// 			return true
// 			break
// 		}
// 	}
// 	return false

// }
