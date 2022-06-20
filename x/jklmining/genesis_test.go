package jklmining_test

import (
	"testing"

	keepertest "github.com/jackal-dao/canine/testutil/keeper"
	"github.com/jackal-dao/canine/testutil/nullify"
	"github.com/jackal-dao/canine/x/jklmining"
	"github.com/jackal-dao/canine/x/jklmining/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		SaveRequestsList: []types.SaveRequests{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		MinersList: []types.Miners{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		MinedList: []types.Mined{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		MinedCount: 2,
		MinerClaimsList: []types.MinerClaims{
			{
				Hash: "0",
			},
			{
				Hash: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.JklminingKeeper(t)
	jklmining.InitGenesis(ctx, *k, genesisState)
	got := jklmining.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.SaveRequestsList, got.SaveRequestsList)
	require.ElementsMatch(t, genesisState.MinersList, got.MinersList)
	require.ElementsMatch(t, genesisState.MinedList, got.MinedList)
	require.Equal(t, genesisState.MinedCount, got.MinedCount)
	require.ElementsMatch(t, genesisState.MinerClaimsList, got.MinerClaimsList)
	// this line is used by starport scaffolding # genesis/test/assert
}
