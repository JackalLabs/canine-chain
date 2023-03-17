package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	merkletree "github.com/wealdtech/go-merkletree"
	"github.com/wealdtech/go-merkletree/sha3"

	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func (k msgServer) Postproof(goCtx context.Context, msg *types.MsgPostproof) (*types.MsgPostproofResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	meter := ctx.BlockGasMeter()
	usedGas := meter.GasConsumed()
	var refunded uint64

	hashes := strings.Split(msg.Hashlist, ",")

	contract, found := k.GetActiveDeals(ctx, msg.Cid)
	if !found {
		ctx.Logger().Debug("%s, %s\n", "Contract not found", msg.Cid)
		return &types.MsgPostproofResponse{Success: false, ErrorMessage: fmt.Sprintf("contract not found: %s", msg.Cid)}, nil
	}
	refunded += meter.GasConsumed() - usedGas
	meter.RefundGas(meter.GasConsumed()-usedGas, "gas limiting")

	ctx.Logger().Debug("Contract that was found: \n%v\n", contract)

	nn, ok := sdk.NewIntFromString(contract.Blocktoprove)
	if !ok {
		meter.ConsumeGas(refunded, "undoing refund")
		return &types.MsgPostproofResponse{Success: false, ErrorMessage: "cannot parse block to prove"}, nil
	}
	num := nn.Int64()

	h := sha256.New()
	_, err := io.WriteString(h, fmt.Sprintf("%d%s", num, msg.Item))
	if err != nil {
		meter.ConsumeGas(refunded, "undoing refund")
		return &types.MsgPostproofResponse{Success: false, ErrorMessage: err.Error()}, nil
	}
	hashName := h.Sum(nil)

	ctx.Logger().Debug("%v\n", hashes)

	var proof merkletree.Proof

	err = json.Unmarshal([]byte(msg.Hashlist), &proof)
	if err != nil {
		ctx.Logger().Debug("%v\n", err)
		meter.ConsumeGas(refunded, "undoing refund")
		return &types.MsgPostproofResponse{Success: false, ErrorMessage: err.Error()}, nil
	}

	ctx.Logger().Debug("proof: %v\n", proof)

	m, err := hex.DecodeString(contract.Merkle)
	if err != nil {
		ctx.Logger().Error("%v\n", err)
		meter.ConsumeGas(refunded, "undoing refund")
		return &types.MsgPostproofResponse{Success: false, ErrorMessage: err.Error()}, nil
	}
	verified, err := merkletree.VerifyProofUsing(hashName, false, &proof, [][]byte{m}, sha3.New512())
	if err != nil {
		ctx.Logger().Error("%v\n", err)
		meter.ConsumeGas(refunded, "undoing refund")
		return &types.MsgPostproofResponse{Success: false, ErrorMessage: err.Error()}, nil
	}

	if !verified {
		ctx.Logger().Debug("%s\n", "Cannot verify")
		meter.ConsumeGas(refunded, "undoing refund")
		return &types.MsgPostproofResponse{Success: false, ErrorMessage: "cannot verify proof"}, nil
	}

	if contract.Proofverified == "true" {
		return &types.MsgPostproofResponse{Success: false, ErrorMessage: "proof already verified"}, nil
	}

	contract.Proofverified = "true"
	k.SetActiveDeals(ctx, contract)

	refunded += meter.GasConsumed() - usedGas
	meter.RefundGas(meter.GasConsumed()-usedGas, "successful proof refund")

	return &types.MsgPostproofResponse{Success: true, ErrorMessage: ""}, nil
}
