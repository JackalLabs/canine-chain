package simulation

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/jackalLabs/canine-chain/x/rns/keeper"
	"github.com/jackalLabs/canine-chain/x/rns/types"
)

func SimulateMsgAcceptBid(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// 1. find a random name
		// 2. check if it has domains forsale
		// 3. scan bids, see if there exists an active bid
		// 4. if there is an active bid, find the owner account

		// getting a random domain name that is on sale
		allNames := k.GetAllForsale(ctx)
		if len(allNames) < 1 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgAcceptBid, "No domain names listed"), nil, nil
		}
		rNameI := simtypes.RandIntBetween(r, 0, len(allNames))
		rName := allNames[rNameI]

		// scanning bids
		allBids := k.GetAllBids(ctx)
		var bidderName string
		for _, bids := range allBids {
			// assuming all addresses are 42 characters long
			auctionedName := bids.Index[42:]
			if auctionedName == rName.Name {
				bidderName = bids.Index[:42]
				break
			}
		}
		if bidderName == "" {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgAcceptBid, "No active bids available"), nil, nil
		}

		// finding the owner account
		var simAccount simtypes.Account
		for _, acc := range accs {
			if acc.Address.String() == rName.Owner {
				simAccount = acc
				break
			}
		}

		if len(simAccount.Address) < 1 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgAcceptBid, "Could not find owner account"), nil, nil
		}

		// packaging the request
		msg := &types.MsgAcceptBid{
			Creator: simAccount.Address.String(),
			Name:    rName.Name,
			From:    bidderName,
		}

		n, tld, _ := keeper.GetNameAndTLD(rName.Name)
		whois, isFound := k.GetNames(ctx, n, tld)
		simaccstr := simAccount.Address.String()
		fmt.Print(whois, isFound, simaccstr)

		// calculating the fees
		price := sdk.NewInt(0) // accepting bids is free?
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
