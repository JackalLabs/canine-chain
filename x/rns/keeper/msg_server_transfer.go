package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackal-dao/canine/x/rns/types"
)

func (k msgServer) Transfer(goCtx context.Context, msg *types.MsgTransfer) (*types.MsgTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, _ := sdk.AccAddressFromBech32(msg.Creator)

	name, tld, err := getNameAndTLD(msg.Name)
	if err != nil {
		return nil, err
	}

	whois, isFound := k.GetNames(ctx, name, tld)

	admin := whois.Value

	blockHeight := ctx.BlockHeight()

	if isFound {

		if blockHeight > whois.Expires {
			return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Name does not exist or has expired.")
		}

		if admin != sender.String() {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "You are not the owner of that name.")
		}

		if whois.Locked > blockHeight {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "cannot transfer free name")
		}

		whois.Data = "{}"
		whois.Value = msg.Reciever

		// Write whois information to the store
		k.SetNames(ctx, whois)

	} else {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Name does not exist or has expired.")
	}

	return &types.MsgTransferResponse{}, nil
}
