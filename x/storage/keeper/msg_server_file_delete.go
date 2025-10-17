package keeper

import (
	"context"
	"encoding/hex"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
)

func (k msgServer) DeleteFile(goCtx context.Context, msg *types.MsgDeleteFile) (*types.MsgDeleteFileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	file, found := k.GetFile(ctx, msg.Merkle, msg.Creator, msg.Start)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "file not found")
	}

	if file.Expires != 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "can not delete files before they expire")
	}

	paymentInfo, found := k.GetStoragePaymentInfo(ctx, msg.Creator) // needs payment info if they're gonna delete a file
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "payment info not found")
	}

	paymentInfo.SpaceUsed -= file.FileSize // remove usage from payment info
	if paymentInfo.SpaceUsed < 0 {
		paymentInfo.SpaceUsed = 0 // cap at 0
	}

	k.SetStoragePaymentInfo(ctx, paymentInfo)

	k.RemoveFile(ctx, msg.Merkle, msg.Creator, msg.Start)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCancelContract,
			sdk.NewAttribute(types.AttributeKeySigner, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyContract, hex.EncodeToString(msg.Merkle)),
		),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeJackalMessage,
			sdk.NewAttribute(types.AttributeKeySigner, msg.Creator),
		),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return &types.MsgDeleteFileResponse{}, nil
}
