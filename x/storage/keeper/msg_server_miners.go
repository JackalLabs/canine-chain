package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackal-dao/canine/x/storage/types"
)

func (k msgServer) CreateProviders(goCtx context.Context, msg *types.MsgCreateProviders) (*types.MsgCreateProvidersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetProviders(
		ctx,
		msg.Address,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	providers := types.Providers{
		Creator:    msg.Creator,
		Address:    msg.Address,
		Ip:         msg.Ip,
		Totalspace: msg.Totalspace,
	}

	k.SetProviders(
		ctx,
		providers,
	)
	return &types.MsgCreateProvidersResponse{}, nil
}

func (k msgServer) UpdateProviders(goCtx context.Context, msg *types.MsgUpdateProviders) (*types.MsgUpdateProvidersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetProviders(
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

	providers := types.Providers{
		Creator:    msg.Creator,
		Address:    msg.Address,
		Ip:         msg.Ip,
		Totalspace: msg.Totalspace,
	}

	k.SetProviders(ctx, providers)

	return &types.MsgUpdateProvidersResponse{}, nil
}

func (k msgServer) DeleteProviders(goCtx context.Context, msg *types.MsgDeleteProviders) (*types.MsgDeleteProvidersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetProviders(
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

	k.RemoveProviders(
		ctx,
		msg.Address,
	)

	return &types.MsgDeleteProvidersResponse{}, nil
}
