package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// InitKeyPrefix is the prefix to retrieve all Init
	InitKeyPrefix = "Init/value/"
)

// InitKey returns the store key to retrieve a Init from the index fields
func InitKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
