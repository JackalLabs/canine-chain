package bindings

type JackalMsg struct {
	/// Contracts can make Files
	PostFile *PostFile `json:"post_file,omitempty"`
}

// NOTE: Creator field is automatically the contract address
type PostFile struct {
	Merkle        []byte `json:"merkle"`
	FileSize      int64  `json:"file_size"`
	ProofInterval int64  `json:"proof_interval"`
	ProofType     int64  `json:"proof_type"`
	MaxProofs     int64  `json:"max_proofs"`
	Expires       int64  `json:"expires"`
	Note          string `json:"note"`
}
