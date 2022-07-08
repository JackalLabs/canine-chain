package telescope_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/themarstonconnell/telescope/testutil/keeper"
	"github.com/themarstonconnell/telescope/testutil/nullify"
	"github.com/themarstonconnell/telescope/x/telescope"
	"github.com/themarstonconnell/telescope/x/telescope/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		WhoisList: []types.Whois{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		NamesList: []types.Names{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		BidsList: []types.Bids{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		ForsaleList: []types.Forsale{
			{
				Name: "0",
			},
			{
				Name: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TelescopeKeeper(t)
	telescope.InitGenesis(ctx, *k, genesisState)
	got := telescope.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.WhoisList, got.WhoisList)
	require.ElementsMatch(t, genesisState.NamesList, got.NamesList)
	require.ElementsMatch(t, genesisState.BidsList, got.BidsList)
	require.ElementsMatch(t, genesisState.ForsaleList, got.ForsaleList)
	// this line is used by starport scaffolding # genesis/test/assert
}
