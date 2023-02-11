package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/jackalLabs/canine-chain/x/rns/keeper"
	"github.com/jackalLabs/canine-chain/x/rns/types"
)

func SimulateMsgBuy(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgBuy{
			Creator: simAccount.Address.String(),
		}

		// choosing a random name listed on the market
		allSale := k.GetAllForsale(ctx)
		if len(allSale) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "No domains for sale"), nil, nil
		}
		bName := allSale[r.Intn(len(allSale))]

		// ensuring the sim accounts isn't the owner
		if bName.Owner == simAccount.Address.String() {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to choose buyer"), nil, nil
		}

		// ensuring the simAccount can buy the domain
		price, err := sdk.ParseCoinNormalized(bName.Price)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Unable to build "), nil, nil
		}
		sdkPrice := price.Amount

		jBalance := bk.GetBalance(ctx, simAccount.Address, "ujkl")
		// It is impossible to specify default bond denom through param.json
		// to naturally fund the accounts with ujkl.
		// The other option is genesis.json but it is not possible to sign transactions
		// due to private and pubkeys are generated independent of addresses
		// resulting pubkey does not match signer address error.
		if jBalance.Amount.LTE(sdkPrice) {
			c := sdk.NewCoin("ujkl", sdkPrice)

			err := bk.MintCoins(ctx, types.ModuleName, sdk.NewCoins(c))
			if err != nil {
				return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unabled to fund account"), nil, err
			}

			err = bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, simAccount.Address, sdk.NewCoins(c))
			if err != nil {
				return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unabled to fund account"), nil, err
			}
		}

		// generating the fees
		spendable := bk.SpendableCoins(ctx, simAccount.Address)
		coins, hasNeg := spendable.SafeSub(sdk.NewCoins(sdk.NewCoin("ujkl", sdkPrice)))

		var fees sdk.Coins
		if !hasNeg {
			var err error
			fees, err = simtypes.RandomFees(r, ctx, coins)
			if err != nil {
				return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate fees"), nil, err
			}
		}

		// filling the message details
		msg.Name = bName.Name

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
