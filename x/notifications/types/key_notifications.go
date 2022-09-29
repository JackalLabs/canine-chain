package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// NotificationsKeyPrefix is the prefix to retrieve all Notifications
	NotificationsKeyPrefix = "Notifications/value/"
)

// NotificationsKey returns the store key to retrieve a Notifications from the index fields
func NotificationsKey(
	count uint64,
	address string,
) []byte {
	var key []byte

	countBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(countBytes, count)
	userAddress := []byte(address)
	key = append(key, countBytes...)
	key = append(key, []byte("/")...)
	key = append(key, userAddress...)
	key = append(key, []byte("/")...)

	return key
}
