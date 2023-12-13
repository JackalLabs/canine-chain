package types

import (
	"encoding/binary"
	"fmt"
)

var _ binary.ByteOrder

const (
	// NotificationsKeyPrefix is the prefix to retrieve all Notifications
	NotificationsKeyPrefix = "Notification/"
)

// NotificationsKey returns the store key to retrieve a Notifications from the index fields
func NotificationsKey(
	to string,
	from string,
	time int64,
) []byte {
	return []byte(fmt.Sprintf("%s/%s/%d", to, from, time))
}

// BlockKey returns the store key to retrieve a block object from the index fields
func BlockKey(
	owner string,
	address string,
) []byte {
	return []byte(fmt.Sprintf("%s/%s", owner, address))
}
