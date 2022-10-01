package keeper

import (
	"context"
	"encoding/json"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/notifications/types"
)

func (k msgServer) AddSenders(goCtx context.Context, msg *types.MsgAddSenders) (*types.MsgAddSendersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	notiCounter, found := k.GetNotiCounter(ctx, msg.Creator)
	if !found {
		return nil, types.ErrNotiCounterNotFound
	}

	//Check if the whoever is adding senders is already a permitted sender. Need to add owner as a Sender in the beginning
	// hasEdit := HasEditAccess(file, msg.Creator)
	// if !hasEdit {
	// 	return nil, types.ErrNoAccess
	// }

	psacc := notiCounter.PermittedSenders

	jsacc := make(map[string]string) //Perhaps I could just use an array
	json.Unmarshal([]byte(psacc), &jsacc)

	ids := strings.Split(msg.SenderIds, ",")

	for _, v := range ids {
		jsacc[v] = "user is a sender"
	}

	saccbytes, err := json.Marshal(jsacc)
	if err != nil {
		return nil, types.ErrCantUnmarshall
	}
	newSenders := string(saccbytes)

	notiCounter.PermittedSenders = newSenders

	k.SetNotiCounter(ctx, notiCounter)

	return &types.MsgAddSendersResponse{}, nil
}
