package jklaccounts_test

import (
	"testing"

	keepertest "github.com/jackal-dao/canine/testutil/keeper"
	"github.com/jackal-dao/canine/testutil/nullify"
	"github.com/jackal-dao/canine/x/jklaccounts"
	"github.com/jackal-dao/canine/x/jklaccounts/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		AccountsList: []types.Accounts{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.JklaccountsKeeper(t)
	jklaccounts.InitGenesis(ctx, *k, genesisState)
	got := jklaccounts.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.AccountsList, got.AccountsList)
	// this line is used by starport scaffolding # genesis/test/assert
}
