package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/themarstonconnell/telescope/testutil/keeper"
	"github.com/themarstonconnell/telescope/x/telescope/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.TelescopeKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
