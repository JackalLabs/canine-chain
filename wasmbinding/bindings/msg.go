package bindings

import storagetypes "github.com/jackalLabs/canine-chain/v4/x/storage/types"

type JackalMsg struct {
	// STORAGE MODULE
	PostFile          *storagetypes.MsgPostFile          `json:"post_file,omitempty"`
	DeleteFile        *storagetypes.MsgDeleteFile        `json:"delete_file,omitempty"`
	BuyStorage        *storagetypes.MsgBuyStorage        `json:"buy_storage,omitempty"`
	RequestReportForm *storagetypes.MsgRequestReportForm `json:"request_report_form,omitempty"`
}
