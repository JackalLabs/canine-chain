package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

func (k msgServer) SetProviderKeybase(goCtx context.Context, msg *types.MsgSetProviderKeybase) (*types.MsgSetProviderKeybaseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	provider, found := k.GetProviders(ctx, msg.Creator)

	if !found {
		return nil, types.ErrProviderNotFound
	}

	provider.KeybaseIdentity = msg.Keybase

	k.SetProviders(ctx, provider)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeJackalMessage,
			sdk.NewAttribute(types.AttributeKeySigner, msg.Creator),
		),
	)

	return &types.MsgSetProviderKeybaseResponse{}, nil
}
