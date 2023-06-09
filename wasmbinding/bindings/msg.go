package bindings

type JackalMsg struct {
	/// Contracts can create files
	/// will they be namespaced under the contract's address?
	/// A contract may create any number of independent files.
	MakeRoot   *MakeRoot   `json:"make_root,omitempty"`
	PostFiles  *PostFiles  `json:"post_files,omitempty"`
	DeleteFile *DeleteFile `json:"delete_file,omitempty"`
	BuyStorage *BuyStorage `json:"buy_storage,omitempty"`
}

type PostFiles struct {
	Account        string `json:"account"`
	HashParent     string `json:"hashparent"`
	HashChild      string `json:"hashchild"`
	Contents       string `json:"contents"`
	Viewers        string `json:"viewers"`
	Editors        string `json:"editors"`
	TrackingNumber string `json:"trackingnumber"`
}

type MakeRoot struct {
	Editors        string `json:"editors"`
	Viewers        string `json:"viewers"`
	TrackingNumber string `json:"trackingnumber"`
}

// / creator == broadcaster of the msg
type DeleteFile struct {
	HashPath string `json:"hashpath"` // the full merklePath
	Account  string `json:"account"`
}

type BuyStorage struct {
	ForAddress   string `json:"foraddress"` // the full merklePath
	Duration     string `json:"duration"`
	Bytes        string `json:"bytes"` // the full merklePath
	PaymentDenom string `json:"paymentdenom"`
}
