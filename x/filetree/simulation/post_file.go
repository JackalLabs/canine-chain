package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/jackalLabs/canine-chain/x/filetree/keeper"
	"github.com/jackalLabs/canine-chain/x/filetree/types"
)

func SimulateMsgPostFile(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgPostFile{
			Creator: simAccount.Address.String(),
		}

		files, err := types.CreateRootFolder(simAccount.Address.String())
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate root folder"), nil, err
		}

		_, found := k.GetFiles(ctx, files.Address, files.Owner)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to find root folder"), nil, nil
		}

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "PostFile simulation not implemented"), nil, nil
	}
}
