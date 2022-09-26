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

func createNLProviderRecord(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.LProviderRecord {
	items := make([]types.LProviderRecord, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetLProviderRecord(ctx, items[i])
	}
	return items
}

func TestLProviderRecordGet(t *testing.T) {
	keeper, ctx := keepertest.LpKeeper(t)
	items := createNLProviderRecord(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetLProviderRecord(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestLProviderRecordRemove(t *testing.T) {
	keeper, ctx := keepertest.LpKeeper(t)
	items := createNLProviderRecord(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLProviderRecord(ctx,
			item.Index,
		)
		_, found := keeper.GetLProviderRecord(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestLProviderRecordGetAll(t *testing.T) {
	keeper, ctx := keepertest.LpKeeper(t)
	items := createNLProviderRecord(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllLProviderRecord(ctx)),
	)
}
