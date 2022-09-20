package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// FilesKeyPrefix is the prefix to retrieve all Files
	FilesKeyPrefix = "Files/value/"
)

// FilesKey returns the store key to retrieve a Files from the index fields
func FilesKey(
	address string,
	ownerAddress string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	ownerAddressBytes := []byte(ownerAddress)

	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)
	key = append(key, ownerAddressBytes...)
	key = append(key, []byte("/")...)

	return key
}
