package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// LPoolKeyPrefix is the prefix to retrieve all LPool
	LPoolKeyPrefix = "LPool/value/"
)

// LPoolKey returns the store key to retrieve a LPool from the index fields
func LPoolKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
