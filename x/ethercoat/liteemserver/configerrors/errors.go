package configerrors

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const EthLikeModule string = "eth-like"

var (
	ErrBlankAPINamespace       = sdkerrors.Register(EthLikeModule, 2300, "")
	ErrNegativeHTTPTimeout     = sdkerrors.Register(EthLikeModule, 2301, "")
	ErrNegativeHTTPIdleTimeout = sdkerrors.Register(EthLikeModule, 2302, "")
)
