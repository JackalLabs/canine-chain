package types

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
