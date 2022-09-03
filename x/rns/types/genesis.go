package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		WhoisList:   []Whois{},
		NamesList:   []Names{},
		BidsList:    []Bids{},
		ForsaleList: []Forsale{},
		InitList:    []Init{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in whois
	whoisIndexMap := make(map[string]struct{})

	for _, elem := range gs.WhoisList {
		index := string(WhoisKey(elem.Index))
		if _, ok := whoisIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for whois")
		}
		whoisIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in names
	namesIndexMap := make(map[string]struct{})

	for _, elem := range gs.NamesList {
		index := string(NamesKey(elem.Name, elem.Tld))
		if _, ok := namesIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for names")
		}
		namesIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in bids
	bidsIndexMap := make(map[string]struct{})

	for _, elem := range gs.BidsList {
		index := string(BidsKey(elem.Index))
		if _, ok := bidsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for bids")
		}
		bidsIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in forsale
	forsaleIndexMap := make(map[string]struct{})

	for _, elem := range gs.ForsaleList {
		index := string(ForsaleKey(elem.Name))
		if _, ok := forsaleIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for forsale")
		}
		forsaleIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in init
	initIndexMap := make(map[string]struct{})

	for _, elem := range gs.InitList {
		index := string(InitKey(elem.Address))
		if _, ok := initIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for init")
		}
		initIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
