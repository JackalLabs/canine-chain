package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PayBlocksKeyPrefix is the prefix to retrieve all PayBlocks
	PayBlocksKeyPrefix = "PayBlocks/value/"
)

// PayBlocksKey returns the store key to retrieve a PayBlocks from the index fields
func PayBlocksKey(
	blockid string,
) []byte {
	var key []byte

	blockidBytes := []byte(blockid)
	key = append(key, blockidBytes...)
	key = append(key, []byte("/")...)

	return key
}
