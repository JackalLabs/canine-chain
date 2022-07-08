package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/themarstonconnell/telescope/x/telescope/types"
)

func (k msgServer) List(goCtx context.Context, msg *types.MsgList) (*types.MsgListResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetForsale(ctx, msg.Name)

	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Name already listed.")
	}

	name, nfound := k.GetNames(ctx, msg.Name)

	if !nfound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Name does not exist or has expired.")
	}

	if name.Value != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "You do not own this name.")
	}

	expires, _ := sdk.NewIntFromString(name.Expires)
	block_height := ctx.BlockHeight()

	if block_height > expires.Int64() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Name does not exist or has expired.")
	}

	newsale := types.Forsale{
		Name:  msg.Name,
		Price: msg.Price,
		Owner: msg.Creator,
	}

	k.SetForsale(ctx, newsale)

	return &types.MsgListResponse{}, nil
}
