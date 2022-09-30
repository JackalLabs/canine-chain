package keeper

import (
	"context"

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

	var counter = types.NotiCounter{
		Address: msg.Creator,
		Counter: 0,
	}

	k.SetNotiCounter(
		ctx,
		counter,
	)

	return &types.MsgSetCounterResponse{}, nil
}
