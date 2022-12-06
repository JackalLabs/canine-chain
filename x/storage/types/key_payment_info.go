package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// StoragePaymentInfoKeyPrefix is the prefix to retrieve all StoragePaymentInfo
	StoragePaymentInfoKeyPrefix = "StoragePaymentInfo/value/"
)

// StoragePaymentInfoKey returns the store key to retrieve a StoragePaymentInfo from the index fields
func StoragePaymentInfoKey(
	blockid string,
) []byte {
	var key []byte

	blockidBytes := []byte(blockid)
	key = append(key, blockidBytes...)
	key = append(key, []byte("/")...)

	return key
}
