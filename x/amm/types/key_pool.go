package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PoolKeyPrefix is the prefix to retrieve all LPool
	PoolKeyPrefix = "Pool/value/"
)

// PoolKey returns the store key to retrieve a Pool from the index fields
func PoolKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
