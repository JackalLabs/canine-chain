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

func SimulateMsgCancelContract(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgCancelContract{}

		// choose a contract
		contracts := k.GetAllActiveDeals(ctx)
		if len(contracts) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgSignContract, "no contracts exist"), nil, nil
		}
		contract := contracts[simtypes.RandIntBetween(r, 0, len(contracts))]

		simAccount, found := simtypes.FindAccount(
			accs, sdk.MustAccAddressFromBech32(contract.Signer),
		)

		if !found {
			return simtypes.NoOpMsg(
				types.ModuleName, types.TypeMsgSignContract,
				"unable to find contract signee in []simtypes.Account",
			), nil, nil
		}

		_, found = k.GetFidCid(ctx, contract.Fid)
		if !found {
			return simtypes.NoOpMsg(
				types.ModuleName, types.TypeMsgSignContract,
				"unable to find fid to cid",
			), nil, nil
		}

		msg.Creator = contract.Signer
		msg.Cid = contract.Cid

		spendable := bk.SpendableCoins(ctx, simAccount.Address)
		fees, err := simtypes.RandomFees(r, ctx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgPostContract, "failed to generate fee"), nil, err
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
