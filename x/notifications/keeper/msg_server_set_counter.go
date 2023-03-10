package keeper

import (
	"context"
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/x/notifications/types"
)

// This needs to be inside of filetree init
func (k msgServer) SetCounter(goCtx context.Context, msg *types.MsgSetCounter) (*types.MsgSetCounterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetNotiCounter(
		ctx,
		msg.Creator,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "counter already set")
	}

	// Place holder map for blocked senders

	placeholderMap := make([]string, 0, 1000)
	marshalledBlockedSenders, err := json.Marshal(placeholderMap)
	if err != nil {
		return nil, types.ErrCantUnmarshall
	}

	BlockedSenders := string(marshalledBlockedSenders)

	counter := types.NotiCounter{
		Address:        msg.Creator,
		Counter:        0,
		BlockedSenders: BlockedSenders,
	}

	k.SetNotiCounter(
		ctx,
		counter,
	)

	return &types.MsgSetCounterResponse{NotiCounter: counter.Counter}, nil
}
