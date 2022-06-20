package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/jackal-dao/canine/x/jklmining/keeper"
	"github.com/jackal-dao/canine/x/jklmining/types"
)

func SimulateMsgClaimSave(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgClaimSave{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the ClaimSave simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "ClaimSave simulation not implemented"), nil, nil
	}
}
