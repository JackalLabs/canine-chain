package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/storage module sentinel errors
var (
	ErrDivideByZero        = sdkerrors.Register(ModuleName, 1110, "DivideByZero")
	ErrProviderNotFound    = sdkerrors.Register(ModuleName, 1111, "Provider not found. Please init your provider.")
	ErrNotValidTotalSpace  = sdkerrors.Register(ModuleName, 1112, "Not a valid total space. Please enter total number of bytes to provide.")
	ErrCannotVerifyProof   = sdkerrors.Register(ModuleName, 1113, "Cannot verify Proof")
	ErrDealNotFound        = sdkerrors.Register(ModuleName, 1114, "Cannot find active deal")
	ErrFormNotFound        = sdkerrors.Register(ModuleName, 1115, "Cannot find attestation form")
	ErrAttestInvalid       = sdkerrors.Register(ModuleName, 1116, "Cannot attest to form")
	ErrAttestAlreadyExists = sdkerrors.Register(ModuleName, 1117, "Attest already exists")
	ErrCannotVerifyProof  = sdkerrors.Register(ModuleName, 1118, "Cannot verify Proof")
	ErrNoCid              = sdkerrors.Register(ModuleName, 1119, "cid does not exist")
)
