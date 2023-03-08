package keeper

import (
	"context"
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/notifications/types"
)

func (k msgServer) BlockSenders(goCtx context.Context, msg *types.MsgBlockSenders) (*types.MsgBlockSendersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	notiCounter, found := k.GetNotiCounter(ctx, msg.Creator)
	if !found {
		return nil, types.ErrNotiCounterNotFound
	}

	// This message is already set to only allow the msg.Creator to add to their own notiCounter, but add this in just in case
	if !(notiCounter.Address == msg.Creator) {
		return nil, types.ErrOnlyOwnerCanBlock
	}

	BlockedSenders := notiCounter.BlockedSenders

	placeholderMap := make([]string, 0, 1000) // Perhaps I could just use an array
	json.Unmarshal([]byte(BlockedSenders), &placeholderMap)

	temporaryMap := make([]string, 0, 1000) // Perhaps I could just use an array
	json.Unmarshal([]byte(msg.SenderIds), &temporaryMap)

	placeholderMap = append(placeholderMap, temporaryMap...)

	marshalledSenders, err := json.Marshal(placeholderMap)
	if err != nil {
		return nil, types.ErrCantUnmarshall
	}

	updatedBlockedSenders := string(marshalledSenders)

	notiCounter.BlockedSenders = updatedBlockedSenders

	k.SetNotiCounter(ctx, notiCounter)

	return &types.MsgBlockSendersResponse{}, nil
}
