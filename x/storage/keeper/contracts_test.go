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

func createNContracts(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Contracts {
	items := make([]types.Contracts, n)
	for i := range items {
		items[i].Cid = strconv.Itoa(i)

		keeper.SetContracts(ctx, items[i])
	}
	return items
}

func TestContractsGet(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	items := createNContracts(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetContracts(ctx,
			item.Cid,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestContractsRemove(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	items := createNContracts(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveContracts(ctx,
			item.Cid,
		)
		_, found := keeper.GetContracts(ctx,
			item.Cid,
		)
		require.False(t, found)
	}
}

func TestContractsGetAll(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	items := createNContracts(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllContracts(ctx)),
	)
}
