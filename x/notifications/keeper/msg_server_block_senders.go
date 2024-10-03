package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/v4/x/notifications/types"
)

func (k msgServer) BlockSenders(goCtx context.Context, msg *types.MsgBlockSenders) (*types.MsgBlockSendersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	for _, toBlock := range msg.ToBlock {

		address, err := k.rns.Resolve(ctx, toBlock)
		if err != nil {
			return nil, sdkerrors.Wrapf(err, "cannot parse address %s from message", toBlock)
		}

		b := types.Block{
			Address:        msg.Creator,
			BlockedAddress: address.String(),
		}

		k.SetBlock(ctx, b)

	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventBlockSenders,
			sdk.NewAttribute(types.AttributeSigner, msg.Creator),
		),
	)

	return &types.MsgBlockSendersResponse{}, nil
}
