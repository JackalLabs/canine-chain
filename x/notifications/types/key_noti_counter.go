package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// NotiCounterKeyPrefix is the prefix to retrieve all NotiCounter
	NotiCounterKeyPrefix = "NotiCounter/value/"
)

// NotiCounterKey returns the store key to retrieve a NotiCounter from the index fields
func NotiCounterKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
