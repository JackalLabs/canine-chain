package keeper_test

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/dsig/keeper"
	"github.com/jackal-dao/canine/x/dsig/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

//nolint:unused
func createNUserUploads(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.UserUploads {
	items := make([]types.UserUploads, n)
	for i := range items {
		items[i].Fid = strconv.Itoa(i)

		keeper.SetUserUploads(ctx, items[i])
	}
	return items
}

/*
func TestUserUploadsGet(t *testing.T) {
	keeper, ctx := keepertest.DsigKeeper(t)
	items := createNUserUploads(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetUserUploads(ctx,
			item.Fid,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
*/

/*
func TestUserUploadsRemove(t *testing.T) {
	keeper, ctx := keepertest.DsigKeeper(t)
	items := createNUserUploads(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveUserUploads(ctx,
			item.Fid,
		)
		_, found := keeper.GetUserUploads(ctx,
			item.Fid,
		)
		require.False(t, found)
	}
}
*/

/*
func TestUserUploadsGetAll(t *testing.T) {
	keeper, ctx := keepertest.DsigKeeper(t)
	items := createNUserUploads(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllUserUploads(ctx)),
	)
}
*/
