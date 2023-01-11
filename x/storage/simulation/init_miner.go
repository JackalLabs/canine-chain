package simulation

import (
	"math/rand"
	"strconv"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/jackalLabs/canine-chain/x/storage/keeper"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func SimulateMsgInitProvider(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgInitProvider{
			Creator:    simAccount.Address.String(),
			Ip:         RandIPv4Url(r),
			Totalspace: strconv.Itoa(simtypes.RandIntBetween(r, 1_000_000_000, 1_000_000_000_000_000)),
			Keybase:    simtypes.RandStringOfLength(r, 10),
		}

		spendable := bk.SpendableCoins(ctx, simAccount.Address)
		fees, err := simtypes.RandomFees(r, ctx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgInitProvider, "unable to generate fees"), nil, err
		}

		txCtx := simulation.OperationInput{
			R:             r,
			App:           app,
			TxGen:         simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:           nil,
			Msg:           msg,
			MsgType:       msg.Type(),
			Context:       ctx,
			SimAccount:    simAccount,
			AccountKeeper: ak,
			ModuleName:    types.ModuleName,
		}

		return simulation.GenAndDeliverTx(txCtx, fees)
	}
}
