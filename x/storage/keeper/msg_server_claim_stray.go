package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackal-dao/canine/x/storage/types"
)

func (k msgServer) ClaimStray(goCtx context.Context, msg *types.MsgClaimStray) (*types.MsgClaimStrayResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	stray, ok := k.GetStrays(ctx, msg.Cid)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "stray contract either no longer is stray, or has been removed by the user")
	}

	deal := types.ActiveDeals{
		Cid:           stray.Cid,
		Signee:        stray.Signee,
		Provider:      msg.Creator,
		Startblock:    fmt.Sprintf("%d", ctx.BlockHeight()),
		Endblock:      fmt.Sprintf("%d", ctx.BlockHeight()),
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
