package simulation

import (
	"math/rand"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

// Randomly generate single account
func genDepositAccount(r *rand.Rand) simtypes.Account {
	return simtypes.RandomAccounts(r, 1)[0]
}

func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyDepositAccount),
			func(r *rand.Rand) string {
				return genDepositAccount(r).Address.String()
			},
		),
	}
}
