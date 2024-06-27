package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/jackalLabs/canine-chain/v4/x/storage/keeper"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
)

func SimulateMsgSetProviderIP(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, _ string,
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

		ip := RandIPv4Url(r)

		msg := &types.MsgSetProviderIP{
			Creator: provider.Creator,
			Ip:      ip,
		}

		spendable := bk.SpendableCoins(ctx, simAccount.Address)
		fees, err := simtypes.RandomFees(r, ctx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgSetProviderIP, "unable to generate fees"), nil, err
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
