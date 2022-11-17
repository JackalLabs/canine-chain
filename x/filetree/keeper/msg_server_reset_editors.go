package keeper

import (
	"context"
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/filetree/types"
)

func (k msgServer) ResetEditors(goCtx context.Context, msg *types.MsgResetEditors) (*types.MsgResetEditorsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	file, found := k.GetFiles(ctx, msg.Address, msg.Fileowner)
	if !found {
		return nil, types.ErrFileNotFound
	}

	isOwner := IsOwner(file, msg.Creator)
	if !isOwner {
		return nil, types.ErrNotOwner
	}

	ownerEditorAddress := MakeEditorAddress(file.TrackingNumber, msg.Creator)

	peacc := file.EditAccess
	//Unmarshall current edit access to this blank map
	jeacc := make(map[string]string)
	json.Unmarshal([]byte(peacc), &jeacc)

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

	return &types.MsgResetEditorsResponse{}, nil
}
