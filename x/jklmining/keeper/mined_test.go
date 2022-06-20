package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/jackal-dao/canine/testutil/keeper"
	"github.com/jackal-dao/canine/testutil/nullify"
	"github.com/jackal-dao/canine/x/jklmining/keeper"
	"github.com/jackal-dao/canine/x/jklmining/types"
	"github.com/stretchr/testify/require"
)

func createNMined(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Mined {
	items := make([]types.Mined, n)
	for i := range items {
		items[i].Id = keeper.AppendMined(ctx, items[i])
	}
	return items
}

func TestMinedGet(t *testing.T) {
	keeper, ctx := keepertest.JklminingKeeper(t)
	items := createNMined(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetMined(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestMinedRemove(t *testing.T) {
	keeper, ctx := keepertest.JklminingKeeper(t)
	items := createNMined(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMined(ctx, item.Id)
		_, found := keeper.GetMined(ctx, item.Id)
		require.False(t, found)
	}
}

func TestMinedGetAll(t *testing.T) {
	keeper, ctx := keepertest.JklminingKeeper(t)
	items := createNMined(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMined(ctx)),
	)
}

func TestMinedCount(t *testing.T) {
	keeper, ctx := keepertest.JklminingKeeper(t)
	items := createNMined(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetMinedCount(ctx))
}
