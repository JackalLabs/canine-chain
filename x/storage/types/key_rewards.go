package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	RewardsKeyPrefix = "Rewards/value/"
)

func RewardsKey( address string) []byte {
	var key []byte
	addr := []byte(address)

	key = append(key, addr)
	key = append(key, []byte("/")...)

	return key
}
