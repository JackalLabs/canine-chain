package keeper

import (
	"context"
	"encoding/json"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/filetree/types"
)

func (k msgServer) RemoveEditors(goCtx context.Context, msg *types.MsgRemoveEditors) (*types.MsgRemoveEditorsResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	file, found := k.GetFiles(ctx, msg.Address, msg.Fileowner)
	if !found {
		return nil, types.ErrFileNotFound
	}

	isOwner := IsOwner(file, msg.Creator)

	if !isOwner {
		return nil, types.ErrNotOwner
	}

	//Continue
	peacc := file.EditAccess

	jeacc := make(map[string]string)
	json.Unmarshal([]byte(peacc), &jeacc)

	ids := strings.Split(msg.EditorIds, ",")
	for _, v := range ids {
		delete(jeacc, v)
	}

	vaccbytes, err := json.Marshal(jeacc)
	if err != nil {
		return nil, types.ErrCantMarshall
	}
	newEditors := string(vaccbytes)

	file.EditAccess = newEditors

	k.SetFiles(ctx, file)

	//notify editors
	bool, error := notify(k, ctx, msg.NotifyEditors, msg.NotiForEditors, msg.Creator, file.Address, file.Owner)
	if !bool {
		return nil, error
	}

	return &types.MsgRemoveEditorsResponse{}, nil
}
