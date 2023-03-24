package types

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ binary.ByteOrder

const (
	// ProviderRecordKeyPrefix is the prefix to retrieve all ProviderRecord
	ProviderRecordKeyPrefix = "ProviderRecord/value/"
	RefKeyPrefix             = "ProviderRecordRef/"
	// A separator inserted between keys.
)

// ProviderRecordKey returns the store key to retrieve a ProviderRecord
func ProviderRecordKey(
	poolName string,
	provider string,
) []byte {

	poolBytes := []byte(poolName)
	addrBytes := []byte(provider)

	return CombineKeys(poolBytes, addrBytes)
}

// ProviderRecordKey returns the store key to retrieve a ProviderRecord
// reference.
func ProviderRecordRefKey(
	poolName string,
	provider sdk.AccAddress,
) []byte {
	poolBytes := []byte(poolName)
	addrBytes := []byte(provider.String())

	return CombineKeys(addrBytes, poolBytes)
}

// Takes ProviderRecord struct to generate store key.
// Key format is: {poolName}{provider}
func GetProviderKey(record ProviderRecord) []byte {
	return ProviderRecordKey(record.PoolName, record.Provider)
}

// Takes ProviderRecord struct to generate reference key.
// Key format is: {provider}{provider}
func GetProviderRefKey(record ProviderRecord) []byte {

	poolBytes := []byte(record.PoolName)
	addrBytes := []byte(record.Provider)

	return CombineKeys(addrBytes, poolBytes)
}
