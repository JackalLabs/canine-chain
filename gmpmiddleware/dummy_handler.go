package gmp_middleware

import sdk "github.com/cosmos/cosmos-sdk/types"

type BankK interface {
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
}

type DummyHandler struct {
	bank BankK
}

func NewDummyHandler(k BankK) *DummyHandler {
	return &DummyHandler{
		bank: k,
	}
}

func (h DummyHandler) HandleGeneralMessage(ctx sdk.Context, srcChain, srcAddress string, destAddress string, payload []byte) error {
	ctx.Logger().Info("HandleGeneralMessage called",
		"srcChain", srcChain,
		"srcAddress", srcAddress,
		"destAddress", destAddress,
		"payload", payload,
		"module", "x/gmp-middleware",
	)
	return nil
}

func (h DummyHandler) HandleGeneralMessageWithToken(ctx sdk.Context, srcChain, srcAddress string, destAddress string, payload []byte, coin sdk.Coin) error {
	ctx.Logger().Info("HandleGeneralMessageWithToken called",
		"srcChain", srcChain,
		"srcAddress", srcAddress,
		"destAddress", destAddress,
		"payload", payload,
		"coin", coin,
	)

	return nil
}
