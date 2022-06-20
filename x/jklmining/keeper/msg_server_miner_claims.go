package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackal-dao/canine/x/jklmining/types"
)

func (k msgServer) CreateMinerClaims(goCtx context.Context, msg *types.MsgCreateMinerClaims) (*types.MsgCreateMinerClaimsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetMinerClaims(
		ctx,
		msg.Hash,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var minerClaims = types.MinerClaims{
		Creator: msg.Creator,
		Hash:    msg.Hash,
	}

	k.SetMinerClaims(
		ctx,
		minerClaims,
	)
	return &types.MsgCreateMinerClaimsResponse{}, nil
}

func (k msgServer) UpdateMinerClaims(goCtx context.Context, msg *types.MsgUpdateMinerClaims) (*types.MsgUpdateMinerClaimsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetMinerClaims(
		ctx,
		msg.Hash,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var minerClaims = types.MinerClaims{
		Creator: msg.Creator,
		Hash:    msg.Hash,
	}

	k.SetMinerClaims(ctx, minerClaims)

	return &types.MsgUpdateMinerClaimsResponse{}, nil
}

func (k msgServer) DeleteMinerClaims(goCtx context.Context, msg *types.MsgDeleteMinerClaims) (*types.MsgDeleteMinerClaimsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetMinerClaims(
		ctx,
		msg.Hash,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveMinerClaims(
		ctx,
		msg.Hash,
	)

	return &types.MsgDeleteMinerClaimsResponse{}, nil
}
