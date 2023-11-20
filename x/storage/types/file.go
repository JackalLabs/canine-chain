package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func (f *UnifiedFile) MakeProofKey(prover string) string {
	return string(ProofKey(prover, f.Merkle, f.Owner, f.Start))
}

func (f *UnifiedFile) ContainsProver(prover string) bool {
	for _, proof := range f.Proofs {
		if proof == string(ProofKey(prover, f.Merkle, f.Owner, f.Start)) {
			return true
		}
	}
	return false
}

func (f *UnifiedFile) ProvenProviderLastBlock(ctx sdk.Context, k ProofLoader, height int64, prover string) bool {
	proof, err := f.GetProver(ctx, k, prover)
	if err != nil {
		return false
	}

	return f.ProvenLastBlock(height, proof.LastProven)
}

func getRoundedWindow(currentHeight int64, start int64, window int64) int64 {
	k := currentHeight - start
	we := k - (k % window) + start

	return we
}

func (f *UnifiedFile) ProvenLastBlock(height int64, lastProven int64) bool {
	window := getRoundedWindow(height, f.Start, f.ProofInterval)

	lastWindowStart := window - f.ProofInterval

	return lastProven >= lastWindowStart // if last proven has been since the window start we can ski it
}

func (f *UnifiedFile) ProvenThisBlock(height int64, lastProven int64) bool {
	window := getRoundedWindow(height, f.Start, f.ProofInterval)

	return lastProven >= window // if last proven has been since the window start we can ski it
}

func (f *UnifiedFile) IsYoung(height int64) bool {
	return f.Start+f.ProofInterval >= height
}
