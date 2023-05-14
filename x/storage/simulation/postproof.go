package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/jackalLabs/canine-chain/x/storage/keeper"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func SimulateMsgPostproof(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgPostproof{}

		deals := k.GetAllActiveDeals(ctx)
		if len(deals) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to find active deals"), nil, nil
		}

		deal := deals[r.Intn(len(deals))]

		if deal.ProofVerified {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "proof already verified, skipping"), nil, nil
		}

		provider, found := k.GetProviders(ctx, deal.Provider)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to find provider for an active deal"), nil, nil
		}

		simAccount, found := simtypes.FindAccount(
			accs, sdk.MustAccAddressFromBech32(provider.Address),
		)

		if !found {
			return simtypes.NoOpMsg(
				types.ModuleName, msg.Type(), "unable to find provider account from []simtypes.Account"), nil, nil
		}

		msg.Item, msg.Hashlist = GetMerkleProof()
		msg.Creator = simAccount.Address.String()
		msg.Cid = deal.Cid

		spendable := bk.SpendableCoins(ctx, simAccount.Address)
		fees, err := simtypes.RandomFees(r, ctx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgSignContract, "unable to generate fees"), nil, err
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
