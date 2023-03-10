package simulation

import (
	"encoding/json"
	"errors"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/jackalLabs/canine-chain/x/notifications/keeper"
	"github.com/jackalLabs/canine-chain/x/notifications/types"
)

func SimulateMsgBlockSenders(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		counters := k.GetAllNotiCounter(ctx)
		if len(counters) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgBlockSenders, "unable to find account"), nil, nil
		}

		counter := counters[r.Intn(len(counters))]
		simAccount, found := simtypes.FindAccount(accs, sdk.MustAccAddressFromBech32(counter.Address))
		if !found {
			panic(errors.New("noti counter created with non-existing account"))
		}

		msg := &types.MsgBlockSenders{
			Creator: simAccount.Address.String(),
		}

		blockAddr, _ := simtypes.RandomAcc(r, accs)
		if IsBlocked(blockAddr.Address.String(), counter) {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgBlockSenders, "account already blocked"), nil, nil
		}

		senderIds := make([]string, 1)
		senderIds[0] = blockAddr.Address.String()

		j, err := json.Marshal(senderIds)
		if err != nil {
			panic(err)
		}

		msg.SenderIds = string(j)

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
