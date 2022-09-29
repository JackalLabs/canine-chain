package types

import (
	"bytes"
)

const (
	KeysSeparator = "/"
)

func CombineKeys(keys ...[]byte) []byte {
	return bytes.Join(keys, []byte(KeysSeparator))
}
