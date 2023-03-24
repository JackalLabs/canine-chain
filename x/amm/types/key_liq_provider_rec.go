package types

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ binary.ByteOrder

const (
	// LiqProviderRecKeyPrefix is the prefix to retrieve all LiqProviderRec
	LiqProviderRecKeyPrefix = "LiqProviderRec/value/"
	RefKeyPrefix             = "LiqProviderRecRef/"
	// A separator inserted between keys.
)

// LiqProviderRecKey returns the store key to retrieve a LiqProviderRec
func LiqProviderRecKey(
	poolName string,
	provider string,
) []byte {

	poolBytes := []byte(poolName)
	addrBytes := []byte(provider)

	return CombineKeys(poolBytes, addrBytes)
}

// LiqProviderRecKey returns the store key to retrieve a LiqProviderRec
// reference.
func LiqProviderRecRefKey(
	poolName string,
	provider sdk.AccAddress,
) []byte {
	poolBytes := []byte(poolName)
	addrBytes := []byte(provider.String())

	return CombineKeys(addrBytes, poolBytes)
}

// Takes LiqProviderRec struct to generate store key.
// Key format is: {poolName}{provider}
func GetProviderKey(record LiqProviderRec) []byte {
	return LiqProviderRecKey(record.PoolName, record.Provider)
}

// Takes LiqProviderRec struct to generate reference key.
// Key format is: {provider}{provider}
func GetProviderRefKey(record LiqProviderRec) []byte {

	poolBytes := []byte(record.PoolName)
	addrBytes := []byte(record.Provider)

	return CombineKeys(addrBytes, poolBytes)
}
