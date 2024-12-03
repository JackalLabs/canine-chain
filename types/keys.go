package types

import (
	"crypto/sha256"
	"encoding/hex"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	ProtocolOwnedLiquidityName = "protocol_owned_liq"
)

func GetPOLAccount() (sdk.AccAddress, error) {
	return GetAccount(ProtocolOwnedLiquidityName)
}

func GetAccount(name string) (sdk.AccAddress, error) {
	s := sha256.New()
	s.Write([]byte(name))
	m := s.Sum(nil)
	mh := hex.EncodeToString(m)
	adr, err := sdk.AccAddressFromHex(mh)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "cannot get account account")
	}
	return adr, nil
}
