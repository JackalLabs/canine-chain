package filetree_test

import (
	"testing"

	keepertest "github.com/jackal-dao/canine/testutil/keeper"
	"github.com/jackal-dao/canine/testutil/nullify"
	"github.com/jackal-dao/canine/x/filetree"
	"github.com/jackal-dao/canine/x/filetree/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		FilesList: []types.Files{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		PubkeyList: []types.Pubkey{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		Tracker: &types.Tracker{
			TrackingNumber: 92,
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FiletreeKeeper(t)
	filetree.InitGenesis(ctx, *k, genesisState)
	got := filetree.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.FilesList, got.FilesList)
	require.ElementsMatch(t, genesisState.PubkeyList, got.PubkeyList)
	require.Equal(t, genesisState.Tracker, got.Tracker)
	// this line is used by starport scaffolding # genesis/test/assert
}
