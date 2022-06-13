package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// MinersKeyPrefix is the prefix to retrieve all Miners
	MinersKeyPrefix = "Miners/value/"
)

// MinersKey returns the store key to retrieve a Miners from the index fields
func MinersKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
