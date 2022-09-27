package lp

import (
	"math/rand"

	"github.com/jackal-dao/canine/testutil/sample"
	lpsimulation "github.com/jackal-dao/canine/x/lp/simulation"
	"github.com/jackal-dao/canine/x/lp/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = lpsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateLPool = "op_weight_msg_create_l_pool"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateLPool int = 100

	opWeightMsgDepositLPool = "op_weight_msg_deposit_l_pool"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDepositLPool int = 100

	opWeightMsgWithdrawLPool = "op_weight_msg_withdraw_l_pool"
	// TODO: Determine the simulation weight value
	defaultWeightMsgWithdrawLPool int = 100

	opWeightMsgSwap = "op_weight_msg_swap"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSwap int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	lpGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&lpGenesis)
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

	var weightMsgCreateLPool int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateLPool, &weightMsgCreateLPool, nil,
		func(_ *rand.Rand) {
			weightMsgCreateLPool = defaultWeightMsgCreateLPool
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateLPool,
		lpsimulation.SimulateMsgCreateLPool(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDepositLPool int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDepositLPool, &weightMsgDepositLPool, nil,
		func(_ *rand.Rand) {
			weightMsgDepositLPool = defaultWeightMsgDepositLPool
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDepositLPool,
		lpsimulation.SimulateMsgDepositLPool(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgWithdrawLPool int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgWithdrawLPool, &weightMsgWithdrawLPool, nil,
		func(_ *rand.Rand) {
			weightMsgWithdrawLPool = defaultWeightMsgWithdrawLPool
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWithdrawLPool,
		lpsimulation.SimulateMsgWithdrawLPool(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSwap int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSwap, &weightMsgSwap, nil,
		func(_ *rand.Rand) {
			weightMsgSwap = defaultWeightMsgSwap
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSwap,
		lpsimulation.SimulateMsgSwap(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
