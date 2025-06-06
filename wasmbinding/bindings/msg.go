package bindings

import (
	filetreetypes "github.com/jackalLabs/canine-chain/v4/x/filetree/types"
	notificationstypes "github.com/jackalLabs/canine-chain/v4/x/notifications/types"
	storagetypes "github.com/jackalLabs/canine-chain/v4/x/storage/types"
)

type JackalMsg struct {
	// STORAGE
	PostFile          *storagetypes.MsgPostFile          `json:"post_file,omitempty"`
	DeleteFile        *storagetypes.MsgDeleteFile        `json:"delete_file,omitempty"`
	BuyStorage        *storagetypes.MsgBuyStorage        `json:"buy_storage,omitempty"`
	RequestReportForm *storagetypes.MsgRequestReportForm `json:"request_report_form,omitempty"`

	// FILETREE
	PostFileTree      *filetreetypes.MsgPostFile          `json:"post_file_tree,omitempty"`
	AddViewers        *filetreetypes.MsgAddViewers        `json:"add_viewers,omitempty"`
	PostKey           *filetreetypes.MsgPostKey           `json:"post_key,omitempty"`
	DeleteFileTree    *filetreetypes.MsgDeleteFile        `json:"delete_file_tree,omitempty"`
	RemoveViewers     *filetreetypes.MsgRemoveViewers     `json:"remove_viewers,omitempty"`
	ProvisionFileTree *filetreetypes.MsgProvisionFileTree `json:"provision_file_tree,omitempty"`
	AddEditors        *filetreetypes.MsgAddEditors        `json:"add_editors,omitempty"`
	RemoveEditors     *filetreetypes.MsgRemoveEditors     `json:"remove_editors,omitempty"`
	ResetEditors      *filetreetypes.MsgResetEditors      `json:"reset_editors,omitempty"`
	ResetViewers      *filetreetypes.MsgResetViewers      `json:"reset_viewers,omitempty"`
	ChangeOwner       *filetreetypes.MsgChangeOwner       `json:"change_owner,omitempty"`

	// NOTIFICATIONS
	CreateNotification *notificationstypes.MsgCreateNotification `json:"create_notification,omitempty"`
	DeleteNotification *notificationstypes.MsgDeleteNotification `json:"delete_notification,omitempty"`
	BlockSenders       *notificationstypes.MsgBlockSenders       `json:"block_senders,omitempty"`
}
