package main

import "fmt"

const (
	UptimeLeftKey = "UPTL-"
	FileKey       = "FILE-"
	DowntimeKey   = "DWNT-"
)

//nolint:unused
func makeUptimeKey(cid string) []byte {
	return []byte(fmt.Sprintf("%s%s", UptimeLeftKey, cid))
}

//nolint:unused
func makeFileKey(cid string) []byte {
	return []byte(fmt.Sprintf("%s%s", FileKey, cid))
}

//nolint:unused
func makeDowntimeKey(cid string) []byte {
	return []byte(fmt.Sprintf("%s%s", DowntimeKey, cid))
}
