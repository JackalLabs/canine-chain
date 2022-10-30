package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/filetree/types"
)

func (k msgServer) DeleteFile(goCtx context.Context, msg *types.MsgDeleteFile) (*types.MsgDeleteFileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	ownerAddress := MakeOwnerAddress(msg.HashPath, msg.Account)

	file, found := k.GetFiles(ctx, msg.HashPath, ownerAddress)
	if !found {
		return nil, types.ErrFileNotFound
	}

	isOwner := IsOwner(file, msg.Creator)
	if !isOwner {
		return nil, types.ErrCannotDelete
	}

	k.RemoveFiles(ctx, msg.HashPath, ownerAddress)

	return &types.MsgDeleteFileResponse{}, nil
}
