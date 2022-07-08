package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// BidsKeyPrefix is the prefix to retrieve all Bids
	BidsKeyPrefix = "Bids/value/"
)

// BidsKey returns the store key to retrieve a Bids from the index fields
func BidsKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
