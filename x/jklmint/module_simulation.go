package jklmint

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	jklmintsimulation "github.com/jackalLabs/canine-chain/v5/x/jklmint/simulation"
	"github.com/jackalLabs/canine-chain/v5/x/jklmint/types"
)

// avoid unused import issue
var (
	//	_ = sample.AccAddress
	_ = jklmintsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	jklmintGenesis := types.GenesisState{
		// Params: types.NewParams(sdk.DefaultBondDenom, 4, 6),
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&jklmintGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	jklmintParams := types.DefaultParams()
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyMintDenom), func(_ *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(jklmintParams.MintDenom))
		}),
	}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	_ = simState
	operations := make([]simtypes.WeightedOperation, 0)

	return operations
}
