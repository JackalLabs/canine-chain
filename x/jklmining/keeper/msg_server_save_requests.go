package keeper

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackal-dao/canine/x/jklmining/types"
)

func (k msgServer) CreateSaveRequests(goCtx context.Context, msg *types.MsgCreateSaveRequests) (*types.MsgCreateSaveRequestsResponse, error) {
	return nil, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "command not supported")
	/**
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetSaveRequests(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var saveRequests = types.SaveRequests{
		Creator:  msg.Creator,
		Index:    msg.Index,
		Size_:    msg.Size_,
		Approved: msg.Approved,
	}

	k.SetSaveRequests(
		ctx,
		saveRequests,
	)
	return &types.MsgCreateSaveRequestsResponse{}, nil
	*/
}

func (k msgServer) UpdateSaveRequests(goCtx context.Context, msg *types.MsgUpdateSaveRequests) (*types.MsgUpdateSaveRequestsResponse, error) {
	return nil, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "command not supported")
	/**
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetSaveRequests(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var saveRequests = types.SaveRequests{
		Creator:  msg.Creator,
		Index:    msg.Index,
		Size_:    msg.Size_,
		Approved: msg.Approved,
	}

	k.SetSaveRequests(ctx, saveRequests)

	return &types.MsgUpdateSaveRequestsResponse{}, nil
	*/
}

func (k msgServer) DeleteSaveRequests(goCtx context.Context, msg *types.MsgDeleteSaveRequests) (*types.MsgDeleteSaveRequestsResponse, error) {
	return nil, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "command not supported")
	/**
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetSaveRequests(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveSaveRequests(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteSaveRequestsResponse{}, nil
	*/
}
