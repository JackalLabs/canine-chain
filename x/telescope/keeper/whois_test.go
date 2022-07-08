package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/jackal-dao/canine/testutil/keeper"
	"github.com/jackal-dao/canine/testutil/nullify"
	"github.com/jackal-dao/canine/x/telescope/keeper"
	"github.com/jackal-dao/canine/x/telescope/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNWhois(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Whois {
	items := make([]types.Whois, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetWhois(ctx, items[i])
	}
	return items
}

func TestWhoisGet(t *testing.T) {
	keeper, ctx := keepertest.TelescopeKeeper(t)
	items := createNWhois(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetWhois(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestWhoisRemove(t *testing.T) {
	keeper, ctx := keepertest.TelescopeKeeper(t)
	items := createNWhois(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveWhois(ctx,
			item.Index,
		)
		_, found := keeper.GetWhois(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestWhoisGetAll(t *testing.T) {
	keeper, ctx := keepertest.TelescopeKeeper(t)
	items := createNWhois(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllWhois(ctx)),
	)
}
