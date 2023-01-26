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

func SimulateMsgTransfer(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgTransfer{
			Creator: simAccount.Address.String(),
		}

		if len(accs) < 2 {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Need more than two accounts to transfer names"), nil, nil
		}

		// finding all the registered names
		nReq := &types.QueryListOwnedNamesRequest{
			Address: simAccount.Address.String(),
		}
		wctx := sdk.WrapSDKContext(ctx)
		regNames, err := k.ListOwnedNames(wctx, nReq)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Couldn't fetch names"), nil, err
		}
		// unmarshalling the results
		names := regNames.GetNames()
		if names == nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "No names to transfer"), nil, nil
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
		receiverAcc := &simtypes.Account{}
		res, _ := simtypes.RandomAcc(r, accs)
		receiverAcc = &res
		// ensuring the sender and receiver are different
		for receiverAcc.Address.String() == simAccount.Address.String() {
			res, _ := simtypes.RandomAcc(r, accs)
			receiverAcc = &res
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
