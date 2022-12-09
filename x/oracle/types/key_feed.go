package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// FeedKeyPrefix is the prefix to retrieve all Feed
	FeedKeyPrefix = "Feed/value/"
)

// FeedKey returns the store key to retrieve a Feed from the index fields
func FeedKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
