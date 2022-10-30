package dsig_test

import (
	"dsig/testutil/nullify"
	"testing"

	keepertest "dsig/testutil/keeper"

	"github.com/jackalLabs/canine-chain/x/dsig"
	"github.com/jackalLabs/canine-chain/x/dsig/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		UserUploadsList: []types.UserUploads{
			{
				Fid: "0",
			},
			{
				Fid: "1",
			},
		},
		FormList: []types.Form{
			{
				Ffid: "0",
			},
			{
				Ffid: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DsigKeeper(t)
	dsig.InitGenesis(ctx, *k, genesisState)
	got := dsig.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.UserUploadsList, got.UserUploadsList)
	require.ElementsMatch(t, genesisState.FormList, got.FormList)
	// this line is used by starport scaffolding # genesis/test/assert
}