package types

import "fmt"

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:   DefaultParams(),
		FeedList: []Feed{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in whois
	feedMap := make(map[string]struct{})

	for _, elem := range gs.FeedList {
		index := string(FeedKey(elem.Name))
		if _, ok := feedMap[index]; ok {
			return fmt.Errorf("duplicated index for whois")
		}
		feedMap[index] = struct{}{}
	}

	return gs.Params.Validate()
}
