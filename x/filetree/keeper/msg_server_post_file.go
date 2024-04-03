package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/filetree/types"
)

func (k msgServer) PostFile(goCtx context.Context, msg *types.MsgPostFile) (*types.MsgPostFileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	parentOwnerString := MakeOwnerAddress(msg.HashParent, msg.Account)

	parentFile, found := k.GetFiles(ctx, msg.HashParent, parentOwnerString)
	if !found {
		return nil, types.ErrParentFileNotFound
	}

	hasEdit, err := HasEditAccess(parentFile, msg.Creator)
	if err != nil {
		// Error raised when json unmarshalling has failed
		return nil, err
	}

	if !hasEdit {
		return nil, types.ErrCannotWrite
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

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypePostFile,
			sdk.NewAttribute(types.AttributeKeySigner, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyFileAddress, fullMerklePath),
			sdk.NewAttribute(types.AttributeKeyOwner, owner),
		),
	)

	return &types.MsgPostFileResponse{Path: fullMerklePath}, nil
}
