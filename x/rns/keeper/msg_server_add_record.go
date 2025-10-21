package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/v5/x/rns/types"
)

func (k msgServer) AddRecord(goCtx context.Context, msg *types.MsgAddRecord) (*types.MsgAddRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	mname := strings.ToLower(msg.Name)

	name, tld, err := GetNameAndTLD(mname)
	if err != nil {
		return nil, err
	}

	whois, isFound := k.GetNames(ctx, name, tld)

	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "name does not exist or has expired")
	}

	if ctx.BlockHeight() > whois.Expires {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "name does not exist or has expired")
	}

	if msg.Creator != whois.Value {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "you do not own this name")
	}

	if strings.Contains(msg.Value, ".") {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "cannot have a '.' in a record")
	}

	// checking if the subdomain is already added
	for _, sd := range whois.Subdomains {
		if sd.Name == msg.Record {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "Subdomain already exists")
		}
	}

	// initializing the subdomains
	if whois.Subdomains == nil {
		whois.Subdomains = []*types.Names{}
	}

	// creating a new subdomain type
	record := types.Names{
		Name:       strings.ToLower(msg.Record),
		Expires:    whois.Expires,
		Value:      msg.Value,
		Data:       msg.Data,
		Subdomains: nil,
		Tld:        whois.Tld,
	}

	whois.Subdomains = append(whois.Subdomains, &record)

	k.SetNames(ctx, whois)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventAddRecord,
			sdk.NewAttribute(types.AttributeName, msg.Name),
			sdk.NewAttribute(types.AttributeOwner, msg.Creator),
			sdk.NewAttribute(types.AttributeKeySigner, msg.Creator),
		),
	)

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

	return &types.MsgAddRecordResponse{}, nil
}
