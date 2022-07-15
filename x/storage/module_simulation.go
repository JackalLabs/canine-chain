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

	opWeightMsgCreateMiners = "op_weight_msg_miners"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateMiners int = 100

	opWeightMsgUpdateMiners = "op_weight_msg_miners"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateMiners int = 100

	opWeightMsgDeleteMiners = "op_weight_msg_miners"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteMiners int = 100

	opWeightMsgSetMinerIp = "op_weight_msg_set_miner_ip"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetMinerIp int = 100

	opWeightMsgSetMinerTotalspace = "op_weight_msg_set_miner_totalspace"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetMinerTotalspace int = 100

	opWeightMsgInitMiner = "op_weight_msg_init_miner"
	// TODO: Determine the simulation weight value
	defaultWeightMsgInitMiner int = 100

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
		MinersList: []types.Miners{
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

	var weightMsgCreateMiners int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateMiners, &weightMsgCreateMiners, nil,
		func(_ *rand.Rand) {
			weightMsgCreateMiners = defaultWeightMsgCreateMiners
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateMiners,
		storagesimulation.SimulateMsgCreateMiners(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateMiners int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateMiners, &weightMsgUpdateMiners, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateMiners = defaultWeightMsgUpdateMiners
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateMiners,
		storagesimulation.SimulateMsgUpdateMiners(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteMiners int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteMiners, &weightMsgDeleteMiners, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteMiners = defaultWeightMsgDeleteMiners
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteMiners,
		storagesimulation.SimulateMsgDeleteMiners(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetMinerIp int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSetMinerIp, &weightMsgSetMinerIp, nil,
		func(_ *rand.Rand) {
			weightMsgSetMinerIp = defaultWeightMsgSetMinerIp
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetMinerIp,
		storagesimulation.SimulateMsgSetMinerIp(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetMinerTotalspace int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSetMinerTotalspace, &weightMsgSetMinerTotalspace, nil,
		func(_ *rand.Rand) {
			weightMsgSetMinerTotalspace = defaultWeightMsgSetMinerTotalspace
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetMinerTotalspace,
		storagesimulation.SimulateMsgSetMinerTotalspace(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgInitMiner int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgInitMiner, &weightMsgInitMiner, nil,
		func(_ *rand.Rand) {
			weightMsgInitMiner = defaultWeightMsgInitMiner
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgInitMiner,
		storagesimulation.SimulateMsgInitMiner(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
