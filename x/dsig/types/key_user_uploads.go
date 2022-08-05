package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// UserUploadsKeyPrefix is the prefix to retrieve all UserUploads
	UserUploadsKeyPrefix = "UserUploads/value/"
)

// UserUploadsKey returns the store key to retrieve a UserUploads from the index fields
func UserUploadsKey(
	fid string,
) []byte {
	var key []byte

	fidBytes := []byte(fid)
	key = append(key, fidBytes...)
	key = append(key, []byte("/")...)

	return key
}
