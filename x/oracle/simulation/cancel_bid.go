package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/jackalLabs/canine-chain/v5/x/rns/keeper"
	"github.com/jackalLabs/canine-chain/v5/x/rns/types"
)

func SimulateMsgCancelBid(
	_ types.AccountKeeper,
	_ types.BankKeeper,
	_ keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, _ *baseapp.BaseApp, _ sdk.Context, accs []simtypes.Account, _ string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCancelBid{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CancelBid simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CancelBid simulation not implemented"), nil, nil
	}
}
