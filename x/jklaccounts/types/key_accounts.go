package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// AccountsKeyPrefix is the prefix to retrieve all Accounts
	AccountsKeyPrefix = "Accounts/value/"
)

// AccountsKey returns the store key to retrieve a Accounts from the index fields
func AccountsKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
