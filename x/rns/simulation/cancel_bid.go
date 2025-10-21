package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/jackalLabs/canine-chain/v5/x/rns/keeper"
	"github.com/jackalLabs/canine-chain/v5/x/rns/types"
)

func SimulateMsgCancelBid(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, _ string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// choosing a random account with a bid open
		nreq := &types.QueryAllBids{}
		wctx := sdk.WrapSDKContext(ctx)
		allBidsResp, err := k.AllBids(wctx, nreq)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCancelBid, "Unable to collect bids"), nil, err
		}
		allBids := allBidsResp.GetBids()
		if len(allBids) < 1 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCancelBid, "No bids to collect"), nil, nil
		}
		randomBidI := simtypes.RandIntBetween(r, 0, len(allBids))
		rBid := allBids[randomBidI]

		bidAddress, err := sdk.AccAddressFromBech32(rBid.Bidder)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCancelBid, "Unable to convert username"), nil, err
		}

		simAccount, ok := simtypes.FindAccount(accs, bidAddress)
		if !ok {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCancelBid, "Unable to find bidder"), nil, err
		}

		// populating the message
		msg := &types.MsgCancelBid{
			Creator: simAccount.Address.String(),
			Name:    rBid.Name,
		}

		// generating the fees
		price := sdk.NewInt(0) // cancelling bid is free?
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

		// configuring the tx
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
