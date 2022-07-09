package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/jackal-dao/canine/testutil/keeper"
	"github.com/jackal-dao/canine/testutil/nullify"
	"github.com/jackal-dao/canine/x/rns/keeper"
	"github.com/jackal-dao/canine/x/rns/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNNames(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Names {
	items := make([]types.Names, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetNames(ctx, items[i])
	}
	return items
}

func TestNamesGet(t *testing.T) {
	keeper, ctx := keepertest.RnsKeeper(t)
	items := createNNames(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNames(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestNamesRemove(t *testing.T) {
	keeper, ctx := keepertest.RnsKeeper(t)
	items := createNNames(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNames(ctx,
			item.Index,
		)
		_, found := keeper.GetNames(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestNamesGetAll(t *testing.T) {
	keeper, ctx := keepertest.RnsKeeper(t)
	items := createNNames(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNames(ctx)),
	)
}
