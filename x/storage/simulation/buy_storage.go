package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/jackalLabs/canine-chain/v5/x/storage/keeper"
	"github.com/jackalLabs/canine-chain/v5/x/storage/types"
)

func SimulateMsgBuyStorage(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, _ string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgBuyStorage{
			Creator:      simAccount.Address.String(),
			ForAddress:   simAccount.Address.String(),
			PaymentDenom: "ujkl",
		}

		_, found := k.GetStoragePaymentInfo(ctx, simAccount.Address.String())
		if found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgBuyStorage, "user already paid for storage, skipping"), nil, nil
		}

		size := simtypes.RandIntBetween(r, 1_000_000_000, 10_000_000_000)

		var t int64 = 30

		hours := sdk.NewDec(t * 24)
		cost := k.GetStorageCost(ctx, int64(size), hours.TruncateInt64())

		msg.Bytes = int64(size)
		msg.DurationDays = t

		jBalance := bk.GetBalance(ctx, simAccount.Address, "ujkl")
		// It is impossible to specify default bond denom through param.json
		// to naturally fund the accounts with ujkl.
		// The other option is genesis.json but it is not possible to sign transactions
		// due to private and pubkeys are generated independent of addresses
		// resulting pubkey does not match signer address error.
		if jBalance.Amount.LTE(cost) {
			c := sdk.NewCoin("ujkl", cost.MulRaw(2))

			err := bk.MintCoins(ctx, types.ModuleName, sdk.NewCoins(c))
			if err != nil {
				return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgBuyStorage, "unabled to fund account"), nil, err
			}

			err = bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, simAccount.Address, sdk.NewCoins(c))
			if err != nil {
				return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgBuyStorage, "unabled to fund account"), nil, err
			}
		}

		spendable := bk.SpendableCoins(ctx, simAccount.Address)
		coins, hasNeg := spendable.SafeSub(sdk.NewCoins(sdk.NewCoin("ujkl", cost)))

		var fees sdk.Coins

		if !hasNeg {
			var err error
			fees, err = simtypes.RandomFees(r, ctx, coins)
			if err != nil {
				return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgBuyStorage, "unable to generate fees"), nil, err
			}
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
