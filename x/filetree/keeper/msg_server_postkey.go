package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/x/filetree/types"
)

func (k msgServer) PostKey(goCtx context.Context, msg *types.MsgPostKey) (*types.MsgPostKeyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pubKey := types.Pubkey{
		Address: msg.Creator,
		Key:     msg.Key,
	}
	k.SetPubkey(ctx, pubKey)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypePostKey,
			sdk.NewAttribute(types.AttributeKeySigner, msg.Creator),
		),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeJackalMessage,
			sdk.NewAttribute(types.AttributeKeySigner, msg.Creator),
		),
	)

	return &types.MsgPostKeyResponse{}, nil
}
