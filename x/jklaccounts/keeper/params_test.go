package keeper_test

import (
	"testing"

	testkeeper "github.com/jackal-dao/canine/testutil/keeper"
	"github.com/jackal-dao/canine/x/jklaccounts/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.JklaccountsKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
