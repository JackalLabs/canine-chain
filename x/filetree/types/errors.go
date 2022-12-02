package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/filetree module sentinel errors
var (
	ErrNoAccess           = sdkerrors.Register(ModuleName, 1101, "wrong permissions for file")
	ErrFileNotFound       = sdkerrors.Register(ModuleName, 1102, "file not found")
	ErrCantMarshall       = sdkerrors.Register(ModuleName, 1103, "cannot marshall data into json")
	ErrCantUnmarshall     = sdkerrors.Register(ModuleName, 1104, "cannot unmarshall data from json")
	ErrPubKeyNotFound     = sdkerrors.Register(ModuleName, 1105, "user's public key not found. Account not inited or wrong address")
	ErrParentFileNotFound = sdkerrors.Register(ModuleName, 1106, "Parent folder does not exist")
	ErrCannotWrite        = sdkerrors.Register(ModuleName, 1107, "You are not permitted to write to this folder")
	ErrNoViewingAccess    = sdkerrors.Register(ModuleName, 1108, "You do not have viewing access. Failed to decrypt.")
	ErrTrackerNotFound    = sdkerrors.Register(ModuleName, 1109, "Tracking number not found")
	ErrCannotDelete       = sdkerrors.Register(ModuleName, 1110, "You are not authorized to delete this file")
	ErrNotOwner           = sdkerrors.Register(ModuleName, 1111, "Not permitted to remove or reset edit/view access. You are not the owner of this file")
	ErrCantGiveAway       = sdkerrors.Register(ModuleName, 1112, "You do not own this file and cannot give it away")
	ErrAlreadyExists      = sdkerrors.Register(ModuleName, 1113, "Proposed new owner already has a file set with this path name. No duplicates allowed.")
	ErrCannotAllowEdit    = sdkerrors.Register(ModuleName, 1114, "Unathorized. Only the owner can add an editor.")
	ErrCannotAllowView    = sdkerrors.Register(ModuleName, 1115, "Unathorized. Only the owner can add a viewer.")
	ErrMissingAESKey      = sdkerrors.Register(ModuleName, 1116, "AES IV and key required")
)
