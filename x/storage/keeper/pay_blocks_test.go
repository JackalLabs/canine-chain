package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/jackal-dao/canine/testutil/keeper"
	"github.com/jackal-dao/canine/testutil/nullify"
	"github.com/jackal-dao/canine/x/storage/keeper"
	"github.com/jackal-dao/canine/x/storage/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNPayBlocks(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PayBlocks {
	items := make([]types.PayBlocks, n)
	for i := range items {
		items[i].Blockid = strconv.Itoa(i)

		keeper.SetPayBlocks(ctx, items[i])
	}
	return items
}

func TestPayBlocksGet(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	items := createNPayBlocks(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPayBlocks(ctx,
			item.Blockid,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPayBlocksRemove(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	items := createNPayBlocks(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePayBlocks(ctx,
			item.Blockid,
		)
		_, found := keeper.GetPayBlocks(ctx,
			item.Blockid,
		)
		require.False(t, found)
	}
}

func TestPayBlocksGetAll(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	items := createNPayBlocks(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPayBlocks(ctx)),
	)
}
