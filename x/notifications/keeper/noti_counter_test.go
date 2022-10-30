package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/jackalLabs/canine-chain/testutil/keeper"
	"github.com/jackalLabs/canine-chain/testutil/nullify"
	"github.com/jackalLabs/canine-chain/x/notifications/keeper"
	"github.com/jackalLabs/canine-chain/x/notifications/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNNotiCounter(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.NotiCounter {
	items := make([]types.NotiCounter, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetNotiCounter(ctx, items[i])
	}
	return items
}

func TestNotiCounterGet(t *testing.T) {
	keeper, ctx := keepertest.NotificationsKeeper(t)
	items := createNNotiCounter(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNotiCounter(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestNotiCounterRemove(t *testing.T) {
	keeper, ctx := keepertest.NotificationsKeeper(t)
	items := createNNotiCounter(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNotiCounter(ctx,
			item.Address,
		)
		_, found := keeper.GetNotiCounter(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestNotiCounterGetAll(t *testing.T) {
	keeper, ctx := keepertest.NotificationsKeeper(t)
	items := createNNotiCounter(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNotiCounter(ctx)),
	)
}
