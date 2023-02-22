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

func SimulateMsgDelist(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		forsales := k.GetAllForsale(ctx)
		if len(forsales) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgDelist, "unable to find names for sale"), nil, nil
		}
		forsale := forsales[r.Intn(len(forsales))]

		name, tld, err := keeper.GetNameAndTLD(forsale.Name)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgDelist, "unable to get name and tld"), nil, err
		}

		rns, _ := k.GetNames(ctx, name, tld)
		if rns.Value != forsale.Owner {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgDelist, "name is expired"), nil, nil
		}

		msg := &types.MsgDelist{
			Creator: forsale.Owner,
			Name:    forsale.Name,
		}

		simAccount, found := simtypes.FindAccount(accs, sdk.MustAccAddressFromBech32(msg.Creator))
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgDelist, "unable to find names for sale"),
				nil,
				fmt.Errorf("account not found")
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
