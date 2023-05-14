package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// StraysKeyPrefix is the prefix to retrieve all Strays
	StraysKeyPrefix   = "Strays/value/"
	StraysV2KeyPrefix = "StrayV2/value/"
)

// StraysKey returns the store key to retrieve a Strays from the index fields
func StraysKey(
	cid string,
) []byte {
	var key []byte

	cidBytes := []byte(cid)
	key = append(key, cidBytes...)
	key = append(key, []byte("/")...)

	return key
}
