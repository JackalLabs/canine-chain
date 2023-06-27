package simulation

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/jackalLabs/canine-chain/v3/x/rns/types"
)

func RandomizedGenState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}

	rnsGenesis := types.DefaultGenesis()
	p := types.DefaultParams()
	p.DepositAccount = "jkl1arsaayyj5tash86mwqudmcs2fd5jt5zgc3sexc"
	rnsGenesis.Params = p

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(rnsGenesis)
}
