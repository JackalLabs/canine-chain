package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/themarstonconnell/telescope/testutil/keeper"
	"github.com/themarstonconnell/telescope/x/telescope/keeper"
	"github.com/themarstonconnell/telescope/x/telescope/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.TelescopeKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
