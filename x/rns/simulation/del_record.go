package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/jackalLabs/canine-chain/v3/x/rns/keeper"
	"github.com/jackalLabs/canine-chain/v3/x/rns/types"
)

func SimulateMsgDelRecord(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// 1. Find a registered name with subdomains
		// 2. Find the associated account
		// choosing a random account WITH registered domains
		var simAccount simtypes.Account
		var name types.Names
		var nameWithSub string
		// checking if any names are registered
		exists := k.CheckExistence(ctx)
		if !exists {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgDelRecord, "No domains registered yet"), nil, nil
		}
		// scanning all names
		for _, n := range k.GetAllNames(ctx) {
			if n.Subdomains != nil {
				name = n
				nameWithSub = n.Subdomains[0].Name + "." + name.Name
			}
		}
		if nameWithSub == "" {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgDelRecord, "No registered names have subdomains"), nil, nil
		}

		// finding the owner
		for _, o := range accs {
			if o.Address.String() == name.Value {
				simAccount = o
			}
		}
		if simAccount.Address.String() == "" {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgDelRecord, "Could not find owner"), nil, nil
		}

		// initializing the message
		msg := &types.MsgDelRecord{
			Creator: simAccount.Address.String(),
		}

		// generating the fees
		price := sdk.NewInt(0)
		spendable := bk.SpendableCoins(ctx, simAccount.Address)
		coins, hasNeg := spendable.SafeSub(sdk.NewCoins(sdk.NewCoin("ujkl", price)))

		var fees sdk.Coins

		if !hasNeg {
			var err error
			fees, err = simtypes.RandomFees(r, ctx, coins)
			if err != nil {
				return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate fees"), nil, err
			}
		}

		// building the message
		msg.Name = nameWithSub + "." + name.Tld

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
