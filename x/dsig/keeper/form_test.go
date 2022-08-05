package keeper_test

import (
	"strconv"
	"testing"

	keepertest "dsig/testutil/keeper"
	"dsig/testutil/nullify"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/dsig/keeper"
	"github.com/jackal-dao/canine/x/dsig/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNForm(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Form {
	items := make([]types.Form, n)
	for i := range items {
		items[i].Ffid = strconv.Itoa(i)

		keeper.SetForm(ctx, items[i])
	}
	return items
}

func TestFormGet(t *testing.T) {
	keeper, ctx := keepertest.DsigKeeper(t)
	items := createNForm(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetForm(ctx,
			item.Ffid,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestFormRemove(t *testing.T) {
	keeper, ctx := keepertest.DsigKeeper(t)
	items := createNForm(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveForm(ctx,
			item.Ffid,
		)
		_, found := keeper.GetForm(ctx,
			item.Ffid,
		)
		require.False(t, found)
	}
}

func TestFormGetAll(t *testing.T) {
	keeper, ctx := keepertest.DsigKeeper(t)
	items := createNForm(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllForm(ctx)),
	)
}
