package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ForsaleKeyPrefix is the prefix to retrieve all Forsale
	ForsaleKeyPrefix = "Forsale/value/"
)

// ForsaleKey returns the store key to retrieve a Forsale from the index fields
func ForsaleKey(
	name string,
) []byte {
	var key []byte

	nameBytes := []byte(name)
	key = append(key, nameBytes...)
	key = append(key, []byte("/")...)

	return key
}
