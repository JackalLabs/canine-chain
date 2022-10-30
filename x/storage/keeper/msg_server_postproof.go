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

func printProof(proof merkletree.Proof) {
	fmt.Printf("Hashes: %v\nIndex: %d\n", proof.Hashes, proof.Index)
}

func (k msgServer) Postproof(goCtx context.Context, msg *types.MsgPostproof) (*types.MsgPostproofResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	hashes := strings.Split(msg.Hashlist, ",")

	contract, found := k.GetActiveDeals(ctx, msg.Cid)
	if !found {
		fmt.Printf("%s, %s\n", "Contract not found", msg.Cid)

		return nil, fmt.Errorf("contract not found")
	}

	nn, ok := sdk.NewIntFromString(contract.Blocktoprove)
	if !ok {
		return nil, fmt.Errorf("failed to parse block")
	}
	num := nn.Int64()

	h := sha256.New()
	io.WriteString(h, fmt.Sprintf("%d%s", num, msg.Item))
	hashName := h.Sum(nil)

	fmt.Printf("%v\n", hashes)

	var proof merkletree.Proof

	err := json.Unmarshal([]byte(msg.Hashlist), &proof)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}

	fmt.Printf("proof: %v\n", proof)
	printProof(proof)

	m, _ := hex.DecodeString(contract.Merkle)

	verified, err := merkletree.VerifyProof(hashName, &proof, m)
	if err != nil {
		fmt.Printf("%v\n", err)

		return nil, fmt.Errorf("could not build merkle tree")
	}

	if !verified {
		fmt.Printf("%s\n", "Cannot verify")

		return nil, fmt.Errorf("file chunk was not verified")
	}

	deal, found := k.GetActiveDeals(ctx, msg.Cid)
	if !found {
		return nil, fmt.Errorf("deal not found")
	}
	if deal.Proofverified == "false" {
		ctx.GasMeter().RefundGas(ctx.GasMeter().GasConsumed(), "successful proof refund.")
	}
	deal.Proofverified = "true"
	k.SetActiveDeals(ctx, deal)

	return &types.MsgPostproofResponse{Merkle: ""}, nil
}
