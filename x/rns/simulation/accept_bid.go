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

func GetBidsFor(k keeper.Keeper, ctx sdk.Context, name string) (bids []types.Bids) {
	allBids := k.GetAllBids(ctx)
	if len(allBids) == 0 {
		return nil
	}

	for _, bid := range allBids {
		if bid.Name == name {
			bids = append(bids, bid)
		}
	}
	return bids
}

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
		if len(allNames) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgAcceptBid, "No domain names listed"), nil, nil
		}
		rName := allNames[r.Intn(len(allNames))]

		n, tld, err := keeper.GetNameAndTLD(rName.Name)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgAcceptBid, "unable to get name and tld"), nil, err
		}

		name, found := k.GetNames(ctx, n, tld)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgAcceptBid, "unable to get name and tld"), nil, err
		}
		if rName.Owner != name.Value {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgAcceptBid, "bid is no longer valid"), nil, nil
		}

		// scanning bids
		bids := GetBidsFor(k, ctx, rName.Name)
		if len(bids) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgAcceptBid, "unalbe to find bids"), nil, nil
		}
		bid := bids[r.Intn(len(bids))]

		// finding the owner account
		acc := sdk.MustAccAddressFromBech32(rName.Owner)
		simAccount, found := simtypes.FindAccount(accs, acc)

		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgAcceptBid, "owner account not found"),
				nil,
				fmt.Errorf("rns registered with non-existing account")
		}

		// packaging the request
		msg := &types.MsgAcceptBid{
			Creator: rName.Owner,
			Name:    rName.Name,
			From:    bid.Bidder,
		}

		spendable := bk.SpendableCoins(ctx, simAccount.Address)
		fees, err := simtypes.RandomFees(r, ctx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate fees"), nil, err
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
