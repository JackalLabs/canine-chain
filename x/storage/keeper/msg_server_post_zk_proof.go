package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/zk"
	merkletree "github.com/wealdtech/go-merkletree"
	"github.com/wealdtech/go-merkletree/sha3"

	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func (k msgServer) PostZKProof(goCtx context.Context, msg *types.MsgPostZKProof) (*types.MsgPostZKProofResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	meter := ctx.GasMeter()
	usedGas := meter.GasConsumed()

	contract, found := k.GetActiveDeals(ctx, msg.Cid)
	if !found {
		ctx.Logger().Debug("%s, %s\n", "Contract not found", msg.Cid)
		return &types.MsgPostZKProofResponse{Success: false, ErrorMessage: fmt.Sprintf("contract not found: %s", msg.Cid)}, nil
	}

	ctx.Logger().Debug("Contract that was found: \n%v\n", contract)

	nn, ok := sdk.NewIntFromString(contract.Blocktoprove)
	if !ok {
		return &types.MsgPostZKProofResponse{Success: false, ErrorMessage: "cannot parse block to prove"}, nil
	}
	num := nn.Int64()

	wp, err := zk.Decode(&msg.Package)
	if err != nil {
		return &types.MsgPostZKProofResponse{Success: false, ErrorMessage: err.Error()}, nil
	}

	data, err := base64.StdEncoding.DecodeString(msg.Hash)
	if err != nil {
		return &types.MsgPostZKProofResponse{Success: false, ErrorMessage: err.Error()}, nil
	}

	hashVerified, err := zk.VerifyHash(wp, data)
	if err != nil {
		return &types.MsgPostZKProofResponse{Success: false, ErrorMessage: err.Error()}, nil
	}

	if !hashVerified {
		return &types.MsgPostZKProofResponse{Success: false, ErrorMessage: "verification failed"}, nil
	}

	h := sha256.New()
	_, err = io.WriteString(h, fmt.Sprintf("%d%x", num, data))
	if err != nil {
		return &types.MsgPostZKProofResponse{Success: false, ErrorMessage: err.Error()}, nil
	}
	hashName := h.Sum(nil)

	var proof merkletree.Proof

	err = json.Unmarshal([]byte(msg.Hashlist), &proof)
	if err != nil {
		ctx.Logger().Debug("%v\n", err)
		return &types.MsgPostZKProofResponse{Success: false, ErrorMessage: err.Error()}, nil
	}

	ctx.Logger().Debug("proof: %v\n", proof)

	m, err := hex.DecodeString(contract.Merkle)
	if err != nil {
		ctx.Logger().Error("%v\n", err)
		return &types.MsgPostZKProofResponse{Success: false, ErrorMessage: err.Error()}, nil
	}
	verified, err := merkletree.VerifyProofUsing(hashName, false, &proof, [][]byte{m}, sha3.New512())
	if err != nil {
		ctx.Logger().Error("%v\n", err)
		return &types.MsgPostZKProofResponse{Success: false, ErrorMessage: err.Error()}, nil
	}

	if !verified {
		ctx.Logger().Debug("%s\n", "Cannot verify")
		return &types.MsgPostZKProofResponse{Success: false, ErrorMessage: "cannot verify proof"}, nil
	}

	if contract.Proofverified == "true" {
		meter.RefundGas(meter.GasConsumed()-usedGas, "successful proof refund")
		return &types.MsgPostZKProofResponse{Success: false, ErrorMessage: "proof already verified"}, nil
	}

	contract.Proofverified = "true"
	k.SetActiveDeals(ctx, contract)

	meter.RefundGas(meter.GasConsumed()-usedGas, "successful proof refund")

	return &types.MsgPostZKProofResponse{Success: true, ErrorMessage: ""}, nil
}
