package types

import (
	"fmt"

	"github.com/jackalLabs/canine-chain/v5/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
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

	CollateralCollectorName = "storage_collateral_name"
)

func gaugeName(gauge PaymentGauge) string {
	return fmt.Sprintf("gauge:%x", gauge.Id)
}

func GetGaugeAccount(gauge PaymentGauge) (sdk.AccAddress, error) {
	return types.GetAccount(gaugeName(gauge))
}

func KeyPrefix(p string) []byte {
	return []byte(p)
}
