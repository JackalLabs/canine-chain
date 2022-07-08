package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/themarstonconnell/telescope/testutil/keeper"
	"github.com/themarstonconnell/telescope/testutil/nullify"
	"github.com/themarstonconnell/telescope/x/telescope/keeper"
	"github.com/themarstonconnell/telescope/x/telescope/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNForsale(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Forsale {
	items := make([]types.Forsale, n)
	for i := range items {
		items[i].Name = strconv.Itoa(i)

		keeper.SetForsale(ctx, items[i])
	}
	return items
}

func TestForsaleGet(t *testing.T) {
	keeper, ctx := keepertest.TelescopeKeeper(t)
	items := createNForsale(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetForsale(ctx,
			item.Name,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestForsaleRemove(t *testing.T) {
	keeper, ctx := keepertest.TelescopeKeeper(t)
	items := createNForsale(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveForsale(ctx,
			item.Name,
		)
		_, found := keeper.GetForsale(ctx,
			item.Name,
		)
		require.False(t, found)
	}
}

func TestForsaleGetAll(t *testing.T) {
	keeper, ctx := keepertest.TelescopeKeeper(t)
	items := createNForsale(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllForsale(ctx)),
	)
}
