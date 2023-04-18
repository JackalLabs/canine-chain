package types

import (
	"encoding/binary"
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
	poolId uint64,
	provider string,
) []byte {

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, poolId)
	addrBytes := []byte(provider)

	return CombineKeys(bz, addrBytes)
}

// ProviderRecordKey returns the store key to retrieve a ProviderRecord
// reference.
func ProviderRecordRefKey(
	poolId uint64,
	provider string,
) []byte {
	
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, poolId)
	addrBytes := []byte(provider)

	return CombineKeys(addrBytes, bz)
}

// Takes ProviderRecord struct to generate store key.
// Key format is: {poolName}{provider}
func GetProviderKey(record ProviderRecord) []byte {
	return ProviderRecordKey(record.PoolId, record.Provider)
}

// Takes ProviderRecord struct to generate reference key.
// Key format is: {provider}{provider}
func GetProviderRefKey(record ProviderRecord) []byte {
	return ProviderRecordRefKey(record.PoolId, record.Provider)
}
