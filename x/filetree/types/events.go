package types

// filetree module event types
const (
	EventTypeUpdateFile = "file_updated"
	EventTypeRemoveFile = "file_removed"
	EventTypePostKey    = "key_posted"

	AttributeKeyOwner       = "file_owner" // update file
	AttributeKeyFileAddress = "file_address"

	EventTypeAddViewers    = "viewers_added"
	EventTypeAddEditors    = "editors_added"
	EventTypeRemoveEditors = "editors_removed"
	EventTypeRemoveViewers = "viewers_removed"
	EventTypeResetViewers  = "viewers_reset"
	EventTypeResetEditors  = "editors_reset"
	EventTypeChangeOwner   = "owner_changed"
	EventTypeMakeRoot      = "root_made"
	EventTypePostFile      = "file_posted"

	AttributeValueCategory = ModuleName

	AttributeKeySigner   = "signer"
	AttributeKeyNewOwner = "new_owner"
	AttributeKeyKey      = "key"
)
