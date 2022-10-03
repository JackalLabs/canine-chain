package keeper

import (
	"context"
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackal-dao/canine/x/notifications/types"
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

	//Add yourself as a permitted Sender in the beginning so you can notify yourself

	placeholderMap := make([]string, 0, 1000) //Perhaps I could just use an array
	placeholderMap = append(placeholderMap, msg.Creator)
	marshalledSenders, err := json.Marshal(placeholderMap)
	if err != nil {
		return nil, types.ErrCantUnmarshall
	}

	updatedSenders := string(marshalledSenders)

	var counter = types.NotiCounter{
		Address:          msg.Creator,
		Counter:          0,
		PermittedSenders: updatedSenders,
	}

	k.SetNotiCounter(
		ctx,
		counter,
	)

	return &types.MsgSetCounterResponse{}, nil
}
