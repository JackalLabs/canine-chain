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

func createNProviders(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Providers {
	items := make([]types.Providers, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetProviders(ctx, items[i])
	}
	return items
}

func TestProvidersGet(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	items := createNProviders(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetProviders(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestProvidersRemove(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	items := createNProviders(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveProviders(ctx,
			item.Address,
		)
		_, found := keeper.GetProviders(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestProvidersGetAll(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	items := createNProviders(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllProviders(ctx)),
	)
}
