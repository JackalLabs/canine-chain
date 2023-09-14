package types

import (
	errorsmod "cosmossdk.io/errors"
)

var (
	ErrBadMetadataFormatMsg = "wasm metadata not properly formatted for: '%v'. %s"
	ErrBadExecutionMsg      = "cannot execute contract: %v"

	ErrMsgValidation = errorsmod.Register("gmp-middleware", 2, "error in gmp-middleware message validation")
	ErrMarshaling    = errorsmod.Register("gmp-middleware", 3, "cannot marshal the ICS20 packet")
	ErrInvalidPacket = errorsmod.Register("gmp-middleware", 4, "invalid packet data")
	ErrBadResponse   = errorsmod.Register("gmp-middleware", 5, "cannot create response")
	ErrWasmError     = errorsmod.Register("gmp-middleware", 6, "wasm error")
	ErrBadSender     = errorsmod.Register("gmp-middleware", 7, "bad sender")
)
