package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/jackal-dao/canine/x/jklmining/types"
)

func (k msgServer) AllowSave(goCtx context.Context, msg *types.MsgAllowSave) (*types.MsgAllowSaveResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetSaveRequests(
		ctx,
		msg.Passkey,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var saveRequests = types.SaveRequests{
		Creator:  msg.Creator,
		Index:    msg.Passkey,
		Size_:    msg.Size_,
		Approved: "false",
	}

	k.SetSaveRequests(
		ctx,
		saveRequests,
	)

	return &types.MsgAllowSaveResponse{}, nil
}
