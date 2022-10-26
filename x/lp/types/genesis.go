package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// Returns account address that protocol fees are collected to
// Panics if the address couldn't be converted to sdk.AccAddress
func ProtocolFeeToAcc() sdk.AccAddress {
	addr := "jkl1vvjuz20086ujjwk3rmnfhy677lxn7hrvu4zzfr"
	return sdk.MustAccAddressFromBech32(addr)
}

// Returns protocol fee rate
// Panics if the rate couldn't be converted to sdk.Dec
func ProtocolFeeRate() sdk.Dec {
	rate := "0.001"
	return sdk.MustNewDecFromStr(rate)
}

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		LPoolList:           []LPool{},
		LProviderRecordList: []LProviderRecord{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in lPool
	lPoolIndexMap := make(map[string]struct{})

	for _, elem := range gs.LPoolList {
		index := string(LPoolKey(elem.Index))
		if _, ok := lPoolIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for lPool")
		}
		lPoolIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in lProviderRecord
	lProviderRecordIndexMap := make(map[string]struct{})

	for _, elem := range gs.LProviderRecordList {
		index := string(LProviderRecordKey(elem.PoolName, elem.Provider))
		if _, ok := lProviderRecordIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for lProviderRecord")
		}
		lProviderRecordIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
