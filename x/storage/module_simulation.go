package storage

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	storagesimulation "github.com/jackalLabs/canine-chain/x/storage/simulation"
	"github.com/jackalLabs/canine-chain/x/storage/types"
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
	opWeightMsgPostContract = "op_weight_msg_post_contract"
	// TODO: Determine the simulation weight value
	defaultWeightMsgPostContract int = 100

	//nolint:all
	opWeightMsgPostproof = "op_weight_msg_postproof"
	// TODO: Determine the simulation weight value
	defaultWeightMsgPostproof int = 100

	//nolint:all
	opWeightMsgSignContract = "op_weight_msg_sign_contract"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSignContract int = 100

	//nolint:all
	opWeightMsgSetProviderIP = "op_weight_msg_set_provider_ip"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetProviderIP int = 100

	//nolint:all
	opWeightMsgSetProviderTotalspace = "op_weight_msg_set_provider_totalspace"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetProviderTotalspace int = 100

	//nolint:all
	opWeightMsgInitProvider = "op_weight_msg_init_provider"
	// TODO: Determine the simulation weight value
	defaultWeightMsgInitProvider int = 100

	//nolint:all
	opWeightMsgCancelContract = "op_weight_msg_cancel_contract"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCancelContract int = 100

	//nolint:all
	opWeightMsgBuyStorage = "op_weight_msg_buy_storage"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBuyStorage int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	storageGenesis := types.GenesisState{
		ContractsList: []types.Contracts{
			// {
			// 	Creator: sample.AccAddress(),
			// 	Cid:     "0",
			// },
			// {
			// 	Creator: sample.AccAddress(),
			// 	Cid:     "1",
			// },
		},
		ActiveDealsList: []types.ActiveDeals{
			// {
			// 	Creator: sample.AccAddress(),
			// 	Cid:     "0",
			// },
			// {
			// 	Creator: sample.AccAddress(),
			// 	Cid:     "1",
			// },
		},
		ProvidersList: []types.Providers{
			// {
			// 	Creator: sample.AccAddress(),
			// 	Address: "0",
			// },
			// {
			// 	Creator: sample.AccAddress(),
			// 	Address: "1",
			// },
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&storageGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgPostContract int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgPostContract, &weightMsgPostContract, nil,
		func(_ *rand.Rand) {
			weightMsgPostContract = defaultWeightMsgPostContract
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPostContract,
		storagesimulation.SimulateMsgPostContract(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgPostproof int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgPostproof, &weightMsgPostproof, nil,
		func(_ *rand.Rand) {
			weightMsgPostproof = defaultWeightMsgPostproof
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPostproof,
		storagesimulation.SimulateMsgPostproof(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSignContract int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSignContract, &weightMsgSignContract, nil,
		func(_ *rand.Rand) {
			weightMsgSignContract = defaultWeightMsgSignContract
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSignContract,
		storagesimulation.SimulateMsgSignContract(am.accountKeeper, am.bankKeeper, am.keeper),
	))

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

	var weightMsgSetProviderTotalspace int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSetProviderTotalspace, &weightMsgSetProviderTotalspace, nil,
		func(_ *rand.Rand) {
			weightMsgSetProviderTotalspace = defaultWeightMsgSetProviderTotalspace
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetProviderTotalspace,
		storagesimulation.SimulateMsgSetProviderTotalspace(am.accountKeeper, am.bankKeeper, am.keeper),
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

	var weightMsgCancelContract int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCancelContract, &weightMsgCancelContract, nil,
		func(_ *rand.Rand) {
			weightMsgCancelContract = defaultWeightMsgCancelContract
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCancelContract,
		storagesimulation.SimulateMsgCancelContract(am.accountKeeper, am.bankKeeper, am.keeper),
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

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
