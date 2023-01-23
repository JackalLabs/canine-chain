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

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func StringWithCharset(r *rand.Rand, length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[r.Intn(len(charset))]
	}
	return string(b)
}

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
		tldIndex := simtypes.RandIntBetween(r, 0, len(types.SupportedTLDs))
		tld := types.SupportedTLDs[tldIndex]

		// generating a random name
		nameLength := simtypes.RandIntBetween(r, 1, 5)
		name := StringWithCharset(r, nameLength, charset)

		fullDomain := name + "." + tld
		fmt.Print(fullDomain)

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Register simulation not implemented"), nil, nil
	}
}
