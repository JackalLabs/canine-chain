package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func (k msgServer) InitProvider(goCtx context.Context, msg *types.MsgInitProvider) (*types.MsgInitProviderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	provider := types.Providers{
		Address:         msg.Creator,
		Ip:              msg.Ip,
		Totalspace:      msg.Totalspace,
		Creator:         msg.Creator,
		BurnedContracts: "0",
		KeybaseIdentity: msg.Keybase,
	}

	k.SetProviders(ctx, provider)

	return &types.MsgInitProviderResponse{}, nil
}
