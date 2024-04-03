package keeper

import (
	"context"
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/filetree/types"
)

func (k msgServer) ResetEditors(goCtx context.Context, msg *types.MsgResetEditors) (*types.MsgResetEditorsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	file, found := k.GetFiles(ctx, msg.Address, msg.FileOwner)
	if !found {
		return nil, types.ErrFileNotFound
	}

	isOwner := IsOwner(file, msg.Creator)
	if !isOwner {
		return nil, types.ErrNotOwner
	}

	ownerEditorAddress := MakeEditorAddress(file.TrackingNumber, msg.Creator)

	peacc := file.EditAccess
	// Unmarshall current edit access to this blank map
	jeacc := make(map[string]string)
	if err := json.Unmarshal([]byte(peacc), &jeacc); err != nil {
		return nil, types.ErrCantUnmarshall
	}

	ownerKey := jeacc[ownerEditorAddress]

	resetEditors := make(map[string]string)
	resetEditors[ownerEditorAddress] = ownerKey

	eaccbytes, err := json.Marshal(resetEditors)
	if err != nil {
		return nil, types.ErrCantMarshall
	}
	newEditors := string(eaccbytes)

	file.EditAccess = newEditors

	k.SetFiles(ctx, file)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeResetEditors,
			sdk.NewAttribute(types.AttributeKeySigner, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyFileAddress, msg.Address),
		),
	)
	return &types.MsgResetEditorsResponse{}, nil
}
