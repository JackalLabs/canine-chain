package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/jackalLabs/canine-chain/x/dsig/keeper"
	"github.com/jackalLabs/canine-chain/x/dsig/types"
)

func SimulateMsgSignform(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSignform{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the Signform simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Signform simulation not implemented"), nil, nil
	}
}
