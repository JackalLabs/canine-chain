package keeper_test

import (
	"testing"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/cosmos/cosmos-sdk/baseapp"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"

	moduletestutil "github.com/jackalLabs/canine-chain/v5/types/module/testutil" // when importing from sdk,'go mod tidy' keeps trying to import from v0.46.

	sdk "github.com/cosmos/cosmos-sdk/types"
	canineglobaltestutil "github.com/jackalLabs/canine-chain/v5/testutil"
	"github.com/jackalLabs/canine-chain/v5/x/notifications/keeper"
	types "github.com/jackalLabs/canine-chain/v5/x/notifications/types"
)

type DummyRns struct{}

func (d DummyRns) Resolve(ctx sdk.Context, name string) (sdk.AccAddress, error) {
	_ = ctx
	return sdk.AccAddressFromBech32(name)
}

// setupNotificationsKeeper creates a NotificationsKeeper as well as all its dependencies.
func setupNotificationsKeeper(t *testing.T) (
	*keeper.Keeper,
	moduletestutil.TestEncodingConfig,
	sdk.Context,
) {
	key := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)
	testCtx := canineglobaltestutil.DefaultContextWithDB(t, sdk.NewTransientStoreKey("transient_test"), key)
	ctx := testCtx.Ctx.WithBlockHeader(tmproto.Header{Time: tmtime.Now()})

	encCfg := moduletestutil.MakeTestEncodingConfig()
	types.RegisterInterfaces(encCfg.InterfaceRegistry)

	// Create MsgServiceRouter, but don't populate it before creating the Notifications keeper.
	msr := baseapp.NewMsgServiceRouter()

	paramsSubspace := typesparams.NewSubspace(encCfg.Codec,
		types.Amino,
		key,
		memStoreKey,
		"notificationsParams",
	)

	// Notifications keeper initializations
	notificationsKeeper := keeper.NewKeeper(encCfg.Codec, key, memStoreKey, paramsSubspace, DummyRns{})
	notificationsKeeper.SetParams(ctx, types.DefaultParams())

	// Register all handlers for the MegServiceRouter.
	msr.SetInterfaceRegistry(encCfg.InterfaceRegistry)
	types.RegisterMsgServer(msr, keeper.NewMsgServerImpl(*notificationsKeeper))

	return notificationsKeeper, encCfg, ctx
}
