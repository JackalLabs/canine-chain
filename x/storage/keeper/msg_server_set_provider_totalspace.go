package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

func (k msgServer) SetProviderTotalspace(goCtx context.Context, msg *types.MsgSetProviderTotalspace) (*types.MsgSetProviderTotalspaceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	provider, found := k.GetProviders(ctx, msg.Creator)

	if !found {
		return nil, types.ErrProviderNotFound
	}

	validTotalSpace := isValidTotalSpace(msg.Space)

	if !validTotalSpace {
		return nil, types.ErrNotValidTotalSpace
	}

	provider.Totalspace = msg.Space

	k.SetProviders(ctx, provider)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return &types.MsgSetProviderTotalspaceResponse{}, nil
}

func isValidTotalSpace(totalSpace string) bool {
	var isNumber bool

	if _, err := strconv.Atoi(totalSpace); err == nil {
		isNumber = true
	}
	return isNumber
}
