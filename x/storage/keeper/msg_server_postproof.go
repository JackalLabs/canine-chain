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

	proofSize := k.GetParams(ctx).ChunkSize

	file, found := k.GetFile(ctx, msg.Merkle, msg.Owner, msg.Start)
	if !found {
		s := fmt.Sprintf("contract not found: %x/%s/%d", msg.Merkle, msg.Owner, msg.Start)
		ctx.Logger().Debug(s)
		return &types.MsgPostProofResponse{Success: false, ErrorMessage: s}, nil
	}

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

	err := file.Prove(ctx, k, msg.Creator, msg.HashList, proof.ChunkToProve, msg.Item, proofSize)
	if err != nil {
		e := sdkerrors.Wrapf(err, "cannot verify %x against %x", msg.Item, file.Merkle)
		ctx.Logger().Debug(e.Error())
		return &types.MsgPostProofResponse{Success: false, ErrorMessage: e.Error()}, nil
	}

	k.SetFile(ctx, file)

	return &types.MsgPostProofResponse{Success: true, ErrorMessage: ""}, nil
}
