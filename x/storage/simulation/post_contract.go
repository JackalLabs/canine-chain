package simulation

import (
	"strconv"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/jackalLabs/canine-chain/x/storage/keeper"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func SimulateMsgPostContract(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgPostContract{}

		// Choose random provider
		providers := k.GetAllProviders(ctx)
		if len(providers) <= 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgPostContract, "providers are not initiated"), nil, nil
		}
		provider := providers[simtypes.RandIntBetween(r, 0, len(providers))]
		msg.Creator = provider.Address

		// Choose random signee
		users := k.GetAllStoragePaymentInfo(ctx)
		if len(users) <= 0 {
			return simtypes.NoOpMsg(
				types.ModuleName, types.TypeMsgPostContract, "storage payment infos are not initiated"), nil, nil
		}
		msg.Signee = users[simtypes.RandIntBetween(r, 0, len(users))].Address

		msg.Filesize = strconv.Itoa(simtypes.RandIntBetween(r, 1, 100_000_000_000_000))
		fid, err := bech32.ConvertAndEncode(
			"jklf", []byte(simtypes.RandStringOfLength(r, 20)))

		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgPostContract, "failed to generate fid"), nil, err
		}

		msg.Fid = fid
		msg.Merkle = simtypes.RandStringOfLength(r, 20)

		spendable := bk.SpendableCoins(ctx, simAccount.Address)
		fees, err := simtypes.RandomFees(r, ctx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate fees"), nil, err
		}

		txCtx := simulation.OperationInput{
			R: r,
			App: app,
			TxGen: simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc: nil,
			Msg: msg,
			MsgType: msg.Type(),
			Context: ctx,
			SimAccount: simAccount,
			AccountKeeper: ak,
			ModuleName: types.ModuleName,
		}

		return simulation.GenAndDeliverTx(txCtx, fees)
	}
}
