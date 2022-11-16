package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/jackal-dao/canine/testutil/keeper"
	"github.com/jackal-dao/canine/testutil/nullify"
	"github.com/jackal-dao/canine/x/filetree/keeper"
	"github.com/jackal-dao/canine/x/filetree/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNPubkey(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Pubkey {
	items := make([]types.Pubkey, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetPubkey(ctx, items[i])
	}
	return items
}

func TestPubkeyGet(t *testing.T) {
	keeper, ctx := keepertest.FiletreeKeeper(t)
	items := createNPubkey(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPubkey(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPubkeyRemove(t *testing.T) {
	keeper, ctx := keepertest.FiletreeKeeper(t)
	items := createNPubkey(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePubkey(ctx,
			item.Address,
		)
		_, found := keeper.GetPubkey(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestPubkeyGetAll(t *testing.T) {
	keeper, ctx := keepertest.FiletreeKeeper(t)
	items := createNPubkey(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPubkey(ctx)),
	)
}
