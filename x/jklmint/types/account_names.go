package types

import (
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DevelopmentAndGrants = "dev_and_grants"
)

func GetDevAndGrantAccount() (sdk.AccAddress, error) {
	adr := hex.EncodeToString([]byte(DevelopmentAndGrants))
	return sdk.AccAddressFromHex(adr)
}
