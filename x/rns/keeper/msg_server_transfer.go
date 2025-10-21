package keeper

import (
	"context"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/v5/x/rns/types"
)

func (k Keeper) TransferName(ctx sdk.Context, creator string, receiever string, name string) error {
	name = strings.ToLower(name)

	sender, err := sdk.AccAddressFromBech32(creator)
	if err != nil {
		return err
	}

	name, tld, err := GetNameAndTLD(name)
	if err != nil {
		return err
	}

	whois, isFound := k.GetNames(ctx, name, tld)

	admin := whois.Value

	blockHeight := ctx.BlockHeight()

	if !isFound {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "Name does not exist or has expired.")
	}

	if blockHeight > whois.Expires {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "Name does not exist or has expired.")
	}

	if admin != sender.String() {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "You are not the owner of that name.")
	}

	if whois.Locked > blockHeight {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "cannot transfer free name")
	}

	whois.Data = "{}"
	whois.Value = receiever

	// Write whois information to the store
	k.SetNames(ctx, whois)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTransfer,
			sdk.NewAttribute(types.AttributeName, fmt.Sprintf("%s.%s", whois.Name, whois.Tld)),
			sdk.NewAttribute(types.AttributeOwner, creator),
			sdk.NewAttribute(types.AttributeReceiver, receiever),
		),
	)

	return nil
}

func (k msgServer) Transfer(goCtx context.Context, msg *types.MsgTransfer) (*types.MsgTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.TransferName(ctx, msg.Creator, msg.Receiver, msg.Name)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeJackalMessage,
			sdk.NewAttribute(types.AttributeKeySigner, msg.Creator),
		),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTransfer,
			sdk.NewAttribute(types.AttributeKeySigner, msg.Creator),
		),
	)

	return &types.MsgTransferResponse{}, err
}
