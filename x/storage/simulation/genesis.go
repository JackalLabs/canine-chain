package simulation

import (
	"math/rand"
	"strconv"

	"github.com/cosmos/cosmos-sdk/types/module"
	sdksim "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func RandomizedGenState(simState *module.SimulationState) {
	var depAcc string

	simState.AppParams.GetOrGenerate(
		simState.Cdc, string(types.KeyDepositAccount), &depAcc, simState.Rand,
		func(r *rand.Rand) {
			depAcc = genDepositAccount(r).Address.String()
		},
	)

	var providers []types.Providers

	// Generate random amount of provider accounts
	// TODO: Have this value exposed
	randProviderCount := sdksim.RandIntBetween(simState.Rand, 0, 100)

	provAccs := sdksim.RandomAccounts(simState.Rand, randProviderCount)
	simState.Accounts = append(simState.Accounts, provAccs...)

	for i := 0; i < randProviderCount; i++ {
		provider := types.Providers{
			Address: provAccs[i].Address.String(),
			Ip:      RandIPv4(simState.Rand),
			Totalspace: strconv.Itoa(
				// Between 1Tb and 1Pb
				sdksim.RandIntBetween(simState.Rand, 1_000_000_000_000, 1_000_000_000_000_000)),
			Creator:         provAccs[i].Address.String(),
			BurnedContracts: "0",
			KeybaseIdentity: sdksim.RandStringOfLength(simState.Rand, 10),
		}

		providers = append(providers, provider)
	}

	storageGen := types.DefaultGenesis()
	storageGen.ProvidersList = providers
	storageGen.Params = types.DefaultParams()

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(storageGen)
}
