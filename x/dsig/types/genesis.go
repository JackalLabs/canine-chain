package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		UserUploadsList: []UserUploads{},
		FormList:        []Form{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in userUploads
	userUploadsIndexMap := make(map[string]struct{})

	for _, elem := range gs.UserUploadsList {
		index := string(UserUploadsKey(elem.Fid))
		if _, ok := userUploadsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for userUploads")
		}
		userUploadsIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in form
	formIndexMap := make(map[string]struct{})

	for _, elem := range gs.FormList {
		index := string(FormKey(elem.Ffid))
		if _, ok := formIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for form")
		}
		formIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
