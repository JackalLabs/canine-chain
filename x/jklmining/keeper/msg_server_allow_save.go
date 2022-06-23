package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/jackal-dao/canine/x/jklmining/types"
)

func (k msgServer) AllowSave(goCtx context.Context, msg *types.MsgAllowSave) (*types.MsgAllowSaveResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	jak := k.Keeper.jklAccountsKeeper

	jaccount, found := jak.GetAccounts(ctx, msg.Creator)

	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "account not set up")
	}

	av, _ := sdk.NewIntFromString(jaccount.Available)
	us, _ := sdk.NewIntFromString(jaccount.Used)
	sz, _ := sdk.NewIntFromString(msg.Size_)

	if av.Int64()-us.Int64() < sz.Int64() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "not enough space on account")
	}

	// Check if the value already exists
	_, isFound := k.GetSaveRequests(
		ctx,
		msg.Passkey,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var saveRequests = types.SaveRequests{
		Creator:  msg.Creator,
		Index:    msg.Passkey,
		Size_:    msg.Size_,
		Approved: "false",
	}

	k.SetSaveRequests(
		ctx,
		saveRequests,
	)

	return &types.MsgAllowSaveResponse{}, nil
}
