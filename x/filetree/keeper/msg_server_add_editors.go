package keeper

import (
	"context"
	"encoding/json"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/filetree/types"
)

func (k msgServer) AddEditors(goCtx context.Context, msg *types.MsgAddEditors) (*types.MsgAddEditorsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	file, found := k.GetFiles(ctx, msg.Address, msg.Fileowner)
	if !found {
		return nil, types.ErrFileNotFound
	}
	//CONSIDER: ONLY THE OWNER CAN ADD EDITORS?
	hasEdit := HasEditAccess(file, msg.Creator)
	if !hasEdit {
		return nil, types.ErrNoAccess
	}

	peacc := file.EditAccess

	jeacc := make(map[string]string)
	json.Unmarshal([]byte(peacc), &jeacc)

	ids := strings.Split(msg.EditorIds, ",")
	keys := strings.Split(msg.EditorKeys, ",")

	for i, v := range ids {
		jeacc[v] = keys[i]
	}

	eaccbytes, err := json.Marshal(jeacc)
	if err != nil {
		return nil, types.ErrCantMarshall
	}
	newEditors := string(eaccbytes)

	file.EditAccess = newEditors

	k.SetFiles(ctx, file)

	//notify editors
	bool, error := notify(k, ctx, msg.NotifyEditors, msg.NotiForEditors, msg.Creator, file.Address, file.Owner)
	if !bool {
		return nil, error
	}

	return &types.MsgAddEditorsResponse{}, nil
}
