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

func SimulateMsgTransfer(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, _ string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// checking if enough accounts exist
		if len(accs) < 2 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgList, "Need more than 2 accounts to transfer names"), nil, nil
		}

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
		msg := &types.MsgTransfer{
			Creator: simAccount.Address.String(),
		}

		// selecting a random name to transfer
		nameI := simtypes.RandIntBetween(r, 0, len(names))
		tName := names[nameI]

		// checking if the name is transferrable
		if ctx.BlockHeight() > tName.Expires {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Expired domain"), nil, nil
		}
		if tName.Locked > ctx.BlockHeight() {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Can't transfer a free name"), nil, nil
		}

		// generating the fees
		price := sdk.NewInt(0) // transferring is free?
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

		// generating randomly generated account
		receiverAcc, _ := simtypes.RandomAcc(r, accs)
		// ensuring the sender and receiver are different
		for receiverAcc.Address.String() == simAccount.Address.String() {
			receiverAcc, _ = simtypes.RandomAcc(r, accs)
		}

		// transferring to the randomly generated simulation account
		msg.Name = tName.Name + "." + tName.Tld
		msg.Receiver = simAccount.Address.String()

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
