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

func createNMinerClaims(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.MinerClaims {
	items := make([]types.MinerClaims, n)
	for i := range items {
		items[i].Hash = strconv.Itoa(i)

		keeper.SetMinerClaims(ctx, items[i])
	}
	return items
}

func TestMinerClaimsGet(t *testing.T) {
	keeper, ctx := keepertest.JklminingKeeper(t)
	items := createNMinerClaims(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetMinerClaims(ctx,
			item.Hash,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestMinerClaimsRemove(t *testing.T) {
	keeper, ctx := keepertest.JklminingKeeper(t)
	items := createNMinerClaims(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMinerClaims(ctx,
			item.Hash,
		)
		_, found := keeper.GetMinerClaims(ctx,
			item.Hash,
		)
		require.False(t, found)
	}
}

func TestMinerClaimsGetAll(t *testing.T) {
	keeper, ctx := keepertest.JklminingKeeper(t)
	items := createNMinerClaims(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMinerClaims(ctx)),
	)
}
