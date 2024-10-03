package keeper

import (
	"context"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/jackalLabs/canine-chain/v4/x/rns/types"
)

func (k Keeper) UpdateName(ctx sdk.Context, sender string, nm string, data string) error {
	nm = strings.ToLower(nm)
	name, tld, err := GetNameAndTLD(nm)
	if err != nil {
		return err
	}

	whois, isFound := k.GetNames(ctx, name, tld)
	// If a name isn't found in store, error
	if !isFound {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "name does not exist or has expired")
	}

	owner, err := sdk.AccAddressFromBech32(sender)
	if err != nil {
		return sdkerrors.Wrap(err, "cannot parse sender")
	}

	if whois.Value != owner.String() { // error if user doesn't own the name
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "not your name")
	}

	blockHeight := ctx.BlockHeight() // making sure name is still valid
	if blockHeight > whois.Expires {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidHeight, "name is expired")
	}

	whois.Data = data
	// Write whois information to the store
	k.SetNames(ctx, whois)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventUpdate,
			sdk.NewAttribute(types.AttributeName, fmt.Sprintf("%s.%s", whois.Name, whois.Tld)),
			sdk.NewAttribute(types.AttributeOwner, sender),
		),
	)

	return nil
}

func (k msgServer) Update(goCtx context.Context, msg *types.MsgUpdate) (*types.MsgUpdateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Try getting a name from the store

	err := k.UpdateName(ctx, msg.Creator, msg.Name, msg.Data)

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

	return &types.MsgUpdateResponse{}, err
}
