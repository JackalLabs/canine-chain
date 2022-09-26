package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/jackal-dao/canine/testutil/keeper"
	"github.com/jackal-dao/canine/testutil/nullify"
	"github.com/jackal-dao/canine/x/lp/keeper"
	"github.com/jackal-dao/canine/x/lp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNLPool(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.LPool {
	items := make([]types.LPool, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetLPool(ctx, items[i])
	}
	return items
}

func TestLPoolGet(t *testing.T) {
	keeper, ctx := keepertest.LpKeeper(t)
	items := createNLPool(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetLPool(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestLPoolRemove(t *testing.T) {
	keeper, ctx := keepertest.LpKeeper(t)
	items := createNLPool(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLPool(ctx,
			item.Index,
		)
		_, found := keeper.GetLPool(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestLPoolGetAll(t *testing.T) {
	keeper, ctx := keepertest.LpKeeper(t)
	items := createNLPool(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllLPool(ctx)),
	)
}
