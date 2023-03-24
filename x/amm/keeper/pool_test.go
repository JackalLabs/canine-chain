package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/jackalLabs/canine-chain/testutil/keeper"
	"github.com/jackalLabs/canine-chain/testutil/nullify"
	"github.com/jackalLabs/canine-chain/x/amm/keeper"
	"github.com/jackalLabs/canine-chain/x/amm/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNPool(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Pool {
	items := make([]types.Pool, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetPool(ctx, items[i])
	}
	return items
}

func TestPoolGet(t *testing.T) {
	keeper, ctx := keepertest.LpKeeper(t)
	items := createNPool(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPool(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPoolRemove(t *testing.T) {
	keeper, ctx := keepertest.LpKeeper(t)
	items := createNPool(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePool(ctx,
			item.Index,
		)
		_, found := keeper.GetPool(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestPoolGetAll(t *testing.T) {
	keeper, ctx := keepertest.LpKeeper(t)
	items := createNPool(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPool(ctx)),
	)
}
