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

func SimulateMsgDelist(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// Selecting accounts with listed domains
		// choosing a random account WITH registered domains
		var simAccount simtypes.Account
		var names []types.Names
		simAccount, _ = simtypes.RandomAcc(r, accs)
		// checking if any names are registered
		exists := k.CheckExistence(ctx)
		if !exists {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgList, "No domains registered yet"), nil, nil
		}
		for {
			// finding all registered domain names
			wctx := sdk.WrapSDKContext(ctx)
			nReq := &types.QueryListOwnedNamesRequest{
				Address: simAccount.Address.String(),
			}
			// requesting the domain names
			regNames, err := k.ListOwnedNames(wctx, nReq)
			if err != nil {
				return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgList, "Couldn't request names"), nil, nil
			}
			names = regNames.GetNames()
			if names != nil {
				break
			} else {
				simAccount, _ = simtypes.RandomAcc(r, accs)
			}
		}

		msg := &types.MsgDelist{
			Creator: simAccount.Address.String(),
		}
		// finding the first name that the random sim account has listed
		var deList types.Names
		for _, n := range names {
			if _, found := k.GetForsale(ctx, n.Name+"."+n.Tld); found && n.Name != "" {
				deList = n
			}
		}
		if deList.Name == "" {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgList, "No viable registered domains found"), nil, nil
		}
		// delisting the chosen name
		msg.Name = deList.Name + "." + deList.Tld

		// compiling the fees and the message
		// generating the fees
		price := sdk.NewInt(0) // delisting is free?
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
