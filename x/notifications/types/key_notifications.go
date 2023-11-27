package types

import (
	"encoding/binary"
	"fmt"
	"time"
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
	time time.Time,
) []byte {

	return []byte(fmt.Sprintf("%s/%s/%d", to, from, time.Unix()))
}
