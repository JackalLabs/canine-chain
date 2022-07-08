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

func createNBids(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Bids {
	items := make([]types.Bids, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetBids(ctx, items[i])
	}
	return items
}

func TestBidsGet(t *testing.T) {
	keeper, ctx := keepertest.TelescopeKeeper(t)
	items := createNBids(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetBids(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestBidsRemove(t *testing.T) {
	keeper, ctx := keepertest.TelescopeKeeper(t)
	items := createNBids(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveBids(ctx,
			item.Index,
		)
		_, found := keeper.GetBids(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestBidsGetAll(t *testing.T) {
	keeper, ctx := keepertest.TelescopeKeeper(t)
	items := createNBids(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllBids(ctx)),
	)
}
