package bindings

import (
	filetreetypes "github.com/jackalLabs/canine-chain/v4/x/filetree/types"
	storagetypes "github.com/jackalLabs/canine-chain/v4/x/storage/types"
)

type JackalMsg struct {
	// STORAGE MODULE
	PostFile          *storagetypes.MsgPostFile          `json:"post_file,omitempty"`
	DeleteFile        *storagetypes.MsgDeleteFile        `json:"delete_file,omitempty"`
	BuyStorage        *storagetypes.MsgBuyStorage        `json:"buy_storage,omitempty"`
	RequestReportForm *storagetypes.MsgRequestReportForm `json:"request_report_form,omitempty"`

	// FILETREE MODULE
	PostFileTree *filetreetypes.MsgPostFile   `json:"post_file_tree,omitempty"`
	AddViewers   *filetreetypes.MsgAddViewers `json:"add_viewers,omitempty"`
	PostKey      *filetreetypes.MsgPostKey    `json:"post_key,omitempty"`
}

/*

  rpc PostFile(MsgPostFile) returns (MsgPostFileResponse);
  rpc AddViewers(MsgAddViewers) returns (MsgAddViewersResponse);
  rpc PostKey(MsgPostKey) returns (MsgPostKeyResponse);
  rpc DeleteFile(MsgDeleteFile) returns (MsgDeleteFileResponse);
  rpc RemoveViewers(MsgRemoveViewers) returns (MsgRemoveViewersResponse);
  rpc ProvisionFileTree(MsgProvisionFileTree) returns (MsgProvisionFileTreeResponse);

  rpc AddEditors(MsgAddEditors) returns (MsgAddEditorsResponse);
  rpc RemoveEditors(MsgRemoveEditors) returns (MsgRemoveEditorsResponse);
  rpc ResetEditors(MsgResetEditors) returns (MsgResetEditorsResponse);
  rpc ResetViewers(MsgResetViewers) returns (MsgResetViewersResponse);
  rpc ChangeOwner(MsgChangeOwner) returns (MsgChangeOwnerResponse);

*/
