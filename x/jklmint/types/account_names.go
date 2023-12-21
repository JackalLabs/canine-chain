package types

import (
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DevelopmentAndGrants   = "dev_and_grants"
	ProtocolOwnedLiquidity = "protocol_owned_liq"
)

func GetDevAndGrantAccount() (sdk.AccAddress, error) {
	adr := hex.EncodeToString([]byte(DevelopmentAndGrants))
	return sdk.AccAddressFromHex(adr)
}

func GetPOLAccount() (sdk.AccAddress, error) {
	adr := hex.EncodeToString([]byte(ProtocolOwnedLiquidity))
	return sdk.AccAddressFromHex(adr)
}
