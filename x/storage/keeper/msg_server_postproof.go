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

	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

func VerifyDeal(merkle string, hashList string, num int64, item string) bool {
	h := sha256.New()
	_, err := io.WriteString(h, fmt.Sprintf("%d%s", num, item))
	if err != nil {
		return false
	}
	hashName := h.Sum(nil)

	var proof merkletree.Proof // unmarshal proof
	err = json.Unmarshal([]byte(hashList), &proof)
	if err != nil {
		return false
	}

	m, err := hex.DecodeString(merkle)
	if err != nil {
		return false
	}
	verified, err := merkletree.VerifyProofUsing(hashName, false, &proof, [][]byte{m}, sha3.New512())
	if err != nil {
		return false
	}

	return verified
}

func (k msgServer) Postproof(goCtx context.Context, msg *types.MsgPostproof) (*types.MsgPostproofResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	meter := ctx.GasMeter()
	usedGas := meter.GasConsumed()

	hashes := strings.Split(msg.Hashlist, ",")

	contract, found := k.GetActiveDeals(ctx, msg.Cid)
	if !found {
		ctx.Logger().Debug("%s, %s\n", "Contract not found", msg.Cid)
		return &types.MsgPostproofResponse{Success: false, ErrorMessage: fmt.Sprintf("contract not found: %s", msg.Cid)}, nil
	}

	ctx.Logger().Debug("Contract that was found: \n%v\n", contract)

	nn, ok := sdk.NewIntFromString(contract.Blocktoprove)
	if !ok {
		return &types.MsgPostproofResponse{Success: false, ErrorMessage: "cannot parse block to prove"}, nil
	}
	num := nn.Int64()

	ctx.Logger().Debug("%v\n", hashes)

	verified := VerifyDeal(contract.Merkle, msg.Hashlist, num, msg.Item)

	if !verified {
		ctx.Logger().Debug("%s\n", "Cannot verify")
		return &types.MsgPostproofResponse{Success: false, ErrorMessage: "cannot verify proof"}, nil
	}

	if contract.Proofverified == "true" {
		meter.RefundGas(meter.GasConsumed()-usedGas, "successful proof refund")
		return &types.MsgPostproofResponse{Success: false, ErrorMessage: "proof already verified"}, nil
	}

	contract.Proofverified = "true"
	k.SetActiveDeals(ctx, contract)

	meter.RefundGas(meter.GasConsumed()-usedGas, "successful proof refund")

	return &types.MsgPostproofResponse{Success: true, ErrorMessage: ""}, nil
}
