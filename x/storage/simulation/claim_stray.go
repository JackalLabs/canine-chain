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

func SimulateMsgClaimStray(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgClaimStray{}

		strays := k.GetAllStrays(ctx)
		if len(strays) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to find strays"), nil, nil
		}
		stray := strays[r.Intn(len(strays))]

		providers := k.GetAllProviders(ctx)
		if len(providers) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to find providers"), nil, nil
		}
		provider := providers[r.Intn(len(providers))]

		files := k.ListFiles(ctx, stray.Fid)

		for _, file := range files {
			if file == provider.Ip {
				return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to find providers"), nil, nil
			}
		}

		claimer := provider.Address
		if len(provider.AuthClaimers) > 0 {
			claimer = provider.AuthClaimers[r.Intn(len(provider.AuthClaimers))]
		}

		simAccount, found := simtypes.FindAccount(
			accs, sdk.MustAccAddressFromBech32(claimer),
		)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to find claimer account"), nil, nil
		}

		msg.Cid = stray.Cid
		msg.Creator = claimer
		msg.ForAddress = provider.Address

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
