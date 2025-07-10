package keeper

import (
	"context"
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
)

func (k msgServer) DeleteFile(goCtx context.Context, msg *types.MsgDeleteFile) (*types.MsgDeleteFileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

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
