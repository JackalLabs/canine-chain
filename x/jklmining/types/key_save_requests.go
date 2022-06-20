package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// SaveRequestsKeyPrefix is the prefix to retrieve all SaveRequests
	SaveRequestsKeyPrefix = "SaveRequests/value/"
)

// SaveRequestsKey returns the store key to retrieve a SaveRequests from the index fields
func SaveRequestsKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
