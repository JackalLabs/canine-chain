package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/filetree/types"
)

func (k msgServer) Postkey(goCtx context.Context, msg *types.MsgPostkey) (*types.MsgPostkeyResponse, error) {
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
			sdk.NewAttribute(types.AttributeKeyKey, msg.Key),
		),
	)

	return &types.MsgPostkeyResponse{}, nil
}
