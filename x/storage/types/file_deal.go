package types

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerror "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/libs/rand"
	"github.com/wealdtech/go-merkletree"
	"github.com/wealdtech/go-merkletree/sha3"
)

// VerifyProof checks whether a proof is valid against a file
func (f *UniversalFile) VerifyProof(proofData []byte, chunk int64, item string) bool {
	h := sha256.New()
	_, err := io.WriteString(h, fmt.Sprintf("%d%s", chunk, item))
	if err != nil {
		return false
	}
	hashName := h.Sum(nil)

	var proof merkletree.Proof // unmarshal proof
	err = json.Unmarshal(proofData, &proof)
	if err != nil {
		return false
	}

	verified, err := merkletree.VerifyProofUsing(hashName, false, &proof, [][]byte{f.Merkle}, sha3.New512())
	if err != nil {
		return false
	}

	return verified
}

// ResetChunk picks a new chunk to prove for a file
func (f *UniversalFile) ResetChunk(ctx sdk.Context, prover string, chunkSize int64) error {
	proof := f.Proofs[prover]
	if proof == nil {
		return sdkerror.ErrKeyNotFound
	}

	pieces := f.FileSize / chunkSize
	d := f.FileSize % chunkSize
	if d == 0 { // handle edge case where there is exactly full chunks with no extra bits
		pieces--
	}
	var newChunk int64
	if pieces > 0 { // if there is more than one piece we pick a random to prove

		r := rand.NewRand()
		r.Seed(ctx.BlockHeight() + int64(ctx.BlockGasMeter().GasConsumed()))
		newChunk = r.Int63n(pieces)
	}

	proof.ChunkToProve = newChunk

	return nil
}

// SetProven sets the proofs proven status to true and updates the proof window & picks a new chunk to verify
func (f *UniversalFile) SetProven(ctx sdk.Context, prover string, chunkSize int64) error {
	proof := f.Proofs[prover]
	if proof == nil {
		return sdkerror.ErrKeyNotFound
	}

	proof.LastProven = ctx.BlockHeight()
	err := f.ResetChunk(ctx, prover, chunkSize)
	if err != nil {
		return err
	}

	return nil
}

// Prove checks the validity of a proof and updates the proof window & picks a new chunk to verify
func (f *UniversalFile) Prove(ctx sdk.Context, prover string, proofData []byte, chunk int64, item string, chunkSize int64) error {
	valid := f.VerifyProof(proofData, chunk, item)

	if !valid {
		return ErrCannotVerifyProof
	}

	err := f.SetProven(ctx, prover, chunkSize)
	if err != nil {
		return err
	}

	return nil
}
