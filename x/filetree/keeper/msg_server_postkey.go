package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/filetree/types"
)

func (k msgServer) PostKey(goCtx context.Context, msg *types.MsgPostKey) (*types.MsgPostKeyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pubKey := types.Pubkey{
		Address: msg.Creator,
		Key:     msg.Key,
	}
	k.SetPubkey(ctx, pubKey)

	return &types.MsgPostKeyResponse{}, nil
}
