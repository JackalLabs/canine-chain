package keeper_test

import (
	"testing"

	testkeeper "dsig/testutil/keeper"

	"github.com/jackal-dao/canine/x/dsig/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DsigKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
