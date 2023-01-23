package simulation

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/jackalLabs/canine-chain/x/rns/keeper"
	"github.com/jackalLabs/canine-chain/x/rns/types"
)

func SimulateMsgRegister(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgRegister{
			Creator: simAccount.Address.String(),
		}

		// generating a random name
		// generating a random TLD
		tldIndex := simtypes.RandIntBetween(r, 0, len(types.SupportedTLDs)+1)
		tld := types.SupportedTLDs[tldIndex]

		// generating a random name
		nameLength := simtypes.RandIntBetween(r, 1, 5)
		var name string
		i := 0
		for i < nameLength {
			name += "j"
			i++
		}

		fullDomain := name + "." + tld
		fmt.Print(fullDomain)

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Register simulation not implemented"), nil, nil
	}
}
