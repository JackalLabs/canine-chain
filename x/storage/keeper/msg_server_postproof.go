package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
)

func (k msgServer) PostProof(goCtx context.Context, msg *types.MsgPostProof) (*types.MsgPostProofResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	f, found := k.GetFile(ctx, msg.Merkle, msg.Owner, msg.Start)
	if !found {
		s := fmt.Sprintf("contract not found: %x/%s/%d", msg.Merkle, msg.Owner, msg.Start)
		ctx.Logger().Debug(s)
		return &types.MsgPostProofResponse{Success: false, ErrorMessage: s}, nil
	}

	file := &f

	prover := msg.Creator

	var proof *types.FileProof

	if len(file.Proofs) == int(file.MaxProofs) {
		var err error
		proof, err = file.GetProver(ctx, k, prover)
		if err != nil {
			return &types.MsgPostProofResponse{Success: false, ErrorMessage: err.Error()}, nil
		}
	} else {
		if file.ContainsProver(prover) {
			var err error
			proof, err = file.GetProver(ctx, k, prover)
			if err != nil {
				return &types.MsgPostProofResponse{Success: false, ErrorMessage: err.Error()}, nil
			}
		} else {
			proof = file.AddProver(ctx, k, prover)
		}
	}

	if msg.ToProve != proof.ChunkToProve {
		e := fmt.Errorf("wrong chunk to prove for %x. Was %d should be %d", file.Merkle, msg.ToProve, proof.ChunkToProve)
		ctx.Logger().Info(e.Error())
		return &types.MsgPostProofResponse{Success: false, ErrorMessage: e.Error()}, nil
	}

	chunkSize := k.GetParams(ctx).ChunkSize

	if file.ProvenThisBlock(ctx.BlockHeight(), proof.LastProven) {
		ctx.Logger().Info("file was already proven")
	}

	err := file.Prove(ctx, proof, msg.HashList, msg.Item, chunkSize)
	if err != nil {
		e := sdkerrors.Wrapf(err, "cannot verify %x against %x", msg.Item, file.Merkle)
		ctx.Logger().Info(e.Error())
		return &types.MsgPostProofResponse{Success: false, ErrorMessage: e.Error()}, nil
	}

	k.SetProof(ctx, *proof)

	tracker, found := k.GetRewardTracker(ctx, prover) // increase the file trackers size
	if !found {
		tracker = types.RewardTracker{
			Provider: prover,
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
			sdk.NewAttribute(types.AttributeKeySigner, msg.Creator),
		),
	)

	return &types.MsgPostProofResponse{Success: true, ErrorMessage: ""}, nil
}

func (k msgServer) PostProofFor(goCtx context.Context, msg *types.MsgPostProofFor) (*types.MsgPostProofResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	f, found := k.GetFile(ctx, msg.Merkle, msg.Owner, msg.Start)
	if !found {
		err := sdkerrors.Wrapf(types.ErrDealNotFound, "contract not found: %x/%s/%d", msg.Merkle, msg.Owner, msg.Start)
		ctx.Logger().Debug(err.Error())
		return &types.MsgPostProofResponse{Success: false, ErrorMessage: err.Error()}, nil
	}

	file := &f

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

	prover := msg.Provider

	var proof *types.FileProof

	if len(file.Proofs) == int(file.MaxProofs) {
		var err error
		proof, err = file.GetProver(ctx, k, prover)
		if err != nil {
			return &types.MsgPostProofResponse{Success: false, ErrorMessage: err.Error()}, nil
		}
	} else {
		if file.ContainsProver(prover) {
			var err error
			proof, err = file.GetProver(ctx, k, prover)
			if err != nil {
				return &types.MsgPostProofResponse{Success: false, ErrorMessage: err.Error()}, nil
			}
		} else {
			proof = file.AddProver(ctx, k, prover)
		}
	}

	if msg.ToProve != proof.ChunkToProve {
		err := sdkerrors.Wrapf(types.ErrBadProofInput, "wrong chunk to prove for %x. Was %d should be %d", file.Merkle, msg.ToProve, proof.ChunkToProve)
		ctx.Logger().Info(err.Error())
		return &types.MsgPostProofResponse{Success: false, ErrorMessage: err.Error()}, nil
	}

	chunkSize := k.GetParams(ctx).ChunkSize

	if file.ProvenThisBlock(ctx.BlockHeight(), proof.LastProven) {
		ctx.Logger().Info("file was already proven")
	}

	err := file.Prove(ctx, proof, msg.HashList, msg.Item, chunkSize)
	if err != nil {
		e := sdkerrors.Wrapf(err, "cannot verify %x against %x", msg.Item, file.Merkle)
		ctx.Logger().Info(e.Error())
		return &types.MsgPostProofResponse{Success: false, ErrorMessage: e.Error()}, nil
	}

	k.SetProof(ctx, *proof)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeJackalMessage,
			sdk.NewAttribute(types.AttributeKeySigner, msg.Creator),
		),
	)

	return &types.MsgPostProofResponse{Success: true, ErrorMessage: ""}, nil
}
