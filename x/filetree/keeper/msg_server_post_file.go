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
		Contents:       msg.Contents,
		Owner:          owner,
		ViewingAccess:  msg.Viewers,
		EditAccess:     msg.Editors,
		Address:        fullMerklePath,
		TrackingNumber: msg.TrackingNumber,
	}

	updatedTrackingNumber := msg.TrackingNumber + 1

	//need to double check this number
	if msg.TrackingNumber == 18446744073709551615 {
		updatedTrackingNumber = 0
		k.SetTracker(ctx, types.Tracker{
			TrackingNumber: uint64(updatedTrackingNumber),
		})
	} else {
		k.SetTracker(ctx, types.Tracker{
			TrackingNumber: uint64(updatedTrackingNumber),
		})
	}

	k.SetFiles(ctx, file)

	return &types.MsgPostFileResponse{Path: fullMerklePath}, nil
}
