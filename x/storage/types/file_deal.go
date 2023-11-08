package types

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerror "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/libs/rand"
	"github.com/wealdtech/go-merkletree/v2"
	"github.com/wealdtech/go-merkletree/v2/sha3"
)

// VerifyProof checks whether a proof is valid against a file
func (f *UnifiedFile) VerifyProof(proofData []byte, chunk int64, item []byte) bool {
	h := sha256.New()
	_, err := io.WriteString(h, fmt.Sprintf("%d%x", chunk, item))
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

func (f *UnifiedFile) AddProver(ctx sdk.Context, k ProofLoader, prover string) *FileProof {
	if len(f.Proofs) >= int(f.MaxProofs) {
		return nil
	}

	pk := f.MakeProofKey(prover)

	f.Proofs = append(f.Proofs, pk)

	p := FileProof{
		Prover:       prover,
		Merkle:       f.Merkle,
		Owner:        f.Owner,
		Start:        f.Start,
		LastProven:   ctx.BlockHeight(),
		ChunkToProve: 0,
	}

	k.SetProof(ctx, p)
	k.SetFile(ctx, *f)

	return &p
}

func (f *UnifiedFile) Save(ctx sdk.Context, k ProofLoader) {
	k.SetFile(ctx, *f)
}

func (f *UnifiedFile) RemoveProver(ctx sdk.Context, k ProofLoader, prover string) {
	pk := f.MakeProofKey(prover)
	f.RemoveProverWithKey(ctx, k, pk)
}

func (f *UnifiedFile) RemoveProverWithKey(ctx sdk.Context, k ProofLoader, proofKey string) {
	if len(f.Proofs) == 0 {
		return
	}

	for i, proof := range f.Proofs {
		ctx.Logger().Info(fmt.Sprintf("should we remove proof: %s == %s ?", proof, proofKey))
		if proof == proofKey {
			ctx.Logger().Info(fmt.Sprintf("removing proofs: %s == %s ?", proof, proofKey))

			front := f.Proofs[:i]
			back := f.Proofs[i+1:]

			// nolint:all
			f.Proofs = append(front, back...)

			k.RemoveProofWithBuiltKey(ctx, []byte(proofKey))
			f.Save(ctx, k)
		}
	}
}

func (f *UnifiedFile) GetProver(ctx sdk.Context, k ProofLoader, prover string) (*FileProof, error) {
	var proof *FileProof
	for _, proofKey := range f.Proofs {
		if proofKey == f.MakeProofKey(prover) {
			p, found := k.GetProofWithBuiltKey(ctx, []byte(proofKey))
			if found {
				proof = &p
				break
			}

		}
	}
	if proof == nil {
		return nil, sdkerror.Wrapf(sdkerror.ErrKeyNotFound, "cannot find proof with prover %s on file %x/%s/%d | %s", prover, f.Merkle, f.Owner, f.Start, strings.Join(f.Proofs, ", "))
	}

	return proof, nil
}

// ResetChunk picks a new chunk to prove for a file
func (f *UnifiedFile) ResetChunk(ctx sdk.Context, k ProofLoader, prover string, chunkSize int64) error {
	proof, err := f.GetProver(ctx, k, prover)
	if err != nil {
		return err
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

	k.SetProof(ctx, *proof)

	return nil
}

// SetProven sets the proofs proven status to true and updates the proof window & picks a new chunk to verify
func (f *UnifiedFile) SetProven(ctx sdk.Context, k ProofLoader, prover string, chunkSize int64) error {
	proof, err := f.GetProver(ctx, k, prover)
	if err != nil {
		return err
	}

	proof.LastProven = ctx.BlockHeight()
	err = f.ResetChunk(ctx, k, prover, chunkSize)
	if err != nil {
		return err
	}

	return nil
}

// Prove checks the validity of a proof and updates the proof window & picks a new chunk to verify
func (f *UnifiedFile) Prove(ctx sdk.Context, k ProofLoader, prover string, proofData []byte, chunk int64, item []byte, chunkSize int64) error {
	valid := f.VerifyProof(proofData, chunk, item)

	if !valid {
		return ErrCannotVerifyProof
	}

	err := f.SetProven(ctx, k, prover, chunkSize)
	if err != nil {
		return err
	}

	return nil
}
