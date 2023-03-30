package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/amm module sentinel errors
var (
	ErrNegativeLockDuration    = sdkerrors.Register(ModuleName, 1, "lock duration cannot be negative")
	ErrInvalidValue            = sdkerrors.Register(ModuleName, 2, "invalid value")
	ErrLiquidityPoolExists     = sdkerrors.Register(ModuleName, 3, "liquidity pool exists")
	ErrLiquidityPoolNotFound   = sdkerrors.Register(ModuleName, 4, "liquidity pool not found")
	ErrLProviderRecordNotFound = sdkerrors.Register(ModuleName, 5, "liquidity provider record not found")
	ErrLProviderRecordExists   = sdkerrors.Register(ModuleName, 6, "liquidity provider exists")
	ErrNegativeCoin = sdkerrors.Register(ModuleName, 7, "coin amount is negative")
	ErrInvalidPoolName = sdkerrors.Register(ModuleName, 8, "invalid pool name")
)
