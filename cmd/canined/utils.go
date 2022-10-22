package main

import "fmt"

const UPTIME_LEFT_KEY = "UPTL-"
const FILE_KEY = "FILE-"
const DOWNTIME_KEY = "DWNT-"

func makeUptimeKey(cid string) []byte {
	return []byte(fmt.Sprintf("%s%s", UPTIME_LEFT_KEY, cid))
}

func makeFileKey(cid string) []byte {
	return []byte(fmt.Sprintf("%s%s", FILE_KEY, cid))
}

func makeDowntimeKey(cid string) []byte {
	return []byte(fmt.Sprintf("%s%s", DOWNTIME_KEY, cid))
}
