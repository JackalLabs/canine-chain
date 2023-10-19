package types

func (f *UnifiedFile) MakeProofKey(prover string) string {
	return string(ProofKey(prover, f.Merkle, f.Owner, f.Start))
}
