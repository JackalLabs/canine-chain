package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ProvidersKeyPrefix is the prefix to retrieve all Providers
	ProvidersKeyPrefix = "Providers/value/"
)

// ProvidersKey returns the store key to retrieve a Providers from the index fields
func ProvidersKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
