package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/jackal-dao/canine/testutil/keeper"
	"github.com/jackal-dao/canine/testutil/nullify"
	"github.com/jackal-dao/canine/x/notifications/keeper"
	"github.com/jackal-dao/canine/x/notifications/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNNotifications(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Notifications {
	items := make([]types.Notifications, n)
	for i := range items {
		items[i].Count = uint64(i)

		keeper.SetNotifications(ctx, items[i])
	}
	return items
}

func TestNotificationsGet(t *testing.T) {
	keeper, ctx := keepertest.NotificationsKeeper(t)
	items := createNNotifications(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNotifications(ctx,
			item.Count,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestNotificationsRemove(t *testing.T) {
	keeper, ctx := keepertest.NotificationsKeeper(t)
	items := createNNotifications(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNotifications(ctx,
			item.Count,
			item.Address,
		)
		_, found := keeper.GetNotifications(ctx,
			item.Count,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestNotificationsGetAll(t *testing.T) {
	keeper, ctx := keepertest.NotificationsKeeper(t)
	items := createNNotifications(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNotifications(ctx)),
	)
}
