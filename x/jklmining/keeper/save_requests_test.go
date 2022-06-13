package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/jackal-dao/canine/testutil/keeper"
	"github.com/jackal-dao/canine/testutil/nullify"
	"github.com/jackal-dao/canine/x/jklmining/keeper"
	"github.com/jackal-dao/canine/x/jklmining/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNSaveRequests(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SaveRequests {
	items := make([]types.SaveRequests, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetSaveRequests(ctx, items[i])
	}
	return items
}

func TestSaveRequestsGet(t *testing.T) {
	keeper, ctx := keepertest.JklminingKeeper(t)
	items := createNSaveRequests(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetSaveRequests(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestSaveRequestsRemove(t *testing.T) {
	keeper, ctx := keepertest.JklminingKeeper(t)
	items := createNSaveRequests(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSaveRequests(ctx,
			item.Index,
		)
		_, found := keeper.GetSaveRequests(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestSaveRequestsGetAll(t *testing.T) {
	keeper, ctx := keepertest.JklminingKeeper(t)
	items := createNSaveRequests(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSaveRequests(ctx)),
	)
}
