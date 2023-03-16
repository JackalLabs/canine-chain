package simulation

import (
	"errors"
	"math/rand"
	"strconv"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/jackalLabs/canine-chain/x/notifications/keeper"
	"github.com/jackalLabs/canine-chain/x/notifications/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func SimulateMsgCreateNotifications(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		// choose receiver of the notification
		receivers := k.GetAllNotiCounter(ctx)
		if len(receivers) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateNotifications, "unable to find noti counters"), nil, nil
		}
		i := r.Intn(len(receivers))
		receiver := receivers[i]
		// avoid potential collision
		if receiver.Address == simAccount.Address.String() {
			receiver = receivers[(i+1)/len(receivers)]
		}

		if IsBlocked(simAccount.Address.String(), receiver) {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateNotifications, "chosen account is blocked by receiver"), nil, nil
		}

		msg := &types.MsgCreateNotifications{
			Creator:      simAccount.Address.String(),
			Notification: simtypes.RandStringOfLength(r, 20),
			Address:      receiver.Address,
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgDeleteNotifications(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		notifications := k.GetAllNotifications(ctx)
		if len(notifications) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgDeleteNotifications, "unable to find notifications"), nil, nil
		}

		noti := notifications[r.Intn(len(notifications))]

		msg := &types.MsgDeleteNotifications{
			Creator: noti.Address,
		}

		simAccount, found := simtypes.FindAccount(accs, sdk.MustAccAddressFromBech32(msg.Creator))
		if !found {
			panic(errors.New("notification created with non-existing account"))
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}
