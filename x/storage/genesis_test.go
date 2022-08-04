package storage_test

import (
	"testing"

	keepertest "github.com/jackal-dao/canine/testutil/keeper"
	"github.com/jackal-dao/canine/testutil/nullify"
	"github.com/jackal-dao/canine/x/storage"
	"github.com/jackal-dao/canine/x/storage/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ContractsList: []types.Contracts{
			{
				Cid: "0",
			},
			{
				Cid: "1",
			},
		},
		ProofsList: []types.Proofs{
			{
				Cid: "0",
			},
			{
				Cid: "1",
			},
		},
		ActiveDealsList: []types.ActiveDeals{
			{
				Cid: "0",
			},
			{
				Cid: "1",
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
		PayBlocksList: []types.PayBlocks{
			{
				Blockid: "0",
			},
			{
				Blockid: "1",
			},
		},
		ClientUsageList: []types.ClientUsage{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.StorageKeeper(t)
	storage.InitGenesis(ctx, *k, genesisState)
	got := storage.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ContractsList, got.ContractsList)
	require.ElementsMatch(t, genesisState.ProofsList, got.ProofsList)
	require.ElementsMatch(t, genesisState.ActiveDealsList, got.ActiveDealsList)
	require.ElementsMatch(t, genesisState.MinersList, got.MinersList)
	require.ElementsMatch(t, genesisState.PayBlocksList, got.PayBlocksList)
	require.ElementsMatch(t, genesisState.ClientUsageList, got.ClientUsageList)
	// this line is used by starport scaffolding # genesis/test/assert
}
