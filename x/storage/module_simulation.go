package storage

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/jackal-dao/canine/testutil/sample"
	storagesimulation "github.com/jackal-dao/canine/x/storage/simulation"
	"github.com/jackal-dao/canine/x/storage/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = storagesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgPostContract = "op_weight_msg_post_contract"
	// TODO: Determine the simulation weight value
	defaultWeightMsgPostContract int = 100

	opWeightMsgCreateContracts = "op_weight_msg_contracts"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateContracts int = 100

	opWeightMsgUpdateContracts = "op_weight_msg_contracts"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateContracts int = 100

	opWeightMsgDeleteContracts = "op_weight_msg_contracts"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteContracts int = 100

	opWeightMsgCreateProofs = "op_weight_msg_proofs"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateProofs int = 100

	opWeightMsgUpdateProofs = "op_weight_msg_proofs"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateProofs int = 100

	opWeightMsgDeleteProofs = "op_weight_msg_proofs"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteProofs int = 100

	opWeightMsgItem = "op_weight_msg_item"
	// TODO: Determine the simulation weight value
	defaultWeightMsgItem int = 100

	opWeightMsgPostproof = "op_weight_msg_postproof"
	// TODO: Determine the simulation weight value
	defaultWeightMsgPostproof int = 100

	opWeightMsgCreateActiveDeals = "op_weight_msg_active_deals"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateActiveDeals int = 100

	opWeightMsgUpdateActiveDeals = "op_weight_msg_active_deals"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateActiveDeals int = 100

	opWeightMsgDeleteActiveDeals = "op_weight_msg_active_deals"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteActiveDeals int = 100

	opWeightMsgSignContract = "op_weight_msg_sign_contract"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSignContract int = 100

	opWeightMsgCreateProviders = "op_weight_msg_providers"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateProviders int = 100

	opWeightMsgUpdateProviders = "op_weight_msg_providers"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateProviders int = 100

	opWeightMsgDeleteProviders = "op_weight_msg_providers"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteProviders int = 100

	opWeightMsgSetProviderIp = "op_weight_msg_set_provider_ip"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetProviderIp int = 100

	opWeightMsgSetProviderTotalspace = "op_weight_msg_set_provider_totalspace"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetProviderTotalspace int = 100

	opWeightMsgInitProvider = "op_weight_msg_init_provider"
	// TODO: Determine the simulation weight value
	defaultWeightMsgInitProvider int = 100

	opWeightMsgCancelContract = "op_weight_msg_cancel_contract"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCancelContract int = 100

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
			{
				Creator: sample.AccAddress(),
				Cid:     "0",
			},
			{
				Creator: sample.AccAddress(),
				Cid:     "1",
			},
		},
		ProofsList: []types.Proofs{
			{
				Creator: sample.AccAddress(),
				Cid:     "0",
			},
			{
				Creator: sample.AccAddress(),
				Cid:     "1",
			},
		},
		ActiveDealsList: []types.ActiveDeals{
			{
				Creator: sample.AccAddress(),
				Cid:     "0",
			},
			{
				Creator: sample.AccAddress(),
				Cid:     "1",
			},
		},
		ProvidersList: []types.Providers{
			{
				Creator: sample.AccAddress(),
				Address: "0",
			},
			{
				Creator: sample.AccAddress(),
				Address: "1",
			},
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

	var weightMsgCreateContracts int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateContracts, &weightMsgCreateContracts, nil,
		func(_ *rand.Rand) {
			weightMsgCreateContracts = defaultWeightMsgCreateContracts
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateContracts,
		storagesimulation.SimulateMsgCreateContracts(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateContracts int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateContracts, &weightMsgUpdateContracts, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateContracts = defaultWeightMsgUpdateContracts
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateContracts,
		storagesimulation.SimulateMsgUpdateContracts(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteContracts int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteContracts, &weightMsgDeleteContracts, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteContracts = defaultWeightMsgDeleteContracts
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteContracts,
		storagesimulation.SimulateMsgDeleteContracts(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateProofs int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateProofs, &weightMsgCreateProofs, nil,
		func(_ *rand.Rand) {
			weightMsgCreateProofs = defaultWeightMsgCreateProofs
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateProofs,
		storagesimulation.SimulateMsgCreateProofs(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateProofs int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateProofs, &weightMsgUpdateProofs, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateProofs = defaultWeightMsgUpdateProofs
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateProofs,
		storagesimulation.SimulateMsgUpdateProofs(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteProofs int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteProofs, &weightMsgDeleteProofs, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteProofs = defaultWeightMsgDeleteProofs
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteProofs,
		storagesimulation.SimulateMsgDeleteProofs(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgItem int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgItem, &weightMsgItem, nil,
		func(_ *rand.Rand) {
			weightMsgItem = defaultWeightMsgItem
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgItem,
		storagesimulation.SimulateMsgItem(am.accountKeeper, am.bankKeeper, am.keeper),
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

	var weightMsgCreateActiveDeals int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateActiveDeals, &weightMsgCreateActiveDeals, nil,
		func(_ *rand.Rand) {
			weightMsgCreateActiveDeals = defaultWeightMsgCreateActiveDeals
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateActiveDeals,
		storagesimulation.SimulateMsgCreateActiveDeals(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateActiveDeals int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateActiveDeals, &weightMsgUpdateActiveDeals, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateActiveDeals = defaultWeightMsgUpdateActiveDeals
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateActiveDeals,
		storagesimulation.SimulateMsgUpdateActiveDeals(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteActiveDeals int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteActiveDeals, &weightMsgDeleteActiveDeals, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteActiveDeals = defaultWeightMsgDeleteActiveDeals
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteActiveDeals,
		storagesimulation.SimulateMsgDeleteActiveDeals(am.accountKeeper, am.bankKeeper, am.keeper),
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

	var weightMsgCreateProviders int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateProviders, &weightMsgCreateProviders, nil,
		func(_ *rand.Rand) {
			weightMsgCreateProviders = defaultWeightMsgCreateProviders
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateProviders,
		storagesimulation.SimulateMsgCreateProviders(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateProviders int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateProviders, &weightMsgUpdateProviders, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateProviders = defaultWeightMsgUpdateProviders
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateProviders,
		storagesimulation.SimulateMsgUpdateProviders(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteProviders int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteProviders, &weightMsgDeleteProviders, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteProviders = defaultWeightMsgDeleteProviders
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteProviders,
		storagesimulation.SimulateMsgDeleteProviders(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetProviderIp int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSetProviderIp, &weightMsgSetProviderIp, nil,
		func(_ *rand.Rand) {
			weightMsgSetProviderIp = defaultWeightMsgSetProviderIp
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetProviderIp,
		storagesimulation.SimulateMsgSetProviderIp(am.accountKeeper, am.bankKeeper, am.keeper),
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
