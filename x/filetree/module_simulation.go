package filetree

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/jackal-dao/canine/testutil/sample"
	filetreesimulation "github.com/jackal-dao/canine/x/filetree/simulation"
	"github.com/jackal-dao/canine/x/filetree/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = filetreesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgPostFile = "op_weight_msg_post_file"
	// TODO: Determine the simulation weight value
	defaultWeightMsgPostFile int = 100

	opWeightMsgAddViewers = "op_weight_msg_add_viewers"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddViewers int = 100

	opWeightMsgPostkey = "op_weight_msg_postkey"
	// TODO: Determine the simulation weight value
	defaultWeightMsgPostkey int = 100

	opWeightMsgInitAccount = "op_weight_msg_init_account"
	// TODO: Determine the simulation weight value
	defaultWeightMsgInitAccount int = 100

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
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgPostFile, &weightMsgPostFile, nil,
		func(_ *rand.Rand) {
			weightMsgPostFile = defaultWeightMsgPostFile
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPostFile,
		filetreesimulation.SimulateMsgPostFile(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddViewers int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddViewers, &weightMsgAddViewers, nil,
		func(_ *rand.Rand) {
			weightMsgAddViewers = defaultWeightMsgAddViewers
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddViewers,
		filetreesimulation.SimulateMsgAddViewers(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgPostkey int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgPostkey, &weightMsgPostkey, nil,
		func(_ *rand.Rand) {
			weightMsgPostkey = defaultWeightMsgPostkey
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPostkey,
		filetreesimulation.SimulateMsgPostkey(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgInitAccount int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgInitAccount, &weightMsgInitAccount, nil,
		func(_ *rand.Rand) {
			weightMsgInitAccount = defaultWeightMsgInitAccount
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgInitAccount,
		filetreesimulation.SimulateMsgInitAccount(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
