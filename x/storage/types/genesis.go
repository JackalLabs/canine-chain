package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ContractsList:   []Contracts{},
		ProofsList:      []Proofs{},
		ActiveDealsList: []ActiveDeals{},
		ProvidersList:   []Providers{},
		PayBlocksList:   []PayBlocks{},
		ClientUsageList: []ClientUsage{},
		StraysList:      []Strays{},
		FidCidList:      []FidCid{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in contracts
	contractsIndexMap := make(map[string]struct{})

	for _, elem := range gs.ContractsList {
		index := string(ContractsKey(elem.Cid))
		if _, ok := contractsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for contracts")
		}
		contractsIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in proofs
	proofsIndexMap := make(map[string]struct{})

	for _, elem := range gs.ProofsList {
		index := string(ProofsKey(elem.Cid))
		if _, ok := proofsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for proofs")
		}
		proofsIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in activeDeals
	activeDealsIndexMap := make(map[string]struct{})

	for _, elem := range gs.ActiveDealsList {
		index := string(ActiveDealsKey(elem.Cid))
		if _, ok := activeDealsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for activeDeals")
		}
		activeDealsIndexMap[index] = struct{}{}
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
	// Check for duplicated index in payBlocks
	payBlocksIndexMap := make(map[string]struct{})

	for _, elem := range gs.PayBlocksList {
		index := string(PayBlocksKey(elem.Blockid))
		if _, ok := payBlocksIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for payBlocks")
		}
		payBlocksIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in clientUsage
	clientUsageIndexMap := make(map[string]struct{})

	for _, elem := range gs.ClientUsageList {
		index := string(ClientUsageKey(elem.Address))
		if _, ok := clientUsageIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for clientUsage")
		}
		clientUsageIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in strays
	straysIndexMap := make(map[string]struct{})

	for _, elem := range gs.StraysList {
		index := string(StraysKey(elem.Cid))
		if _, ok := straysIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for strays")
		}
		straysIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in fidCid
	fidCidIndexMap := make(map[string]struct{})

	for _, elem := range gs.FidCidList {
		index := string(FidCidKey(elem.Fid))
		if _, ok := fidCidIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for fidCid")
		}
		fidCidIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
