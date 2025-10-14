package keeper

import (
	"context"
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
)

func (k Keeper) UpdateProof(ctx sdk.Context, proof *types.FileProof, file *types.UnifiedFile) {
	k.SetProof(ctx, *proof)

	tracker, found := k.GetRewardTracker(ctx, proof.Prover) // increase the file trackers size
	if !found {
		tracker = types.RewardTracker{
			Provider: proof.Prover,
			Size_:    0,
		}
	}
	tracker.Size_ += file.FileSize
	k.SetRewardTracker(ctx, tracker)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeJackalMessage,
			sdk.NewAttribute(types.AttributeKeySigner, proof.Prover),
		),
	)
}

func (k Keeper) postProof(ctx sdk.Context,
	prover string,
	item []byte,
	hashList []byte,
	merkle []byte,
	owner string,
	start int64,
	toProve int64,
) error {
	f, found := k.GetFile(ctx, merkle, owner, start)
	if !found {
		s := fmt.Sprintf("contract not found: %x/%s/%d", merkle, owner, start)
		ctx.Logger().Debug(s)
		return errors.New(s)
	}
	file := &f

	var proof *types.FileProof

	if len(file.Proofs) == int(file.MaxProofs) {
		var err error
		proof, err = file.GetProver(ctx, k, prover)
		if err != nil {
			return sdkerrors.Wrap(err, "this is not your file")
		}
	} else {
		if file.ContainsProver(prover) {
			var err error
			proof, err = file.GetProver(ctx, k, prover)
			if err != nil {
				return sdkerrors.Wrap(err, "you were supposed to have a proof but don't")
			}
		} else {
			proof = file.AddProver(ctx, k, prover)
		}
	}

	if toProve != proof.ChunkToProve {
		e := fmt.Errorf("wrong chunk to prove for %x. Was %d should be %d", file.Merkle, toProve, proof.ChunkToProve)
		ctx.Logger().Info(e.Error())
		return e
	}

	chunkSize := k.GetParams(ctx).ChunkSize

	if file.ProvenThisBlock(ctx.BlockHeight(), proof.LastProven) {
		ctx.Logger().Info("file was already proven")
	}

	err := file.Prove(ctx, proof, hashList, item, chunkSize)
	if err != nil {
		e := sdkerrors.Wrapf(err, "cannot verify %x against %x", item, file.Merkle)
		ctx.Logger().Info(e.Error())
		return e
	}

	k.UpdateProof(ctx, proof, file)

	return nil
}

func (k msgServer) PostProof(goCtx context.Context, msg *types.MsgPostProof) (*types.MsgPostProofResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.postProof(ctx, msg.Creator, msg.Item, msg.HashList, msg.Merkle, msg.Owner, msg.Start, msg.ToProve)
	if err != nil {
		return &types.MsgPostProofResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}, nil
	}

	return &types.MsgPostProofResponse{
		Success:      true,
		ErrorMessage: "",
	}, nil
}

func (k msgServer) PostProofFor(goCtx context.Context, msg *types.MsgPostProofFor) (*types.MsgPostProofResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	provider, found := k.GetProviders(ctx, msg.Provider)
	if !found {
		return nil, types.ErrProviderNotFound
	}

	found = false
	for _, claimer := range provider.AuthClaimers {
		if claimer == msg.Creator {
			found = true
		}
	}
	if !found {
		return nil, types.ErrProviderNotFound
	}

	err := k.postProof(ctx, msg.Provider, msg.Item, msg.HashList, msg.Merkle, msg.Owner, msg.Start, msg.ToProve)
	if err != nil {
		return &types.MsgPostProofResponse{Success: false, ErrorMessage: err.Error()}, nil
	}

	return &types.MsgPostProofResponse{Success: true, ErrorMessage: ""}, nil
}
