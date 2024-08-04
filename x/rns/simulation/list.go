package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/jackalLabs/canine-chain/v4/x/rns/keeper"
	"github.com/jackalLabs/canine-chain/v4/x/rns/types"
)

func SimulateMsgList(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, _ string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
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
			nReq := &types.QueryListOwnedNames{
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
			}
			simAccount, _ = simtypes.RandomAcc(r, accs)

		}

		// initializing the message
		msg := &types.MsgList{
			Creator: simAccount.Address.String(),
		}

		// choosing a random name that isn't already listed
		unListed := make([]types.Names, 0)
		for _, name := range names {
			if _, found := k.GetForsale(ctx, name.Name+"."+name.Tld); !found && name.Name != "" {
				unListed = append(unListed, name)
			}
		}

		if len(unListed) < 1 {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "All owned domains are listed"), nil, nil
		}
		nameI := simtypes.RandIntBetween(r, 0, len(unListed))
		tName := unListed[nameI]

		// checking if the name is listable
		if ctx.BlockHeight() > tName.Expires {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Expired domain"), nil, nil
		}
		if tName.Locked > ctx.BlockHeight() {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Can't list a free name"), nil, nil
		}

		// generating the fees
		price := sdk.NewInt(0) // listing is free?
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

		msg.Name = tName.Name + "." + tName.Tld
		msg.Price = sdk.NewInt64Coin("ujkl", int64(simtypes.RandIntBetween(r, 0, 10000000)))

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
		OpMsg, _, err := simulation.GenAndDeliverTx(txCtx, fees)

		// creating a future op
		fOp := simtypes.FutureOperation{
			BlockHeight: int(ctx.BlockHeight()) + 15,
			Op:          SimulateMsgBid(ak, bk, k),
		}
		fOps := []simtypes.FutureOperation{fOp}

		return OpMsg, fOps, err
	}
}
