package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackal-dao/canine/x/storage/types"
)

func (k msgServer) CreateProofs(goCtx context.Context, msg *types.MsgCreateProofs) (*types.MsgCreateProofsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetProofs(
		ctx,
		msg.Cid,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	proofs := types.Proofs{
		Creator: msg.Creator,
		Cid:     msg.Cid,
		Item:    msg.Item,
		Hashes:  msg.Hashes,
	}

	k.SetProofs(
		ctx,
		proofs,
	)
	return &types.MsgCreateProofsResponse{}, nil
}

func (k msgServer) UpdateProofs(goCtx context.Context, msg *types.MsgUpdateProofs) (*types.MsgUpdateProofsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetProofs(
		ctx,
		msg.Cid,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	proofs := types.Proofs{
		Creator: msg.Creator,
		Cid:     msg.Cid,
		Item:    msg.Item,
		Hashes:  msg.Hashes,
	}

	k.SetProofs(ctx, proofs)

	return &types.MsgUpdateProofsResponse{}, nil
}

func (k msgServer) DeleteProofs(goCtx context.Context, msg *types.MsgDeleteProofs) (*types.MsgDeleteProofsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetProofs(
		ctx,
		msg.Cid,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveProofs(
		ctx,
		msg.Cid,
	)

	return &types.MsgDeleteProofsResponse{}, nil
}
