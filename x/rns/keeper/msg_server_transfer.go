package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackal-dao/canine/x/rns/types"
)

func (k msgServer) Transfer(goCtx context.Context, msg *types.MsgTransfer) (*types.MsgTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, _ := sdk.AccAddressFromBech32(msg.Creator)

	whois, isFound := k.GetNames(ctx, msg.Name)

	block_height := ctx.BlockHeight()

	if isFound {
		expires, _ := sdk.NewIntFromString(whois.Expires)

		if block_height > expires.Int64() {
			return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Name does not exist or has expired.")
		}

		if whois.Value != owner.String() {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "You are not the owner of that name.")
		}

		// Create an updated whois record
		newWhois := types.Names{
			Index:   msg.Name,
			Name:    msg.Name,
			Expires: whois.Expires,
			Value:   msg.Reciever,
			Data:    "{}",
		}
		// Write whois information to the store
		k.SetNames(ctx, newWhois)

	} else {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Name does not exist or has expired.")
	}

	return &types.MsgTransferResponse{}, nil
}
