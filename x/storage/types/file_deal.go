package types

import (
	"fmt"
	"strings"

	"github.com/jackalLabs/canine-chain/v3/x/storage/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerror "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/libs/rand"
)

// VerifyProof checks whether a proof is valid against a file
func (f *UnifiedFile) VerifyProof(proofData []byte, chunk int64, item []byte) bool {
	return utils.VerifyProof(f.Merkle, proofData, chunk, item)
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

		var gs int64
		gasMeter := ctx.BlockGasMeter()
		if gasMeter != nil {
			gs = int64(gasMeter.GasConsumed())
		}
		h := ctx.BlockHeight()

		r := rand.NewRand()
		r.Seed(gs + h)
		newChunk = r.Int63n(pieces)
	}

	proof.ChunkToProve = newChunk

	k.SetProof(ctx, *proof)

	return nil
}

// ResetChunkWithProof picks a new chunk to prove for a file
func (f *UnifiedFile) ResetChunkWithProof(ctx sdk.Context, proof *FileProof, chunkSize int64) error {
	pieces := f.FileSize / chunkSize
	d := f.FileSize % chunkSize
	if d == 0 { // handle edge case where there is exactly full chunks with no extra bits
		pieces--
	}
	var newChunk int64
	if pieces > 0 { // if there is more than one piece we pick a random to prove

		var gs int64
		gasMeter := ctx.BlockGasMeter()
		if gasMeter != nil {
			gs = int64(gasMeter.GasConsumed())
		}
		h := ctx.BlockHeight()

		r := rand.NewRand()
		r.Seed(gs + h)
		newChunk = r.Int63n(pieces)
	}

	proof.ChunkToProve = newChunk

	return nil
}

// SetProven sets the proofs proven status to true and updates the proof window & picks a new chunk to verify
func (f *UnifiedFile) SetProven(ctx sdk.Context, proof *FileProof, chunkSize int64) error {
	proof.LastProven = ctx.BlockHeight() // sets the newest proof window

	err := f.ResetChunkWithProof(ctx, proof, chunkSize)
	if err != nil {
		return err
	}

	return nil
}

// Prove checks the validity of a proof and updates the proof window & picks a new chunk to verify
func (f *UnifiedFile) Prove(ctx sdk.Context, proof *FileProof, proofData []byte, item []byte, chunkSize int64) error {
	valid := f.VerifyProof(proofData, proof.ChunkToProve, item)

	if !valid {
		return ErrCannotVerifyProof
	}

	err := f.SetProven(ctx, proof, chunkSize)
	if err != nil {
		return err
	}

	return nil
}
