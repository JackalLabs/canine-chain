package types

import (
	"fmt"
)

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Notifications: []Notification{},
		Params:        DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in notifications
	notificationsIndexMap := make(map[string]struct{})

	for _, elem := range gs.Notifications {
		index := string(NotificationsKey(elem.To, elem.From, elem.Time))
		if _, ok := notificationsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for notifications")
		}
		notificationsIndexMap[index] = struct{}{}
	}
	return gs.Params.Validate()
}
