package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackal-dao/canine/x/jklmining/types"
)

func (k msgServer) CreateMiners(goCtx context.Context, msg *types.MsgCreateMiners) (*types.MsgCreateMinersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetMiners(
		ctx,
		msg.Address,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var miners = types.Miners{
		Creator: msg.Creator,
		Address: msg.Address,
		Ip:      msg.Ip,
	}

	k.SetMiners(
		ctx,
		miners,
	)
	return &types.MsgCreateMinersResponse{}, nil
}

func (k msgServer) UpdateMiners(goCtx context.Context, msg *types.MsgUpdateMiners) (*types.MsgUpdateMinersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetMiners(
		ctx,
		msg.Address,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var miners = types.Miners{
		Creator: msg.Creator,
		Address: msg.Address,
		Ip:      msg.Ip,
	}

	k.SetMiners(ctx, miners)

	return &types.MsgUpdateMinersResponse{}, nil
}

func (k msgServer) DeleteMiners(goCtx context.Context, msg *types.MsgDeleteMiners) (*types.MsgDeleteMinersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetMiners(
		ctx,
		msg.Address,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveMiners(
		ctx,
		msg.Address,
	)

	return &types.MsgDeleteMinersResponse{}, nil
}
