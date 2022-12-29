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

func SimulateMsgBuyStorage(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgBuyStorage{
			Creator:    simAccount.Address.String(),
			ForAddress: simAccount.Address.String(),
		}

		jBalance := bk.GetBalance(ctx, simAccount.Address, "ujkl")
		if !jBalance.IsPositive() {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgBuyStorage, "balance is negative"), nil, nil
		}

		cost, err := simtypes.RandPositiveInt(r, jBalance.Amount)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgBuyStorage, "unable to generate positive cost"), nil, err
		}

		jklPrice := k.GetJklPrice(ctx)
		size, duration := RandStoragePlan(r, jklPrice, cost)

		msg.Bytes = strconv.Itoa(int(size))
		msg.Duration = strconv.Itoa(int(duration))

		spendable := bk.SpendableCoins(ctx, simAccount.Address)
		coins, hasNeg := spendable.SafeSub(sdk.NewCoins(sdk.NewCoin("ujkl", cost)))

		var fees sdk.Coins

		if !hasNeg {
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

// Get random size and duration that costs ujklCost
func RandStoragePlan(
	r *rand.Rand, jklPrice sdk.Dec, ujklCost sdk.Int,
) (bytes, duration int64) {
	ujklUnit := sdk.NewDec(1000000)
	pricePerGb := sdk.MustNewDecFromStr("0.008")
	// Calculate GetStorageCost in reverse
	jklcost := sdk.NewDecFromInt(ujklCost).Quo(ujklUnit)
	totalCost := jklcost.Mul(jklPrice)

	/*
		To choose random storage size that won't end up costing more than
		ujkl_cost.
		Randomly choose size between 1 and maximum size with 1 month duration.
		Then get duration based on that size.
	*/

	maxSize := totalCost.Mul(pricePerGb)
	bytes = int64(simtypes.RandIntBetween(r, 1, int(maxSize.TruncateInt64())))
	bytes *= 1_000_000_000

	pricePerMonth := sdk.NewDec(bytes).Mul(pricePerGb)
	duration = totalCost.Quo(pricePerMonth).TruncateInt64()
	return
}
