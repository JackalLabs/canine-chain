package filetree

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	filetreesimulation "github.com/jackalLabs/canine-chain/v5/x/filetree/simulation"
	"github.com/jackalLabs/canine-chain/v5/x/filetree/types"
)

// avoid unused import issue
var (
	// _ = sample.AccAddress
	_ = filetreesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	// opWeightMsgPostFile = "op_weight_msg_post_file"
	// TODO: Determine the simulation weight value
	defaultWeightMsgPostFile int = 100

	// opWeightMsgAddViewers = "op_weight_msg_add_viewers"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddViewers int = 100

	// opWeightMsgPostkey = "op_weight_msg_postkey"
	// TODO: Determine the simulation weight value
	defaultWeightMsgPostkey int = 100

	// opWeightMsgInitAccount = "op_weight_msg_init_account"
	// TODO: Determine the simulation weight value
	// defaultWeightMsgInitAccount int = 100

	// opWeightMsgDeleteFile = "op_weight_msg_delete_file"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteFile int = 100

	// opWeightMsgRemoveViewers = "op_weight_msg_remove_viewers"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRemoveViewers int = 100

	// opWeightMsgMakeRoot = "op_weight_msg_make_root"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMakeRoot int = 100

	// opWeightMsgAddEditors = "op_weight_msg_add_editors"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddEditors int = 100

	// opWeightMsgRemoveEditors = "op_weight_msg_remove_editors"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRemoveEditors int = 100

	// opWeightMsgResetEditors = "op_weight_msg_reset_editors"
	// TODO: Determine the simulation weight value
	defaultWeightMsgResetEditors int = 100

	// opWeightMsgResetViewers = "op_weight_msg_reset_viewers"
	// TODO: Determine the simulation weight value
	defaultWeightMsgResetViewers int = 100

	// opWeightMsgChangeOwner = "op_weight_msg_change_owner"
	// TODO: Determine the simulation weight value
	defaultWeightMsgChangeOwner int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	filetreeGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&filetreeGenesis)
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

	var weightMsgPostFile int
	simState.AppParams.GetOrGenerate(simState.Cdc, "opWeightMsgPostFile", &weightMsgPostFile, nil,
		func(_ *rand.Rand) {
			weightMsgPostFile = defaultWeightMsgPostFile
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPostFile,
		filetreesimulation.SimulateMsgPostFile(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddViewers int
	simState.AppParams.GetOrGenerate(simState.Cdc, "op_weight_msg_post_file", &weightMsgAddViewers, nil,
		func(_ *rand.Rand) {
			weightMsgAddViewers = defaultWeightMsgAddViewers
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddViewers,
		filetreesimulation.SimulateMsgAddViewers(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgPostkey int
	simState.AppParams.GetOrGenerate(simState.Cdc, "op_weight_msg_postkey", &weightMsgPostkey, nil,
		func(_ *rand.Rand) {
			weightMsgPostkey = defaultWeightMsgPostkey
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPostkey,
		filetreesimulation.SimulateMsgPostkey(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteFile int
	simState.AppParams.GetOrGenerate(simState.Cdc, "op_weight_msg_delete_file", &weightMsgDeleteFile, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteFile = defaultWeightMsgDeleteFile
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteFile,
		filetreesimulation.SimulateMsgDeleteFile(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRemoveViewers int
	simState.AppParams.GetOrGenerate(simState.Cdc, "op_weight_msg_remove_viewers", &weightMsgRemoveViewers, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveViewers = defaultWeightMsgRemoveViewers
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRemoveViewers,
		filetreesimulation.SimulateMsgRemoveViewers(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgMakeFolder int
	simState.AppParams.GetOrGenerate(simState.Cdc, "op_weight_msg_make_root", &weightMsgMakeFolder, nil,
		func(_ *rand.Rand) {
			weightMsgMakeFolder = defaultWeightMsgMakeRoot
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMakeFolder,
		filetreesimulation.SimulateMsgMakeRoot(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddEditors int
	simState.AppParams.GetOrGenerate(simState.Cdc, "op_weight_msg_add_editors", &weightMsgAddEditors, nil,
		func(_ *rand.Rand) {
			weightMsgAddEditors = defaultWeightMsgAddEditors
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddEditors,
		filetreesimulation.SimulateMsgAddEditors(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRemoveEditors int
	simState.AppParams.GetOrGenerate(simState.Cdc, "op_weight_msg_remove_editors", &weightMsgRemoveEditors, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveEditors = defaultWeightMsgRemoveEditors
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRemoveEditors,
		filetreesimulation.SimulateMsgRemoveEditors(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgResetEditors int
	simState.AppParams.GetOrGenerate(simState.Cdc, "op_weight_msg_reset_editors", &weightMsgResetEditors, nil,
		func(_ *rand.Rand) {
			weightMsgResetEditors = defaultWeightMsgResetEditors
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgResetEditors,
		filetreesimulation.SimulateMsgResetEditors(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgResetViewers int
	simState.AppParams.GetOrGenerate(simState.Cdc, "op_weight_msg_reset_viewers", &weightMsgResetViewers, nil,
		func(_ *rand.Rand) {
			weightMsgResetViewers = defaultWeightMsgResetViewers
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgResetViewers,
		filetreesimulation.SimulateMsgResetViewers(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgChangeOwner int
	simState.AppParams.GetOrGenerate(simState.Cdc, "op_weight_msg_change_owner", &weightMsgChangeOwner, nil,
		func(_ *rand.Rand) {
			weightMsgChangeOwner = defaultWeightMsgChangeOwner
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgChangeOwner,
		filetreesimulation.SimulateMsgChangeOwner(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
