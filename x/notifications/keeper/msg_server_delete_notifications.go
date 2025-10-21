package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v5/x/notifications/types"
)

// DeleteNotification deletes a given message
func (k msgServer) DeleteNotification(goCtx context.Context, msg *types.MsgDeleteNotification) (*types.MsgDeleteNotificationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.RemoveNotification(ctx, msg.Creator, msg.From, msg.Time)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventRemoveNotification,
			sdk.NewAttribute(types.AttributeSigner, msg.Creator),
		),
	)

	return &types.MsgDeleteNotificationResponse{}, nil
}
