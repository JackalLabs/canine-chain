package types

import (
	"fmt"
)

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		Params:              DefaultParams(),
		FileList:            []UnifiedFile{},
		ProvidersList:       []Providers{},
		PaymentInfoList:     []StoragePaymentInfo{},
		CollateralList:      []Collateral{},
		ActiveProvidersList: []ActiveProviders{},
		ReportForms:         []ReportForm{},
		AttestForms:         []AttestationForm{},
		OracleRequests:      []OracleRequest{},
		OracleEntries:       []OracleEntry{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs *GenesisState) Validate() error {
	// Check for duplicated index in files
	filesIndexMap := make(map[string]struct{})

	for _, elem := range gs.FileList {
		index := string(FilesPrimaryKey(elem.Merkle, elem.Owner, elem.Start))
		if _, ok := filesIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for activeDeals")
		}
		filesIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in providers
	providersIndexMap := make(map[string]struct{})

	for _, elem := range gs.ProvidersList {
		index := string(ProvidersKey(elem.Address))
		if _, ok := providersIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for providers")
		}
		providersIndexMap[index] = struct{}{}
	}

	// Check for duplicated index in fidCid
	payinfoIndexMap := make(map[string]struct{})

	for _, elem := range gs.PaymentInfoList {
		index := string(StoragePaymentInfoKey(elem.Address))
		if _, ok := payinfoIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for address")
		}
		payinfoIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
