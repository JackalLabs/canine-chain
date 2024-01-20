package keeper

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/filetree/types"
)

func (k msgServer) PostFile(goCtx context.Context, msg *types.MsgPostFile) (*types.MsgPostFileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	parentOwnerString := MakeOwnerAddress(msg.HashParent, msg.Account)

	parentFile, found := k.GetFiles(ctx, msg.HashParent, parentOwnerString)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrParentFileNotFound, "cannot find %s", msg.HashParent)
	}

	hasEdit, err := HasEditAccess(parentFile, msg.Creator)
	if err != nil {
		// Error raised when json unmarshalling has failed
		return nil, sdkerrors.Wrapf(err, "cannot check for edit access")
	}

	if !hasEdit {
		return nil, sdkerrors.Wrapf(types.ErrCannotWrite, "does not have edit access")
	}

	// Make the full path
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

	return &types.MsgPostFileResponse{Path: fullMerklePath}, nil
}
