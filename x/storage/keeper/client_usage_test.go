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

func createNClientUsage(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ClientUsage {
	items := make([]types.ClientUsage, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetClientUsage(ctx, items[i])
	}
	return items
}

func TestClientUsageGet(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	items := createNClientUsage(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetClientUsage(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestClientUsageRemove(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	items := createNClientUsage(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveClientUsage(ctx,
			item.Address,
		)
		_, found := keeper.GetClientUsage(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestClientUsageGetAll(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	items := createNClientUsage(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllClientUsage(ctx)),
	)
}
