package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/jackalLabs/canine-chain/x/amm/types"
	"github.com/jackalLabs/canine-chain/x/amm/keeper"
)

func SimulateMsgExitPool(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgExitPool{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the ExitPool simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "ExitPool simulation not implemented"), nil, nil
	}
}
