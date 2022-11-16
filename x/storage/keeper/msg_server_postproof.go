package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	// sdk "github.com/cosmos/cosmos-sdk/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	merkletree "github.com/wealdtech/go-merkletree"

	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func printProof(ctx sdk.Context, proof merkletree.Proof) {
	ctx.Logger().Info("Hashes: %v\nIndex: %d\n", proof.Hashes, proof.Index)
}

func (k msgServer) Postproof(goCtx context.Context, msg *types.MsgPostproof) (*types.MsgPostproofResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	hashes := strings.Split(msg.Hashlist, ",")

	contract, found := k.GetActiveDeals(ctx, msg.Cid)
	if !found {
		ctx.Logger().Debug("%s, %s\n", "Contract not found", msg.Cid)

		return nil, fmt.Errorf("contract not found")
	}

	if contract.Proofverified == "true" {
		return nil, fmt.Errorf("proof already verified")
	}

	ctx.Logger().Debug("Contract that was found: \n%v\n", contract)

	nn, ok := sdk.NewIntFromString(contract.Blocktoprove)
	if !ok {
		return nil, fmt.Errorf("failed to parse block")
	}
	num := nn.Int64()

	h := sha256.New()
	_, err := io.WriteString(h, fmt.Sprintf("%d%s", num, msg.Item))
	if err != nil {
		return nil, err
	}
	hashName := h.Sum(nil)

	ctx.Logger().Debug("%v\n", hashes)

	var proof merkletree.Proof

	err = json.Unmarshal([]byte(msg.Hashlist), &proof)
	if err != nil {
		ctx.Logger().Debug("%v\n", err)
		return nil, err
	}

	ctx.Logger().Debug("proof: %v\n", proof)
	printProof(ctx, proof)

	m, err := hex.DecodeString(contract.Merkle)
	if err != nil {
		ctx.Logger().Error("%v\n", err)
		return nil, fmt.Errorf("could not build merkle tree")
	}
	verified, err := merkletree.VerifyProof(hashName, &proof, m)
	if err != nil {
		ctx.Logger().Error("%v\n", err)

		return nil, fmt.Errorf("could not build merkle tree")
	}

	if !verified {
		ctx.Logger().Debug("%s\n", "Cannot verify")

		return nil, fmt.Errorf("file chunk was not verified")
	}

	if contract.Proofverified == "false" {
		ctx.GasMeter().RefundGas(ctx.GasMeter().GasConsumed(), "successful proof refund.")
	}

	contract.Proofverified = "true"
	k.SetActiveDeals(ctx, contract)

	return &types.MsgPostproofResponse{Merkle: fmt.Sprintf("%v", proof)}, nil
}
