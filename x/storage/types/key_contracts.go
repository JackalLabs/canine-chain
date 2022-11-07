package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ContractsKeyPrefix is the prefix to retrieve all Contracts
	ContractsKeyPrefix = "Contracts/value/"
)

// ContractsKey returns the store key to retrieve a Contracts from the index fields
func ContractsKey(
	cid string,
) []byte {
	var key []byte

	cidBytes := []byte(cid)
	key = append(key, cidBytes...)
	key = append(key, []byte("/")...)

	return key
}
