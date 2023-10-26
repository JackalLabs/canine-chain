package v4_test

import (
	gocontext "context"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	canineglobaltestutil "github.com/jackalLabs/canine-chain/v3/testutil"
	moduletestutil "github.com/jackalLabs/canine-chain/v3/types/module/testutil" // when importing from sdk,'go mod tidy' keeps trying to import from v0.46.
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/filetree"
	"github.com/jackalLabs/canine-chain/v3/x/filetree/keeper"

	"github.com/jackalLabs/canine-chain/v3/x/filetree/types"
	"github.com/stretchr/testify/suite"
)

// SetupFileTreeKeeper creates a filetreeKeeper as well as all its dependencies.
func SetupFileTreeKeeper(t *testing.T) (
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

type UpgradeTestKeeper struct {
	suite.Suite

	cdc            codec.Codec
	ctx            sdk.Context
	filetreeKeeper *keeper.Keeper
	queryClient    types.QueryClient
	msgSrvr        types.MsgServer
}

func (suite *UpgradeTestKeeper) SetupSuite() {
	suite.reset()
}

func (suite *UpgradeTestKeeper) reset() {
	filetreeKeeper, encCfg, ctx := SetupFileTreeKeeper(suite.T())

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, encCfg.InterfaceRegistry)
	types.RegisterQueryServer(queryHelper, filetreeKeeper)
	queryClient := types.NewQueryClient(queryHelper)

	suite.ctx = ctx
	suite.filetreeKeeper = filetreeKeeper
	suite.cdc = encCfg.Codec
	suite.queryClient = queryClient
	suite.msgSrvr = keeper.NewMsgServerImpl(*suite.filetreeKeeper)
}

func setupMsgServer(suite *UpgradeTestKeeper) (types.MsgServer, gocontext.Context) {
	k := suite.filetreeKeeper
	filetree.InitGenesis(suite.ctx, *k, *types.DefaultGenesis())
	ctx := sdk.WrapSDKContext(suite.ctx)
	return keeper.NewMsgServerImpl(*k), ctx
}

func TestFiletreeTestSuite(t *testing.T) {
	suite.Run(t, new(UpgradeTestKeeper))
}
