package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ProofsKeyPrefix is the prefix to retrieve all Proofs
	ProofsKeyPrefix = "Proofs/value/"
)

// ProofsKey returns the store key to retrieve a Proofs from the index fields
func ProofsKey(
	cid string,
) []byte {
	var key []byte

	cidBytes := []byte(cid)
	key = append(key, cidBytes...)
	key = append(key, []byte("/")...)

	return key
}
