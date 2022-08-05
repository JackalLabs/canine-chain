package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// FormKeyPrefix is the prefix to retrieve all Form
	FormKeyPrefix = "Form/value/"
)

// FormKey returns the store key to retrieve a Form from the index fields
func FormKey(
	ffid string,
) []byte {
	var key []byte

	ffidBytes := []byte(ffid)
	key = append(key, ffidBytes...)
	key = append(key, []byte("/")...)

	return key
}
