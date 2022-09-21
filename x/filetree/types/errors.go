package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/filetree module sentinel errors
var (
	ErrNoAccess       = sdkerrors.Register(ModuleName, 1101, "wrong permissions for file")
	ErrFileNotFound   = sdkerrors.Register(ModuleName, 1102, "file not found")
	ErrCantMarshall   = sdkerrors.Register(ModuleName, 1103, "cannot marshall data into json")
	ErrCantUnmarshall = sdkerrors.Register(ModuleName, 1104, "cannot unmarshall data from json")
	ErrPubKeyNotFound = sdkerrors.Register(ModuleName, 1105, "user's public key not found. Account not inited or wrong address")
)
