package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/v3/x/notifications/types"
)

func (k msgServer) CreateNotification(goCtx context.Context, msg *types.MsgCreateNotification) (*types.MsgCreateNotificationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender := msg.Creator
	owner := msg.To

	address, err := k.rns.Resolve(ctx, owner)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "cannot parse address from message")
	}

	if k.IsBlocked(ctx, address.String(), sender) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "you are blocked from sending this user notifications")
	}

	noti := types.Notification{
		To:       address.String(),
		From:     sender,
		Time:     ctx.BlockTime(),
		Contents: msg.Contents,
	}

	k.SetNotification(ctx, noti)

	return &types.MsgCreateNotificationResponse{}, nil
}
