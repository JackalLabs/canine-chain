package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/filetree/types"
)

func (k msgServer) ChangeOwner(goCtx context.Context, msg *types.MsgChangeOwner) (*types.MsgChangeOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	file, found := k.GetFiles(ctx, msg.Address, msg.FileOwner)
	if !found {
		return nil, types.ErrFileNotFound
	}
	//Only the owner of a file can give it away
	isOwner := IsOwner(file, msg.Creator)
	if !isOwner {
		return nil, types.ErrCantGiveAway
	}

	//If the new owner already has a file set with this path,do not change
	//ownership--else their file will be overwridden
	_, fnd := k.GetFiles(ctx, msg.Address, msg.NewOwner)
	if fnd {
		return nil, types.ErrAlreadyExists
	}

	file.Owner = msg.NewOwner

	k.SetFiles(ctx, file)
	//Delete old file
	k.RemoveFiles(ctx, msg.Address, msg.FileOwner)

	return &types.MsgChangeOwnerResponse{}, nil
}
