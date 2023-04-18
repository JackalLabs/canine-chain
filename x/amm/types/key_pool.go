package types

import (
	"encoding/binary"
)

var _ binary.ByteOrder

const (
	// PoolKeyPrefix is the prefix to retrieve all LPool
	PoolKeyPrefix = "Pool/value/"
)

// PoolKey returns the store key to retrieve a Pool from the index fields
func PoolKey(
	id uint64,
) []byte {
	var key []byte
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)

	key = append(key, bz...)
	key = append(key, []byte("/")...)

	return key
}
