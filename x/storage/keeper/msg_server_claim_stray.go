package keeper

import (
	"context"

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

	provider, found := k.GetProviders(ctx, msg.ForAddress)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "not a provider")
	}

	for _, l := range ls {
		if l == provider.Ip {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "cannot claim a stray you own.")
		}
	}

	size := sdk.NewInt(int64(stray.Size()))

	pieces := size.Quo(sdk.NewInt(k.GetParams(ctx).ChunkSize))

	var pieceToStart int64

	if !pieces.IsZero() {
		pieceToStart = ctx.BlockHeight() % pieces.Int64()
	}

	if msg.ForAddress != msg.Creator {
		found := false
		for _, claimer := range provider.AuthClaimers {
			if claimer == msg.Creator {
				found = true
			}
		}
		if !found {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "cannot claim a stray for someone else")
		}

	}

	deal := types.ActiveDealsV2{
		Cid:           stray.Cid,
		Signer:        stray.Signer,
		Provider:      msg.ForAddress,
		StartBlock:    ctx.BlockHeight(),
		EndBlock:      stray.End,
		FileSize:      stray.FileSize,
		ProofVerified: false,
		BlockToProve:  pieceToStart,
		Creator:       msg.Creator,
		ProofsMissed:  0,
		Merkle:        stray.Merkle,
		Fid:           stray.Fid,
	}

	k.SetActiveDeals(ctx, deal)

	k.RemoveStrays(ctx, stray.Cid)

	return &types.MsgClaimStrayResponse{}, nil
}
