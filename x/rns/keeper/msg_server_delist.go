package keeper

import (
	"context"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/x/rns/types"
)

func (k msgServer) Delist(goCtx context.Context, msg *types.MsgDelist) (*types.MsgDelistResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	mname := strings.ToLower(msg.Name)

	sale, found := k.GetForsale(ctx, mname)

	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Name isn't listed.")
	}

	n, tld, err := GetNameAndTLD(mname)
	if err != nil {
		return nil, err
	}

	name, nfound := k.GetNames(ctx, n, tld)

	if !nfound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Name does not exist or has expired.")
	}

	if sale.Owner != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "You do not own this listing.")
	}

	if name.Value != sale.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "This listing has expired.")
	}

	k.RemoveForsale(ctx, mname)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeJackalMessage,
			sdk.NewAttribute(types.AttributeKeySigner, msg.Creator),
		),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return &types.MsgDelistResponse{}, nil
}
