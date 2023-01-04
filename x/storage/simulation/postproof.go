package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
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
		if len(deals) <= 0 {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to find active deals"), nil, nil
		}

		deal := deals[r.Intn(len(deals))]

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

		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}
