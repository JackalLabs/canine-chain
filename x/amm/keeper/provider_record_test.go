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

func createNProviderRecord(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ProviderRecord {
	items := make([]types.ProviderRecord, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetProviderRecord(ctx, items[i])
	}
	return items
}

func TestProviderRecordGet(t *testing.T) {
	keeper, ctx := keepertest.LpKeeper(t)
	items := createNProviderRecord(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetProviderRecord(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestProviderRecordRemove(t *testing.T) {
	keeper, ctx := keepertest.LpKeeper(t)
	items := createNProviderRecord(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveProviderRecord(ctx,
			item.Index,
		)
		_, found := keeper.GetProviderRecord(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestProviderRecordGetAll(t *testing.T) {
	keeper, ctx := keepertest.LpKeeper(t)
	items := createNProviderRecord(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllProviderRecord(ctx)),
	)
}
