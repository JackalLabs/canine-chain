package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PubkeyKeyPrefix is the prefix to retrieve all Pubkey
	PubkeyKeyPrefix = "Pubkey/value/"
)

// PubkeyKey returns the store key to retrieve a Pubkey from the index fields
func PubkeyKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
