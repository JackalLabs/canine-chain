package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/lp module sentinel errors
var (
	ErrNegativeLockDuration    = sdkerrors.Register(ModuleName, 1, "Lock duration cannot be negative")
	ErrInvalidValue            = sdkerrors.Register(ModuleName, 2, "Invalid value")
	ErrLiquidityPoolExists     = sdkerrors.Register(ModuleName, 3, "Liquidity pool exists")
	ErrLiquidityPoolNotFound   = sdkerrors.Register(ModuleName, 4, "Liquidity pool not found")
	ErrLProviderRecordNotFound = sdkerrors.Register(ModuleName, 5, "Liquidity provider record not found")
	ErrLProviderRecordExists   = sdkerrors.Register(ModuleName, 6, "Liquidity provider exists")
)
