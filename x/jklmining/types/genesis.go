package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		SaveRequestsList: []SaveRequests{},
		MinersList:       []Miners{},
		MinedList:        []Mined{},
		MinerClaimsList:  []MinerClaims{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in saveRequests
	saveRequestsIndexMap := make(map[string]struct{})

	for _, elem := range gs.SaveRequestsList {
		index := string(SaveRequestsKey(elem.Index))
		if _, ok := saveRequestsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for saveRequests")
		}
		saveRequestsIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in miners
	minersIndexMap := make(map[string]struct{})

	for _, elem := range gs.MinersList {
		index := string(MinersKey(elem.Address))
		if _, ok := minersIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for miners")
		}
		minersIndexMap[index] = struct{}{}
	}
	// Check for duplicated ID in mined
	minedIdMap := make(map[uint64]bool)
	minedCount := gs.GetMinedCount()
	for _, elem := range gs.MinedList {
		if _, ok := minedIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for mined")
		}
		if elem.Id >= minedCount {
			return fmt.Errorf("mined id should be lower or equal than the last id")
		}
		minedIdMap[elem.Id] = true
	}
	// Check for duplicated index in minerClaims
	minerClaimsIndexMap := make(map[string]struct{})

	for _, elem := range gs.MinerClaimsList {
		index := string(MinerClaimsKey(elem.Hash))
		if _, ok := minerClaimsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for minerClaims")
		}
		minerClaimsIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
