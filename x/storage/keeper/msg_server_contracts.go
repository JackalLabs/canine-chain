package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackal-dao/canine/x/storage/types"
)

func (k msgServer) CreateContracts(goCtx context.Context, msg *types.MsgCreateContracts) (*types.MsgCreateContractsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetContracts(
		ctx,
		msg.Cid,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var contracts = types.Contracts{
		Creator:    msg.Creator,
		Cid:        msg.Cid,
		Priceamt:   msg.Priceamt,
		Pricedenom: msg.Pricedenom,
		Merkle:     msg.Merkle,
		Signee:     msg.Signee,
		Duration:   msg.Duration,
		Filesize:   msg.Filesize,
		Fid:        msg.Fid,
	}

	k.SetContracts(
		ctx,
		contracts,
	)
	return &types.MsgCreateContractsResponse{}, nil
}

func (k msgServer) UpdateContracts(goCtx context.Context, msg *types.MsgUpdateContracts) (*types.MsgUpdateContractsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetContracts(
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

	var contracts = types.Contracts{
		Creator:  msg.Creator,
		Cid:      msg.Cid,
		Merkle:   msg.Merkle,
		Signee:   msg.Signee,
		Duration: msg.Duration,
		Filesize: msg.Filesize,
		Fid:      msg.Fid,
	}

	k.SetContracts(ctx, contracts)

	return &types.MsgUpdateContractsResponse{}, nil
}

func (k msgServer) DeleteContracts(goCtx context.Context, msg *types.MsgDeleteContracts) (*types.MsgDeleteContractsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetContracts(
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

	k.RemoveContracts(
		ctx,
		msg.Cid,
	)

	return &types.MsgDeleteContractsResponse{}, nil
}
