package keeper

import (
	"context"
	"encoding/json"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/filetree/types"
)

func (k msgServer) RemoveViewers(goCtx context.Context, msg *types.MsgRemoveViewers) (*types.MsgRemoveViewersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	file, found := k.GetFiles(ctx, msg.Address, msg.Fileowner)
	if !found {
		return nil, types.ErrFileNotFound
	}
	//This was previously: 'hasEditAccess', but this means that An editor can remove a viewer? So, in a file owned by Charlie, Alice--if an editor--can remove Bob's
	//viewing access while Bob is also an editor. Bob could add himself back in as a viewer but it would be so laborous
	isOwner := IsOwner(file, msg.Creator)

	if !isOwner {
		return nil, types.ErrNotOwner
	}

	pvacc := file.ViewingAccess

	jvacc := make(map[string]string)
	json.Unmarshal([]byte(pvacc), &jvacc)

	ids := strings.Split(msg.ViewerIds, ",")
	for _, v := range ids {
		delete(jvacc, v)
	}

	vaccbytes, err := json.Marshal(jvacc)
	if err != nil {
		return nil, types.ErrCantMarshall
	}
	newviewers := string(vaccbytes)

	file.ViewingAccess = newviewers

	k.SetFiles(ctx, file)

	//notify viewers
	bool, error := notify(k, ctx, msg.Notifyviewers, msg.NotiForViewers, msg.Creator, file.Address, file.Owner)
	if !bool {
		return nil, error
	}

	return &types.MsgRemoveViewersResponse{}, nil
}
