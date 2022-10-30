package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/jackalLabs/canine-chain/testutil/keeper"
	"github.com/jackalLabs/canine-chain/testutil/nullify"
	"github.com/jackalLabs/canine-chain/x/filetree/keeper"
	"github.com/jackalLabs/canine-chain/x/filetree/types"
)

func createTestTracker(keeper *keeper.Keeper, ctx sdk.Context) types.Tracker {
	item := types.Tracker{}
	keeper.SetTracker(ctx, item)
	return item
}

func TestTrackerGet(t *testing.T) {
	keeper, ctx := keepertest.FiletreeKeeper(t)
	item := createTestTracker(keeper, ctx)
	rst, found := keeper.GetTracker(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestTrackerRemove(t *testing.T) {
	keeper, ctx := keepertest.FiletreeKeeper(t)
	createTestTracker(keeper, ctx)
	keeper.RemoveTracker(ctx)
	_, found := keeper.GetTracker(ctx)
	require.False(t, found)
}
