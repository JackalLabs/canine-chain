package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/jackalLabs/canine-chain/testutil/keeper"
	"github.com/jackalLabs/canine-chain/testutil/nullify"
	"github.com/jackalLabs/canine-chain/x/filetree/keeper"
	"github.com/jackalLabs/canine-chain/x/filetree/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNFiles(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Files {
	items := make([]types.Files, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetFiles(ctx, items[i])
	}
	return items
}

func TestFilesGet(t *testing.T) {
	keeper, ctx := keepertest.FiletreeKeeper(t)
	items := createNFiles(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetFiles(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestFilesRemove(t *testing.T) {
	keeper, ctx := keepertest.FiletreeKeeper(t)
	items := createNFiles(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveFiles(ctx,
			item.Address,
		)
		_, found := keeper.GetFiles(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestFilesGetAll(t *testing.T) {
	keeper, ctx := keepertest.FiletreeKeeper(t)
	items := createNFiles(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllFiles(ctx)),
	)
}
