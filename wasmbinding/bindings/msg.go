package bindings

type JackalMsg struct {
	/// Contracts can have a PubKey
	PostKey *PostKey `json:"post_key,omitempty"`
	/// Contracts can make Files
	PostFile *PostFile `json:"post_file,omitempty"`
	/// Contracts can buy storage
	BuyStorage *BuyStorage `json:"buy_storage,omitempty"`
}

// NOTE: Creator field is automatically the contract address
type PostKey struct {
	Key string `json:"key"`
}

// NOTE: Creator field is automatically the contract address
type PostFile struct {
	Merkle        string `json:"merkle"`
	FileSize      int64  `json:"file_size"`
	ProofInterval int64  `json:"proof_interval"`
	ProofType     int64  `json:"proof_type"`
	MaxProofs     int64  `json:"max_proofs"`
	Expires       int64  `json:"expires"`
	Note          string `json:"note"`
}

// NOTE: Creator field is automatically the contract address
type BuyStorage struct {
	ForAddress   string `json:"for_address"`
	DurationDays int64  `json:"duration_days"`
	Bytes        int64  `json:"bytes"`
	PaymentDenom string `json:"payment_denom"`
}
