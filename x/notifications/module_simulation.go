package notifications

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/jackal-dao/canine/testutil/sample"
	notificationssimulation "github.com/jackal-dao/canine/x/notifications/simulation"
	"github.com/jackal-dao/canine/x/notifications/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = notificationssimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateNotifications = "op_weight_msg_notifications"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateNotifications int = 100

	opWeightMsgUpdateNotifications = "op_weight_msg_notifications"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateNotifications int = 100

	opWeightMsgDeleteNotifications = "op_weight_msg_notifications"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteNotifications int = 100

	opWeightMsgSetCounter = "op_weight_msg_set_counter"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetCounter int = 100

	opWeightMsgAddSenders = "op_weight_msg_add_senders"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddSenders int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	notificationsGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		NotificationsList: []types.Notifications{
			{
				Sender: sample.AccAddress(),
				Count:  0,
			},
			{
				Sender: sample.AccAddress(),
				Count:  1,
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&notificationsGenesis)
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

	var weightMsgCreateNotifications int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateNotifications, &weightMsgCreateNotifications, nil,
		func(_ *rand.Rand) {
			weightMsgCreateNotifications = defaultWeightMsgCreateNotifications
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateNotifications,
		notificationssimulation.SimulateMsgCreateNotifications(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateNotifications int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateNotifications, &weightMsgUpdateNotifications, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateNotifications = defaultWeightMsgUpdateNotifications
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateNotifications,
		notificationssimulation.SimulateMsgUpdateNotifications(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteNotifications int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteNotifications, &weightMsgDeleteNotifications, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteNotifications = defaultWeightMsgDeleteNotifications
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteNotifications,
		notificationssimulation.SimulateMsgDeleteNotifications(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetCounter int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSetCounter, &weightMsgSetCounter, nil,
		func(_ *rand.Rand) {
			weightMsgSetCounter = defaultWeightMsgSetCounter
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetCounter,
		notificationssimulation.SimulateMsgSetCounter(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddSenders int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddSenders, &weightMsgAddSenders, nil,
		func(_ *rand.Rand) {
			weightMsgAddSenders = defaultWeightMsgAddSenders
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddSenders,
		notificationssimulation.SimulateMsgAddSenders(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
