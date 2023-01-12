package simulation

import (
	"math/rand"
	"strconv"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/jackalLabs/canine-chain/x/storage/keeper"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func SimulateMsgSetProviderTotalspace(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		providers := k.GetAllProviders(ctx)
		if len(providers) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgSetProviderIP, "unable to find provider"), nil, nil
		}

		provider := providers[rand.Intn(len(providers))]
		simAccount, found := simtypes.FindAccount(accs, sdk.MustAccAddressFromBech32(provider.Creator))
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgSetProviderIP, "unable to find provider"), nil, nil
		}

		msg := &types.MsgSetProviderTotalspace{
			Creator: provider.Creator,
			Space: strconv.Itoa(simtypes.RandIntBetween(r, 1_000_000_000, 1_000_000_000_000_000)),
		}

		spendable := bk.SpendableCoins(ctx, simAccount.Address)
		fees, err := simtypes.RandomFees(r, ctx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgSignContract, "unable to generate fees"), nil, err
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
