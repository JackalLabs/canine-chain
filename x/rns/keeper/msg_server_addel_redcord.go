package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/jackal-dao/canine/x/rns/types"
)

func (k msgServer) DelRecord(goCtx context.Context, msg *types.MsgDelRecord) (*types.MsgDelRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	n, tld, err := getNameAndTLD(msg.Name)
	if err != nil {
		return nil, err
	}

	sub, n, hasSub := getSubdomain(n)

	if !hasSub {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "name does not contain records")
	}

	val, found := k.GetNames(ctx, n, tld)

	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "cannot find name")

	}

	if msg.Creator != val.Value {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "you do not own this name")
	}

	dms := []*types.Names{}
	for _, domain := range val.Subdomains {
		if domain.Name != sub {
			dms = append(dms, domain)
		}
	}

	val.Subdomains = dms
	k.SetNames(ctx, val)

	return &types.MsgDelRecordResponse{}, nil
}
