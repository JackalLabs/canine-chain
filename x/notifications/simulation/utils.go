package simulation

import (
	"encoding/json"

	"github.com/jackalLabs/canine-chain/x/notifications/types"
)

func GetBlockedSenders(counter types.NotiCounter) (address []string) {
	blackList := make([]string, 100)
	err := json.Unmarshal([]byte(counter.BlockedSenders), &blackList)
	if err != nil {
		panic(err)
	}

	return blackList
}

func IsBlocked(address string, counter types.NotiCounter) bool {
	blackList := GetBlockedSenders(counter)

	for _, addr := range blackList {
		if addr == address {
			return true
		}
	}

	return false
}
