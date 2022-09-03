package types

import (
	"encoding/binary"
	"fmt"
)

var _ binary.ByteOrder

const (
	// NamesKeyPrefix is the prefix to retrieve all Names
	NamesKeyPrefix = "Names/value/"
)

// NamesKey returns the store key to retrieve a Names from the index fields
func NamesKey(
	name string,
	tld string,
) []byte {
	var key []byte

	indexBytes := []byte(fmt.Sprintf("%s.%s", name, tld))
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
