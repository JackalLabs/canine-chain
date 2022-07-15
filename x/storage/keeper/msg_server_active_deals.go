package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackal-dao/canine/x/storage/types"
)

func (k msgServer) CreateActiveDeals(goCtx context.Context, msg *types.MsgCreateActiveDeals) (*types.MsgCreateActiveDealsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetActiveDeals(
		ctx,
		msg.Cid,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var activeDeals = types.ActiveDeals{
		Creator:       msg.Creator,
		Cid:           msg.Cid,
		Signee:        msg.Signee,
		Miner:         msg.Miner,
		Startblock:    msg.Startblock,
		Endblock:      msg.Endblock,
		Filesize:      msg.Filesize,
		Proofverified: msg.Proofverified,
		Proofsmissed:  msg.Proofsmissed,
		Blocktoprove:  msg.Blocktoprove,
	}

	k.SetActiveDeals(
		ctx,
		activeDeals,
	)
	return &types.MsgCreateActiveDealsResponse{}, nil
}

func (k msgServer) UpdateActiveDeals(goCtx context.Context, msg *types.MsgUpdateActiveDeals) (*types.MsgUpdateActiveDealsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetActiveDeals(
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

	var activeDeals = types.ActiveDeals{
		Creator:       msg.Creator,
		Cid:           msg.Cid,
		Signee:        msg.Signee,
		Miner:         msg.Miner,
		Startblock:    msg.Startblock,
		Endblock:      msg.Endblock,
		Filesize:      msg.Filesize,
		Proofverified: msg.Proofverified,
		Proofsmissed:  msg.Proofsmissed,
		Blocktoprove:  msg.Blocktoprove,
	}

	k.SetActiveDeals(ctx, activeDeals)

	return &types.MsgUpdateActiveDealsResponse{}, nil
}

func (k msgServer) DeleteActiveDeals(goCtx context.Context, msg *types.MsgDeleteActiveDeals) (*types.MsgDeleteActiveDealsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetActiveDeals(
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

	k.RemoveActiveDeals(
		ctx,
		msg.Cid,
	)

	return &types.MsgDeleteActiveDealsResponse{}, nil
}
