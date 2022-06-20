package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// MinerClaimsKeyPrefix is the prefix to retrieve all MinerClaims
	MinerClaimsKeyPrefix = "MinerClaims/value/"
)

// MinerClaimsKey returns the store key to retrieve a MinerClaims from the index fields
func MinerClaimsKey(
	hash string,
) []byte {
	var key []byte

	hashBytes := []byte(hash)
	key = append(key, hashBytes...)
	key = append(key, []byte("/")...)

	return key
}
