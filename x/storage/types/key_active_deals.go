package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ActiveDealsKeyPrefix is the prefix to retrieve all ActiveDeals
	ActiveDealsKeyPrefix   = "ActiveDeals/value/"
	ActiveDealsV2KeyPrefix = "ActiveDealsV2/value/"
)

// ActiveDealsKey returns the store key to retrieve a ActiveDeals from the index fields
func ActiveDealsKey(
	cid string,
) []byte {
	var key []byte

	cidBytes := []byte(cid)
	key = append(key, cidBytes...)
	key = append(key, []byte("/")...)

	return key
}
