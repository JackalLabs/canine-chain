package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
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
