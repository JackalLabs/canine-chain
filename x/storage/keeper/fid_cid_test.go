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

func createNFidCid(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.FidCid {
	items := make([]types.FidCid, n)
	for i := range items {
		items[i].Fid = strconv.Itoa(i)

		keeper.SetFidCid(ctx, items[i])
	}
	return items
}

func TestFidCidGet(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	items := createNFidCid(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetFidCid(ctx,
			item.Fid,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestFidCidRemove(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	items := createNFidCid(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveFidCid(ctx,
			item.Fid,
		)
		_, found := keeper.GetFidCid(ctx,
			item.Fid,
		)
		require.False(t, found)
	}
}

func TestFidCidGetAll(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	items := createNFidCid(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllFidCid(ctx)),
	)
}
