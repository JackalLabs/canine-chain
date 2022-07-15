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

func createNMiners(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Miners {
	items := make([]types.Miners, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetMiners(ctx, items[i])
	}
	return items
}

func TestMinersGet(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	items := createNMiners(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetMiners(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestMinersRemove(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	items := createNMiners(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMiners(ctx,
			item.Address,
		)
		_, found := keeper.GetMiners(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestMinersGetAll(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	items := createNMiners(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMiners(ctx)),
	)
}
