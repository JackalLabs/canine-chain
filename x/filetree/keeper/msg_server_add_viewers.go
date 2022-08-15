package keeper

import (
	"context"
	"encoding/json"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/filetree/types"
)

func (k msgServer) AddViewers(goCtx context.Context, msg *types.MsgAddViewers) (*types.MsgAddViewersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	file, found := k.GetFiles(ctx, msg.Address)
	if !found {
		return nil, types.ErrFileNotFound
	}

	hasEdit := HasEditAccess(file, msg.Creator)
	if !hasEdit {
		return nil, types.ErrNoAccess
	}

	pvacc := file.ViewingAccess

	jvacc := make(map[string]string)
	json.Unmarshal([]byte(pvacc), &jvacc)

	ids := strings.Split(msg.ViewerIds, ",")
	keys := strings.Split(msg.ViewerKeys, ",")

	for i, v := range ids {
		jvacc[v] = keys[i]
	}

	vaccbytes, err := json.Marshal(jvacc)
	if err != nil {
		return nil, types.ErrCantMarshall
	}
	newviewers := string(vaccbytes)

	file.ViewingAccess = newviewers

	k.SetFiles(ctx, file)

	return &types.MsgAddViewersResponse{}, nil
}
