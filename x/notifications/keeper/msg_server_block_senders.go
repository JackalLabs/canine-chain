package keeper

import (
	"context"
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/notifications/types"
)

func (k msgServer) BlockSenders(goCtx context.Context, msg *types.MsgBlockSenders) (*types.MsgBlockSendersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	notiCounter, found := k.GetNotiCounter(ctx, msg.Creator)
	if !found {
		return nil, types.ErrNotiCounterNotFound
	}

	if !(notiCounter.Address == msg.Creator) {
		return nil, types.ErrOnlyOwnerCanBlock
	}

	BlockedSenders := notiCounter.BlockedSenders

	placeholderMap := make([]string, 0, 1000)
	err := json.Unmarshal([]byte(BlockedSenders), &placeholderMap)
	if err != nil {
		return nil, types.ErrCantUnmarshall
	}

	temporaryMap := make([]string, 0, 1000)
	err = json.Unmarshal([]byte(msg.SenderIds), &temporaryMap)

	if err != nil {
		return nil, types.ErrCantUnmarshall
	}

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
