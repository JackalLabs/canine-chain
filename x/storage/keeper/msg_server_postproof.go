package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	merkletree "github.com/wealdtech/go-merkletree/v2"
	"github.com/wealdtech/go-merkletree/v2/sha3"

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

func VerifyDealRaw(merkle string, hashList []byte, num int64, item []byte) bool {
	h := sha256.New()
	_, err := io.WriteString(h, fmt.Sprintf("%d%x", num, item))
	if err != nil {
		return false
	}
	hashName := h.Sum(nil)

	var proof merkletree.Proof // unmarshal proof
	err = json.Unmarshal(hashList, &proof)
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

	p := k.GetParams(ctx)
	alreadyVerified := contract.IsVerified(ctx.BlockHeight(), p.ProofWindow)

	if alreadyVerified {
		meter.RefundGas(meter.GasConsumed()-usedGas, "successful proof refund")
		return &types.MsgPostproofResponse{Success: false, ErrorMessage: "proof already verified"}, nil
	}

	contract.LastProof = ctx.BlockHeight()
	k.SetActiveDeals(ctx, contract)

	meter.RefundGas(meter.GasConsumed()-usedGas, "successful proof refund")

	return &types.MsgPostproofResponse{Success: true, ErrorMessage: ""}, nil
}

func (k msgServer) PostProof(goCtx context.Context, msg *types.MsgPostProof) (*types.MsgPostproofResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	meter := ctx.GasMeter()
	usedGas := meter.GasConsumed()

	var contract types.ActiveDeals

	k.IterateActiveDeals(ctx, func(deal types.ActiveDeals) bool {
		if deal.Merkle != hex.EncodeToString(msg.Merkle) {
			return false
		}

		if deal.Signee != msg.Owner {
			return false
		}

		start, err := strconv.ParseInt(deal.Startblock, 10, 64)
		if err != nil {
			return false
		}

		if start != msg.Start {
			return false
		}

		contract = deal

		return true
	})

	if contract.Merkle == "" {
		ctx.Logger().Debug("Contract not found: %x\n", msg.Merkle)
		return &types.MsgPostproofResponse{Success: false, ErrorMessage: fmt.Sprintf("contract not found: %x", msg.Merkle)}, nil
	}

	ctx.Logger().Debug("Contract that was found: \n%v\n", contract)

	nn, ok := sdk.NewIntFromString(contract.Blocktoprove)
	if !ok {
		return &types.MsgPostproofResponse{Success: false, ErrorMessage: "cannot parse block to prove"}, nil
	}
	num := nn.Int64()

	ctx.Logger().Debug("%v\n", msg.HashList)

	verified := VerifyDealRaw(contract.Merkle, msg.HashList, num, msg.Item)

	if !verified {
		ctx.Logger().Debug("%s\n", "Cannot verify")
		return &types.MsgPostproofResponse{Success: false, ErrorMessage: "cannot verify proof"}, nil
	}

	p := k.GetParams(ctx)
	alreadyVerified := contract.IsVerified(ctx.BlockHeight(), p.ProofWindow)

	if alreadyVerified {
		meter.RefundGas(meter.GasConsumed()-usedGas, "successful proof refund")
		return &types.MsgPostproofResponse{Success: false, ErrorMessage: "proof already verified"}, nil
	}

	contract.LastProof = ctx.BlockHeight()
	k.SetActiveDeals(ctx, contract)

	meter.RefundGas(meter.GasConsumed()-usedGas, "successful proof refund")

	return &types.MsgPostproofResponse{Success: true, ErrorMessage: ""}, nil
}
