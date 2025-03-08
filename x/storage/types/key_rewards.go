package types

import (
	"encoding/binary"
	"strings"
)

var _ binary.ByteOrder

const (
	ScoreKeyPrefix = "Score/value/"
)

func ScoreKey(address string) []byte {
	var key []byte
	addr := []byte(address)

	key = append(key, addr...)
	key = append(key, []byte("/")...)

	return key
}

func AddressFromScoreKey(key []byte) string {
	s := strings.Split(string(key), "/")
	if len(s) < 3 {
		return ""
	}
	return s[2]
}
