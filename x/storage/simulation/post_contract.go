package simulation

import (
	"math/rand"
	"strconv"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/jackalLabs/canine-chain/x/storage/keeper"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func SimulateMsgPostContract(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		providers := k.GetAllProviders(ctx)
		if len(providers) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgPostContract, "providers are not initiated"), nil, nil
		}

		provider := providers[simtypes.RandIntBetween(r, 0, len(providers))]

		simAccount, found := simtypes.FindAccount(accs, sdk.MustAccAddressFromBech32(provider.Creator))

		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgPostContract, "provider address is unknown"), nil, nil
		}

		msg := &types.MsgPostContract{
			Creator: provider.Creator,
		}

		users := k.GetAllStoragePaymentInfo(ctx)
		if len(users) == 0 {
			return simtypes.NoOpMsg(
				types.ModuleName, types.TypeMsgPostContract, "storage payment infos are not initiated"), nil, nil
		}
		msg.Signee = users[simtypes.RandIntBetween(r, 0, len(users))].Address

		msg.Filesize = strconv.Itoa(len(fileData))
		fid, err := bech32.ConvertAndEncode(
			"jklf", []byte(simtypes.RandStringOfLength(r, 20)))
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgPostContract, "failed to generate fid"), nil, err
		}

		msg.Fid = fid
		msg.Merkle = GetMerkleRoot()

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
