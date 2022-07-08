package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// WhoisKeyPrefix is the prefix to retrieve all Whois
	WhoisKeyPrefix = "Whois/value/"
)

// WhoisKey returns the store key to retrieve a Whois from the index fields
func WhoisKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
