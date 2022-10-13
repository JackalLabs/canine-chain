package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/storage/types"
)

func (k msgServer) SetProviderTotalspace(goCtx context.Context, msg *types.MsgSetProviderTotalspace) (*types.MsgSetProviderTotalspaceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	provider, found := k.GetProviders(ctx, msg.Creator)

	if !found {
		provider = types.Providers{
			Address:         msg.Creator,
			Ip:              "",
			Totalspace:      "0",
			Creator:         msg.Creator,
			BurnedContracts: "0",
		}
	}

	provider.Totalspace = msg.Space

	k.SetProviders(ctx, provider)

	return &types.MsgSetProviderTotalspaceResponse{}, nil
}
