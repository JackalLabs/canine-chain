package bindings

// JackalQuery contains Jackal custom queries.
// TO DO
// Create Rust bindings to be located at:
// https://github.com/jackalLabs/Jackal-bindings/blob/main/packages/bindings/src/query.rs
type JackalQuery struct {
	/// returns a Files struct
	Files *Files `json:"files,omitempty"`
}

type Files struct {
	Address        string `json:"address"`
	Contents       string `json:"contents"`
	Owner          string `json:"owner"`
	ViewingAccess  string `json:"viewing_access"`
	EditAccess     string `json:"edit_access"`
	TrackingNumber string `json:"tracking_number"`
}

// Consider changing
type FilesResponse struct {
	Files string `json:"files"`
}
