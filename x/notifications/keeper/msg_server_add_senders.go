package keeper

import (
	"context"
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/notifications/types"
)

func (k msgServer) AddSenders(goCtx context.Context, msg *types.MsgAddSenders) (*types.MsgAddSendersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	notiCounter, found := k.GetNotiCounter(ctx, msg.Creator)
	if !found {
		return nil, types.ErrNotiCounterNotFound
	}

	// This message is already set to only allow the msg.Creator to add to their own notiCounter, but add this in just in case
	if !(notiCounter.Address == msg.Creator) {
		return nil, types.ErrCannotAddSenders
	}

	currentSenders := notiCounter.PermittedSenders

	placeholderMap := make([]string, 0, 1000) // Perhaps I could just use an array
	json.Unmarshal([]byte(currentSenders), &placeholderMap)

	temporaryMap := make([]string, 0, 1000) // Perhaps I could just use an array
	json.Unmarshal([]byte(msg.SenderIds), &temporaryMap)

	placeholderMap = append(placeholderMap, temporaryMap...)

	marshalledSenders, err := json.Marshal(placeholderMap)
	if err != nil {
		return nil, types.ErrCantUnmarshall
	}

	updatedSenders := string(marshalledSenders)

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
