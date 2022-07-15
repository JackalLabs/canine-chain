package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/jackal-dao/canine/testutil/keeper"
	"github.com/jackal-dao/canine/testutil/nullify"
	"github.com/jackal-dao/canine/x/storage/keeper"
	"github.com/jackal-dao/canine/x/storage/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNProofs(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Proofs {
	items := make([]types.Proofs, n)
	for i := range items {
		items[i].Cid = strconv.Itoa(i)

		keeper.SetProofs(ctx, items[i])
	}
	return items
}

func TestProofsGet(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	items := createNProofs(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetProofs(ctx,
			item.Cid,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestProofsRemove(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	items := createNProofs(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveProofs(ctx,
			item.Cid,
		)
		_, found := keeper.GetProofs(ctx,
			item.Cid,
		)
		require.False(t, found)
	}
}

func TestProofsGetAll(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	items := createNProofs(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllProofs(ctx)),
	)
}
