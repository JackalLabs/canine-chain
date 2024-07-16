package types

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	// ModuleName defines the module name
	ModuleName = "storage"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_storage"

	AddressPrefix = "jkl"
	CidPrefix     = "jklc"

	CollateralCollectorName    = "storage_collateral_name"
	ProtocolOwnedLiquidityName = "protocol_owned_liq"
)

func gaugeName(gauge PaymentGauge) string {
	return fmt.Sprintf("gauge:%x", gauge.Id)
}

func GetGaugeAccount(gauge PaymentGauge) (sdk.AccAddress, error) {
	return GetAccount(gaugeName(gauge))
}

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

func KeyPrefix(p string) []byte {
	return []byte(p)
}
