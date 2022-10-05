package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/filetree/types"
)

func (k msgServer) PostFile(goCtx context.Context, msg *types.MsgPostFile) (*types.MsgPostFileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	parentOwnerString := MakeOwnerAddress(msg.HashParent, msg.Account)

	parentFile, found := k.GetFiles(ctx, msg.HashParent, parentOwnerString)
	if !found {
		return nil, types.ErrParentFileNotFound
	}

	hasEdit := HasEditAccess(parentFile, msg.Creator)
	if !hasEdit {
		return nil, types.ErrCannotWrite
	}

	//Make the full path
	fullMerklePath := types.AddToMerkle(msg.HashParent, msg.HashChild)

	owner := MakeOwnerAddress(fullMerklePath, msg.Account)

	file := types.Files{
		Address:        fullMerklePath,
		Contents:       msg.Contents,
		Owner:          owner,
		ViewingAccess:  msg.Viewers,
		EditAccess:     msg.Editors,
		TrackingNumber: msg.TrackingNumber,
	}

	incrementTracker(k, ctx, msg)
	k.SetFiles(ctx, file)

	//notify viewers
	bool, error := notify(k, ctx, msg.NotifyViewers, string("you have viewer access from postfile"), msg.Creator, fullMerklePath, owner)
	if !bool {
		return nil, error
	}

	//notify editors
	ok, err := notify(k, ctx, msg.NotifyEditors, string("you have editor access from postfile"), msg.Creator, fullMerklePath, owner)
	if !ok {
		return nil, err
	}

	return &types.MsgPostFileResponse{Path: fullMerklePath}, nil
}
