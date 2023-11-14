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

func (f *UnifiedFile) Proven(ctx sdk.Context, k ProofLoader, height int64, prover string) bool {
	interval := f.ProofInterval

	lastWindowEnd := height - (height % interval)
	lastWindowStart := lastWindowEnd - interval

	proof, err := f.GetProver(ctx, k, prover)
	if err != nil {
		return false
	}

	lastProven := proof.LastProven

	return lastProven > lastWindowStart
}

func FileNeedsProven(start int64, interval int64, height int64) bool {
	lastWindowEnd := start - (start % interval)

	return height > lastWindowEnd
}
