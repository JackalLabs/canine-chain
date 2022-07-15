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

func createNActiveDeals(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ActiveDeals {
	items := make([]types.ActiveDeals, n)
	for i := range items {
		items[i].Cid = strconv.Itoa(i)

		keeper.SetActiveDeals(ctx, items[i])
	}
	return items
}

func TestActiveDealsGet(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	items := createNActiveDeals(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetActiveDeals(ctx,
			item.Cid,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestActiveDealsRemove(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	items := createNActiveDeals(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveActiveDeals(ctx,
			item.Cid,
		)
		_, found := keeper.GetActiveDeals(ctx,
			item.Cid,
		)
		require.False(t, found)
	}
}

func TestActiveDealsGetAll(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	items := createNActiveDeals(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllActiveDeals(ctx)),
	)
}
