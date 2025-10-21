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
	"github.com/jackalLabs/canine-chain/v5/x/notifications/keeper"
	"github.com/jackalLabs/canine-chain/v5/x/notifications/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func SimulateMsgCreateNotifications(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	_ keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, _ string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		receiver, _ := simtypes.RandomAcc(r, accs)

		msg := &types.MsgCreateNotification{
			Creator:  simAccount.Address.String(),
			To:       receiver.Address.String(),
			Contents: "{}",
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
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, _ string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		notifications := k.GetAllNotifications(ctx)
		if len(notifications) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgDeleteNotification, "unable to find notifications"), nil, nil
		}

		noti := notifications[r.Intn(len(notifications))]

		msg := &types.MsgDeleteNotification{
			Creator: noti.To,
			Time:    noti.Time,
			From:    noti.From,
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
