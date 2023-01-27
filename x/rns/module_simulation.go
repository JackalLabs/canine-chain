package rns

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	//	"github.com/jackalLabs/canine-chain/testutil/sample"
	rnssimulation "github.com/jackalLabs/canine-chain/x/rns/simulation"
	"github.com/jackalLabs/canine-chain/x/rns/types"
)

// TODO: rewrite tests but don't use ignite

// avoid unused import issue
var (
	//	_ = sample.AccAddress
	_ = rnssimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

//nolint:gosec // these aren't hard-coded credentials
const (
	opWeightMsgRegister = "op_weight_msg_register"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegister int = 100

	opWeightMsgBid = "op_weight_msg_bid"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBid int = 100

	opWeightMsgAcceptBid = "op_weight_msg_accept_bid"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAcceptBid int = 0

	opWeightMsgCancelBid = "op_weight_msg_cancel_bid"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCancelBid int = 0

	opWeightMsgList = "op_weight_msg_list"
	// TODO: Determine the simulation weight value
	defaultWeightMsgList int = 100

	opWeightMsgBuy = "op_weight_msg_buy"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBuy int = 0

	opWeightMsgDelist = "op_weight_msg_delist"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDelist int = 0

	opWeightMsgTransfer = "op_weight_msg_transfer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTransfer int = 100
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	rnsGenesis := types.GenesisState{}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&rnsGenesis)
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

	var weightMsgRegister int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRegister, &weightMsgRegister, nil,
		func(_ *rand.Rand) {
			weightMsgRegister = defaultWeightMsgRegister
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegister,
		rnssimulation.SimulateMsgRegister(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgBid int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgBid, &weightMsgBid, nil,
		func(_ *rand.Rand) {
			weightMsgBid = defaultWeightMsgBid
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBid,
		rnssimulation.SimulateMsgBid(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAcceptBid int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAcceptBid, &weightMsgAcceptBid, nil,
		func(_ *rand.Rand) {
			weightMsgAcceptBid = defaultWeightMsgAcceptBid
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAcceptBid,
		rnssimulation.SimulateMsgAcceptBid(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCancelBid int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCancelBid, &weightMsgCancelBid, nil,
		func(_ *rand.Rand) {
			weightMsgCancelBid = defaultWeightMsgCancelBid
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCancelBid,
		rnssimulation.SimulateMsgCancelBid(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgList int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgList, &weightMsgList, nil,
		func(_ *rand.Rand) {
			weightMsgList = defaultWeightMsgList
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgList,
		rnssimulation.SimulateMsgList(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgBuy int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgBuy, &weightMsgBuy, nil,
		func(_ *rand.Rand) {
			weightMsgBuy = defaultWeightMsgBuy
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBuy,
		rnssimulation.SimulateMsgBuy(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDelist int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDelist, &weightMsgDelist, nil,
		func(_ *rand.Rand) {
			weightMsgDelist = defaultWeightMsgDelist
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDelist,
		rnssimulation.SimulateMsgDelist(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgTransfer int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgTransfer, &weightMsgTransfer, nil,
		func(_ *rand.Rand) {
			weightMsgTransfer = defaultWeightMsgTransfer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTransfer,
		rnssimulation.SimulateMsgTransfer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	return operations
}
