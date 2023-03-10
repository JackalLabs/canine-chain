package notifications

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	notificationssimulation "github.com/jackalLabs/canine-chain/x/notifications/simulation"
	"github.com/jackalLabs/canine-chain/x/notifications/types"
)

// avoid unused import issue
var (
	// _ = sample.AccAddress
	_ = notificationssimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

//nolint:gosec // these aren't hard-coded credentials
const (
	opWeightMsgCreateNotifications = "op_weight_msg_notifications"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateNotifications int = 100

	opWeightMsgDeleteNotifications = "op_weight_msg_notifications"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteNotifications int = 100

	opWeightMsgSetCounter = "op_weight_msg_set_counter"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetCounter int = 100

	opWeightMsgBlockSenders = "op_weight_msg_block_senders"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBlockSenders int = 2
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	notificationsGenesis := types.GenesisState{
		Params:            types.DefaultParams(),
		NotificationsList: []types.Notifications{},
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

	var weightMsgBlockSenders int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgBlockSenders, &weightMsgBlockSenders, nil,
		func(_ *rand.Rand) {
			weightMsgBlockSenders = defaultWeightMsgBlockSenders
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBlockSenders,
		notificationssimulation.SimulateMsgBlockSenders(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
