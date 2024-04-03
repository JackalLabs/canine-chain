package simulation

import (
	"fmt"
	"math/rand"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/jackalLabs/canine-chain/v3/x/rns/types"
)

func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	_ = r
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyDepositAccount),
			func(_ *rand.Rand) string {
				return fmt.Sprintf("\"%s\"", "jkl14c3j672kvw9l5uleh4x9uds2fre5vl7yun4ss8")
			},
		),
	}
}
