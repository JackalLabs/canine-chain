package keeper

import (
	"context"
	"encoding/json"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/filetree/types"
)

func (k msgServer) AddViewers(goCtx context.Context, msg *types.MsgAddViewers) (*types.MsgAddViewersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	file, found := k.GetFiles(ctx, msg.Address, msg.Fileowner)
	if !found {
		return nil, types.ErrFileNotFound
	}

	isOwner := IsOwner(file, msg.Creator)
	if !isOwner {
		return nil, types.ErrCannotAllowView
	}

	pvacc := file.ViewingAccess

	jvacc := make(map[string]string)
	if err := json.Unmarshal([]byte(pvacc), &jvacc); err != nil {
		return nil, types.ErrCantUnmarshall
	}

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

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAddViewers,
			sdk.NewAttribute(types.AttributeKeySigner, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyFileAddress, msg.Address),
		),
	)

	return &types.MsgAddViewersResponse{}, nil
}
