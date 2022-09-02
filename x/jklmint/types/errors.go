package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/jklmint module sentinel errors
var (
	ErrCannotParseFloat = sdkerrors.Register(ModuleName, 1101, "cannot parse float")
	ErrZeroDivision     = sdkerrors.Register(ModuleName, 1102, "cannot use zero value for division")
)
