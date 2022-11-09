package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func (k msgServer) SetProviderTotalspace(goCtx context.Context, msg *types.MsgSetProviderTotalspace) (*types.MsgSetProviderTotalspaceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	provider, found := k.GetProviders(ctx, msg.Creator)

	if !found {
		return nil, types.ErrProviderNotFound
	}

	provider.Totalspace = msg.Space

	k.SetProviders(ctx, provider)

	return &types.MsgSetProviderTotalspaceResponse{}, nil
}
