package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ClientUsageKeyPrefix is the prefix to retrieve all ClientUsage
	ClientUsageKeyPrefix = "ClientUsage/value/"
)

// ClientUsageKey returns the store key to retrieve a ClientUsage from the index fields
func ClientUsageKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
