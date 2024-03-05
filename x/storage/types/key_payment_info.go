package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// StoragePaymentInfoKeyPrefix is the prefix to retrieve all StoragePaymentInfo
	StoragePaymentInfoKeyPrefix = "StoragePaymentInfo/value/"

	PaymentGaugeKeyPrefix = "PaymentGauge/value/"
)

// StoragePaymentInfoKey returns the store key to retrieve a StoragePaymentInfo from the index fields
func StoragePaymentInfoKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}

// PaymentGaugeKey returns the store key to retrieve a PaymentGauge from the index fields
func PaymentGaugeKey(
	id []byte,
) []byte {
	var key []byte

	key = append(key, id...)
	key = append(key, []byte("/")...)

	return key
}
