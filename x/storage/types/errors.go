package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/storage module error messages
var (
	ErrDivideByZero        = sdkerrors.Register(ModuleName, 1110, "divide by zero")
	ErrProviderNotFound    = sdkerrors.Register(ModuleName, 1111, "provider not found please init your provider")
	ErrNotValidTotalSpace  = sdkerrors.Register(ModuleName, 1112, "not a valid total space please enter total number of bytes to provide")
	ErrDealNotFound        = sdkerrors.Register(ModuleName, 1114, "cannot find active deal")
	ErrFormNotFound        = sdkerrors.Register(ModuleName, 1115, "cannot find attestation form")
	ErrAttestInvalid       = sdkerrors.Register(ModuleName, 1116, "cannot attest to form")
	ErrAttestAlreadyExists = sdkerrors.Register(ModuleName, 1117, "attest already exists")
	ErrCannotVerifyProof   = sdkerrors.Register(ModuleName, 1118, "cannot verify proof")
	ErrNoCid               = sdkerrors.Register(ModuleName, 1119, "cid does not exist")
	ErrProviderExists      = sdkerrors.Register(ModuleName, 1120, "provider already exists")
	ErrCancelContract      = sdkerrors.Register(ModuleName, 1121, "cannot cancel contract")
)
