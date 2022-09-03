package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackal-dao/canine/x/rns/types"
)

func (k msgServer) AddRecord(goCtx context.Context, msg *types.MsgAddRecord) (*types.MsgAddRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	name, tld, err := getNameAndTLD(msg.Name)
	if err != nil {
		return nil, err
	}

	whois, isFound := k.GetNames(ctx, name, tld)

	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "name does not exist or has expired")
	}

	expires, ok := sdk.NewIntFromString(whois.Expires)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "cannot parse expiration")
	}

	if ctx.BlockHeight() > expires.Int64() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "name does not exist or has expired")
	}

	if msg.Creator != whois.Value {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "you do not own this name")
	}

	if strings.Contains(msg.Value, ".") {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "cannot have a '.' in a record")
	}

	record := types.Names{
		Name:       msg.Record,
		Expires:    whois.Expires,
		Value:      msg.Value,
		Data:       msg.Data,
		Subdomains: nil,
		Tld:        whois.Tld,
	}

	if whois.Subdomains == nil {
		whois.Subdomains = []*types.Names{}
	}
	whois.Subdomains = append(whois.Subdomains, &record)

	k.SetNames(ctx, whois)

	return &types.MsgAddRecordResponse{}, nil
}
