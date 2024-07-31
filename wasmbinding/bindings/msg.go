package bindings

type JackalMsg struct {
	/// Contracts can create files
	/// will they be namespaced under the contract's address?
	/// A contract may create any number of independent files.
	PostKey *PostKey `json:"post_key,omitempty"`
}

// WARNING DANGER
// NOTE: Sender can currently be spoofed
// We have notes everywhere to flesh out different authentication methods
type PostKey struct {
	Sender string `json:"sender"`
	Key    string `json:"key"`
}
