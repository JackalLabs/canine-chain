package storage

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	storagesimulation "github.com/jackalLabs/canine-chain/v3/x/storage/simulation"
)

// avoid unused import issue
var (
	// _ = sample.AccAddress
	_ = storagesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	//nolint:all
	opWeightMsgSetProviderIP = "op_weight_msg_set_provider_ip"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetProviderIP int = 10

	//nolint:all
	opWeightMsgSetProviderTotalSpace = "op_weight_msg_set_provider_totalspace"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetProviderTotalSpace int = 10

	//nolint:all
	opWeightMsgInitProvider = "op_weight_msg_init_provider"
	// TODO: Determine the simulation weight value
	defaultWeightMsgInitProvider int = 60

	//nolint:all
	opWeightMsgBuyStorage = "op_weight_msg_buy_storage"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBuyStorage int = 100

	//nolint:all
	opWeightMsgAddProviderClaimer          = "op_weight_msg_add_provider_claimer"
	defaultWeightMsgAddProviderClaimer int = 100

	//nolint:all
	opWeightMsgRemoveProviderClaimer          = "op_weight_msg_remove_provider_claimer"
	defaultWeightMsgRemoveProviderClaimer int = 10
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	storagesimulation.RandomizedGenState(simState)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(r *rand.Rand) []simtypes.ParamChange {
	return storagesimulation.ParamChanges(r)
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgSetProviderIP int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSetProviderIP, &weightMsgSetProviderIP, nil,
		func(_ *rand.Rand) {
			weightMsgSetProviderIP = defaultWeightMsgSetProviderIP
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetProviderIP,
		storagesimulation.SimulateMsgSetProviderIP(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetProviderTotalSpace int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSetProviderTotalSpace, &weightMsgSetProviderTotalSpace, nil,
		func(_ *rand.Rand) {
			weightMsgSetProviderTotalSpace = defaultWeightMsgSetProviderTotalSpace
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetProviderTotalSpace,
		storagesimulation.SimulateMsgSetProviderTotalSpace(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgInitProvider int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgInitProvider, &weightMsgInitProvider, nil,
		func(_ *rand.Rand) {
			weightMsgInitProvider = defaultWeightMsgInitProvider
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgInitProvider,
		storagesimulation.SimulateMsgInitProvider(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgBuyStorage int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgBuyStorage, &weightMsgBuyStorage, nil,
		func(_ *rand.Rand) {
			weightMsgBuyStorage = defaultWeightMsgBuyStorage
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBuyStorage,
		storagesimulation.SimulateMsgBuyStorage(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddProviderClaimer int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddProviderClaimer, &weightMsgAddProviderClaimer, nil,
		func(_ *rand.Rand) {
			weightMsgAddProviderClaimer = defaultWeightMsgAddProviderClaimer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddProviderClaimer,
		storagesimulation.SimulateMsgAddProviderClaimer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRemoveProviderClaimer int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRemoveProviderClaimer, &weightMsgRemoveProviderClaimer, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveProviderClaimer = defaultWeightMsgRemoveProviderClaimer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRemoveProviderClaimer,
		storagesimulation.SimulateMsgRemoveProviderClaimer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	return operations
}
