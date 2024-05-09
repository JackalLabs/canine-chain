package simulation

import (
	"math/rand"
	"strconv"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/jackalLabs/canine-chain/v3/x/storage/keeper"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

func SimulateMsgInitProvider(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, _ string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		_, found := k.GetProviders(ctx, simAccount.Address.String())
		if found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgInitProvider, "provider already exists"), nil, nil
		}

		coins := sdk.NewCoins(sdk.NewInt64Coin("ujkl", 10_000_000_000_000_000))
		err := bk.MintCoins(ctx, types.ModuleName, coins)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgInitProvider, "failed to mint collateral"), nil, err
		}
		err = bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, simAccount.Address, coins)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgInitProvider, "failed to generate collateral"), nil, err
		}

		msg := &types.MsgInitProvider{
			Creator:    simAccount.Address.String(),
			Ip:         RandIPv4Url(r),
			TotalSpace: strconv.FormatInt(int64(simtypes.RandIntBetween(r, 1_000_000_000, 1_000_000_000_000_000)), 10),
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
