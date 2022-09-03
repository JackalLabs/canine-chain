package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/rns module sentinel errors
var (
	ErrNoTLD    = sdkerrors.Register(ModuleName, 1100, "could not extract the tld from the name provided")
	ErrReserved = sdkerrors.Register(ModuleName, 1101, "tld is reserved by the system")
)
