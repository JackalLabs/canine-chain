package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/jackal-dao/canine/testutil/keeper"
	"github.com/jackal-dao/canine/testutil/nullify"
	"github.com/jackal-dao/canine/x/jklaccounts/keeper"
	"github.com/jackal-dao/canine/x/jklaccounts/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNAccounts(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Accounts {
	items := make([]types.Accounts, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetAccounts(ctx, items[i])
	}
	return items
}

func TestAccountsGet(t *testing.T) {
	keeper, ctx := keepertest.JklaccountsKeeper(t)
	items := createNAccounts(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetAccounts(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestAccountsRemove(t *testing.T) {
	keeper, ctx := keepertest.JklaccountsKeeper(t)
	items := createNAccounts(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAccounts(ctx,
			item.Address,
		)
		_, found := keeper.GetAccounts(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestAccountsGetAll(t *testing.T) {
	keeper, ctx := keepertest.JklaccountsKeeper(t)
	items := createNAccounts(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAccounts(ctx)),
	)
}
