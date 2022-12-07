package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func (k msgServer) ClaimStray(goCtx context.Context, msg *types.MsgClaimStray) (*types.MsgClaimStrayResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	stray, ok := k.GetStrays(ctx, msg.Cid)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "stray contract either no longer is stray, or has been removed by the user")
	}

	ls := k.ListFiles(ctx, stray.Fid)

	provider, found := k.GetProviders(ctx, msg.Creator)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "not a provider")
	}

	for _, l := range ls {
		if l == provider.Ip {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "cannot claim a stray you own.")
		}
	}

	deal := types.ActiveDeals{
		Cid:           stray.Cid,
		Signee:        stray.Signee,
		Provider:      msg.Creator,
		Startblock:    fmt.Sprintf("%d", ctx.BlockHeight()),
		Endblock:      "0",
		Filesize:      stray.Filesize,
		Proofverified: "false",
		Blocktoprove:  fmt.Sprintf("%d", ctx.BlockHeight()/1024),
		Creator:       msg.Creator,
		Proofsmissed:  "0",
		Merkle:        stray.Merkle,
		Fid:           stray.Fid,
	}

	k.SetActiveDeals(ctx, deal)

	k.RemoveStrays(ctx, stray.Cid)

	return &types.MsgClaimStrayResponse{}, nil
}
