package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/storage module sentinel errors
var (
	ErrDivideByZero       = sdkerrors.Register(ModuleName, 1110, "DivideByZero")
	ErrProviderNotFound   = sdkerrors.Register(ModuleName, 1111, "Provider not found. Please init your provider.")
	ErrNotValidTotalSpace = sdkerrors.Register(ModuleName, 1112, "Not a valid total space. Please enter total number of bytes to provide.")
	ErrCannotVerifyProof  = sdkerrors.Register(ModuleName, 1113, "Cannot verify Proof")
	ErrNoCid              = sdkerrors.Register(ModuleName, 1114, "cid does not exist")
)
