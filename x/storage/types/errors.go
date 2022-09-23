package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/storage module sentinel errors
var (
	ErrDivideByZero = sdkerrors.Register(ModuleName, 1110, "DivideByZero")
)
