package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/jackalLabs/canine-chain/v4/x/rns/types"
)

func (k msgServer) List(goCtx context.Context, msg *types.MsgList) (*types.MsgListResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	mname := strings.ToLower(msg.Name)

	_, found := k.GetForsale(ctx, mname)

	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Name already listed.")
	}

	n, tld, err := GetNameAndTLD(mname)
	if err != nil {
		return nil, err
	}

	name, nfound := k.GetNames(ctx, n, tld)

	if !nfound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Name does not exist or has expired.")
	}

	if name.Value != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "You do not own this name.")
	}

	blockHeight := ctx.BlockHeight()

	if name.Locked > blockHeight {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "cannot transfer free name")
	}

	if blockHeight > name.Expires {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Name does not exist or has expired.")
	}

	newsale := types.Forsale{
		Name:  mname,
		Price: msg.Price.String(),
		Owner: msg.Creator,
	}

	k.SetForsale(ctx, newsale)

	return &types.MsgListResponse{}, nil
}
