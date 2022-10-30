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

	k.SetFiles(ctx, file)


	bool, error := notify(k, ctx, msg.ViewersToNotify, msg.NotiForViewers, msg.Creator, fullMerklePath, owner)
	if !bool {
		return nil, error
	}


	ok, err := notify(k, ctx, msg.EditorsToNotify, msg.NotiForEditors, msg.Creator, fullMerklePath, owner)
	if !ok {
		return nil, err
	}

	return &types.MsgPostFileResponse{Path: fullMerklePath}, nil
}
