package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/jackalLabs/canine-chain/testutil/keeper"
	"github.com/jackalLabs/canine-chain/testutil/nullify"
	"github.com/jackalLabs/canine-chain/x/storage/keeper"
	"github.com/jackalLabs/canine-chain/x/storage/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNStrays(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Strays {
	items := make([]types.Strays, n)
	for i := range items {
		items[i].Cid = strconv.Itoa(i)

		keeper.SetStrays(ctx, items[i])
	}
	return items
}

func TestStraysGet(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	items := createNStrays(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetStrays(ctx,
			item.Cid,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestStraysRemove(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	items := createNStrays(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStrays(ctx,
			item.Cid,
		)
		_, found := keeper.GetStrays(ctx,
			item.Cid,
		)
		require.False(t, found)
	}
}

func TestStraysGetAll(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	items := createNStrays(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStrays(ctx)),
	)
}
