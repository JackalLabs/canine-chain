package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/jackal-dao/canine/x/lp/keeper"
	"github.com/jackal-dao/canine/x/lp/types"
)

func SimulateMsgWithdrawLPool(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgWithdrawLPool{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the WithdrawLPool simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "WithdrawLPool simulation not implemented"), nil, nil
	}
}
