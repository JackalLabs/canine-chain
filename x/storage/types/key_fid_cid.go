package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// FidCidKeyPrefix is the prefix to retrieve all FidCid
	FidCidKeyPrefix = "FidCid/value/"
)

// FidCidKey returns the store key to retrieve a FidCid from the index fields
func FidCidKey(
	fid string,
) []byte {
	var key []byte

	fidBytes := []byte(fid)
	key = append(key, fidBytes...)
	key = append(key, []byte("/")...)

	return key
}
