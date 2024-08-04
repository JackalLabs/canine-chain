package keeper

import (
	"context"
	"encoding/json"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/x/filetree/types"
)

func (k msgServer) RemoveViewers(goCtx context.Context, msg *types.MsgRemoveViewers) (*types.MsgRemoveViewersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	file, found := k.GetFiles(ctx, msg.Address, msg.FileOwner)
	if !found {
		return nil, types.ErrFileNotFound
	}
	// This was previously: 'hasEditAccess', but this means that An editor can remove a viewer? So, in a file owned by Charlie, Alice--if an editor--can remove Bob's
	// viewing access while Bob is also an editor. Bob could add himself back in as a viewer but it would be so laborous
	isOwner := IsOwner(file, msg.Creator)

	if !isOwner {
		return nil, types.ErrNotOwner
	}

	pvacc := file.ViewingAccess

	jvacc := make(map[string]string)
	if err := json.Unmarshal([]byte(pvacc), &jvacc); err != nil {
		return nil, types.ErrCantUnmarshall
	}

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
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRemoveViewers,
			sdk.NewAttribute(types.AttributeKeySigner, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyFileAddress, msg.Address),
		),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeJackalMessage,
			sdk.NewAttribute(types.AttributeKeySigner, msg.Creator),
		),
	)
	return &types.MsgRemoveViewersResponse{}, nil
}
