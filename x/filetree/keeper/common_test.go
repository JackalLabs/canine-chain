package keeper_test

import (
	"testing"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/cosmos/cosmos-sdk/baseapp"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"

	moduletestutil "github.com/jackalLabs/canine-chain/v3/types/module/testutil" // when importing from sdk,'go mod tidy' keeps trying to import from v0.46.

	sdk "github.com/cosmos/cosmos-sdk/types"
	canineglobaltestutil "github.com/jackalLabs/canine-chain/v3/testutil"
	"github.com/jackalLabs/canine-chain/v3/x/filetree/keeper"
	types "github.com/jackalLabs/canine-chain/v3/x/filetree/types"
)

// setupFiletreeKeeper creates a filetreeKeeper as well as all its dependencies.
func setupFiletreeKeeper(t *testing.T) (
	*keeper.Keeper,
	moduletestutil.TestEncodingConfig,
	sdk.Context,
) {
	key := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)
	testCtx := canineglobaltestutil.DefaultContextWithDB(t, key, sdk.NewTransientStoreKey("transient_test"))
	ctx := testCtx.Ctx.WithBlockHeader(tmproto.Header{Time: tmtime.Now()})

	encCfg := moduletestutil.MakeTestEncodingConfig()
	types.RegisterInterfaces(encCfg.InterfaceRegistry)

	// Create MsgServiceRouter, but don't populate it before creating the filetree keeper.
	msr := baseapp.NewMsgServiceRouter()

	paramsSubspace := typesparams.NewSubspace(encCfg.Codec,
		types.Amino,
		key,
		memStoreKey,
		"FiletreeParams",
	)

	// filetree keeper initializations
	filetreeKeeper := keeper.NewKeeper(encCfg.Codec, key, memStoreKey, paramsSubspace)
	filetreeKeeper.SetParams(ctx, types.DefaultParams())

	// Register all handlers for the MegServiceRouter.
	msr.SetInterfaceRegistry(encCfg.InterfaceRegistry)
	types.RegisterMsgServer(msr, keeper.NewMsgServerImpl(*filetreeKeeper))

	return filetreeKeeper, encCfg, ctx
}
