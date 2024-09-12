package bindings

import storagetypes "github.com/jackalLabs/canine-chain/v4/x/storage/types"

type JackalMsg struct {
	/// Contracts can make Files
	PostFile *storagetypes.MsgPostFile `json:"post_file,omitempty"`
}
