package types

import (
	"encoding/binary"
	"fmt"
)

var _ binary.ByteOrder

const (
	OracleRequestKey = "OracleRequest/value/"
	OracleEntryKey   = "OracleEntry/value/"
)

// RequestKey returns the store key to retrieve an oracle request from the index fields ordered by merkle
func RequestKey(
	owner string,
	merkle []byte,
	chunk int64,
) []byte {
	return []byte(fmt.Sprintf("%x/%d/%s/", merkle, chunk, owner))
}

// EntryKey returns the store key to retrieve an oracle entry from the index fields ordered by merkle
func EntryKey(
	owner string,
	merkle []byte,
	chunk int64,
) []byte {
	return []byte(fmt.Sprintf("%x/%d/%s/", merkle, chunk, owner))
}
