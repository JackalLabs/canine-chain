package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/jackalLabs/canine-chain/v3/x/rns/types"
)

func (k msgServer) DelRecord(goCtx context.Context, msg *types.MsgDelRecord) (*types.MsgDelRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	mname := strings.ToLower(msg.Name)

	n, tld, err := GetNameAndTLD(mname)
	if err != nil {
		return nil, err
	}

	sub, n, hasSub := GetSubdomain(n)

	if !hasSub {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "name does not contain records")
	}

	val, found := k.GetNames(ctx, n, tld)

	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "cannot find name")
	}

	if ctx.BlockHeight() > val.Expires {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "name does not exist or has expired")
	}

	if msg.Creator != val.Value {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "you do not own this name")
	}

	removed := false
	dms := []*types.Names{}
	for _, domain := range val.Subdomains {
		if domain.Name != sub {
			dms = append(dms, domain)
			continue
		}
		removed = true
	}
	if !removed {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "record does not exist for this name")
	}

	val.Subdomains = dms
	k.SetNames(ctx, val)

	return &types.MsgDelRecordResponse{}, nil
}
