package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/jackal-dao/canine/x/jklaccounts/keeper"
	"github.com/jackal-dao/canine/x/jklaccounts/types"
)

func SimulateMsgChoosePlan(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgChoosePlan{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the ChoosePlan simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "ChoosePlan simulation not implemented"), nil, nil
	}
}
