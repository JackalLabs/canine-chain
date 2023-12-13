package v4_test

import (
	"fmt"
	"testing"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/golang/mock/gomock"
	minttypes "github.com/jackalLabs/canine-chain/v3/x/jklmint/types"
	oracletypes "github.com/jackalLabs/canine-chain/v3/x/oracle/types"
	storagekeeper "github.com/jackalLabs/canine-chain/v3/x/storage/keeper"
	storagetestutil "github.com/jackalLabs/canine-chain/v3/x/storage/testutil"
	storagemoduletypes "github.com/jackalLabs/canine-chain/v3/x/storage/types"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	canineglobaltestutil "github.com/jackalLabs/canine-chain/v3/testutil"
	moduletestutil "github.com/jackalLabs/canine-chain/v3/types/module/testutil" // when importing from sdk,'go mod tidy' keeps trying to import from v0.46.
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/filetree"
	"github.com/jackalLabs/canine-chain/v3/x/filetree/keeper"

	"github.com/jackalLabs/canine-chain/v3/x/filetree/types"
	"github.com/stretchr/testify/suite"
)

var modAccount = authtypes.NewModuleAddress(types.ModuleName)

// SetupFileTreeKeeper creates a filetreeKeeper as well as all its dependencies.
func SetupFileTreeKeeper(t *testing.T) (
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

// SetupStorageKeeper creates a storageKeeper as well as all its dependencies.
func SetupStorageKeeper(t *testing.T) (
	*storagekeeper.Keeper,
	*storagetestutil.MockBankKeeper,
	*storagetestutil.MockAccountKeeper,
	moduletestutil.TestEncodingConfig,
	sdk.Context,
) {
	key := sdk.NewKVStoreKey(storagemoduletypes.StoreKey)
	// memStoreKey := storetypes.NewMemoryStoreKey(storagemoduletypes.MemStoreKey)
	tkey := sdk.NewTransientStoreKey("transient_test")
	testCtx := canineglobaltestutil.DefaultContextWithDB(t, tkey, key)
	ctx := testCtx.Ctx.WithBlockHeader(tmproto.Header{Time: tmtime.Now()})

	encCfg := moduletestutil.MakeTestEncodingConfig()
	storagemoduletypes.RegisterInterfaces(encCfg.InterfaceRegistry)
	banktypes.RegisterInterfaces(encCfg.InterfaceRegistry)
	authtypes.RegisterInterfaces(encCfg.InterfaceRegistry)

	// Create MsgServiceRouter, but don't populate it before creating the storage keeper.
	msr := baseapp.NewMsgServiceRouter()

	// gomock initializations
	ctrl := gomock.NewController(t)
	bankKeeper := storagetestutil.NewMockBankKeeper(ctrl)
	accountKeeper := storagetestutil.NewMockAccountKeeper(ctrl)
	oracleKeeper := storagetestutil.NewMockOracleKeeper(ctrl)
	trackMockBalances(bankKeeper)
	accountKeeper.EXPECT().GetModuleAddress(storagemoduletypes.ModuleName).Return(modAccount).AnyTimes()

	oracleKeeper.EXPECT().GetFeed(gomock.Any(), gomock.Any()).Return(oracletypes.Feed{
		Data:  `{"price":"0.24","24h_change":"0"}`,
		Name:  "jklprice",
		Owner: "cosmos1arsaayyj5tash86mwqudmcs2fd5jt5zgp07gl8",
	}, true).AnyTimes()

	paramsSubspace := typesparams.NewSubspace(encCfg.Codec,
		storagemoduletypes.Amino,
		key,
		tkey,
		"StorageParams",
	)

	// storage keeper initializations
	storageKeeper := storagekeeper.NewKeeper(encCfg.Codec, key, paramsSubspace, bankKeeper, accountKeeper, oracleKeeper)
	storageKeeper.SetParams(ctx, storagemoduletypes.DefaultParams())

	// Register all handlers for the MegServiceRouter.
	msr.SetInterfaceRegistry(encCfg.InterfaceRegistry)
	storagemoduletypes.RegisterMsgServer(msr, storagekeeper.NewMsgServerImpl(*storageKeeper))
	banktypes.RegisterMsgServer(msr, nil) // Nil is fine here as long as we never execute the proposal's Msgs.

	return storageKeeper, bankKeeper, accountKeeper, encCfg, ctx
}

// SetupStorageKeeper creates a storageKeeper as well as all its dependencies.
func SetUpKeepers(t *testing.T) (
	*storagekeeper.Keeper,
	*keeper.Keeper,
	*storagetestutil.MockBankKeeper,
	*storagetestutil.MockAccountKeeper,
	moduletestutil.TestEncodingConfig,
	sdk.Context,
) {
	skey := sdk.NewKVStoreKey(storagemoduletypes.StoreKey)
	fkey := sdk.NewKVStoreKey(types.StoreKey)

	// memStoreKey := storetypes.NewMemoryStoreKey(storagemoduletypes.MemStoreKey)
	tkey := sdk.NewTransientStoreKey("transient_test")
	testCtx := canineglobaltestutil.DefaultContextWithDB(t, tkey, skey, fkey)
	ctx := testCtx.Ctx.WithBlockHeader(tmproto.Header{Time: tmtime.Now()})

	encCfg := moduletestutil.MakeTestEncodingConfig()
	storagemoduletypes.RegisterInterfaces(encCfg.InterfaceRegistry)
	types.RegisterInterfaces(encCfg.InterfaceRegistry)
	banktypes.RegisterInterfaces(encCfg.InterfaceRegistry)
	authtypes.RegisterInterfaces(encCfg.InterfaceRegistry)

	// Create MsgServiceRouter, but don't populate it before creating the storage keeper.
	msr := baseapp.NewMsgServiceRouter()

	// gomock initializations
	ctrl := gomock.NewController(t)
	bankKeeper := storagetestutil.NewMockBankKeeper(ctrl)
	accountKeeper := storagetestutil.NewMockAccountKeeper(ctrl)
	oracleKeeper := storagetestutil.NewMockOracleKeeper(ctrl)
	trackMockBalances(bankKeeper)
	accountKeeper.EXPECT().GetModuleAddress(storagemoduletypes.ModuleName).Return(modAccount).AnyTimes()

	oracleKeeper.EXPECT().GetFeed(gomock.Any(), gomock.Any()).Return(oracletypes.Feed{
		Data:  `{"price":"0.24","24h_change":"0"}`,
		Name:  "jklprice",
		Owner: "cosmos1arsaayyj5tash86mwqudmcs2fd5jt5zgp07gl8",
	}, true).AnyTimes()

	storParamsSubspace := typesparams.NewSubspace(encCfg.Codec,
		storagemoduletypes.Amino,
		skey,
		tkey,
		"StorageParams",
	)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	filParamsSubspace := typesparams.NewSubspace(encCfg.Codec,
		types.Amino,
		fkey,
		memStoreKey,
		"FiletreeParams",
	)

	// storage keeper initializations
	storageKeeper := storagekeeper.NewKeeper(encCfg.Codec, skey, storParamsSubspace, bankKeeper, accountKeeper, oracleKeeper)
	storageKeeper.SetParams(ctx, storagemoduletypes.DefaultParams())

	filetreeKeeper := keeper.NewKeeper(encCfg.Codec, fkey, memStoreKey, filParamsSubspace)
	filetreeKeeper.SetParams(ctx, types.DefaultParams())

	// Register all handlers for the MegServiceRouter.
	msr.SetInterfaceRegistry(encCfg.InterfaceRegistry)
	storagemoduletypes.RegisterMsgServer(msr, storagekeeper.NewMsgServerImpl(*storageKeeper))
	banktypes.RegisterMsgServer(msr, nil) // Nil is fine here as long as we never execute the proposal's Msgs.
	types.RegisterMsgServer(msr, keeper.NewMsgServerImpl(*filetreeKeeper))

	return storageKeeper, filetreeKeeper, bankKeeper, accountKeeper, encCfg, ctx
}

type UpgradeTestKeeper struct {
	suite.Suite

	cdc            codec.Codec
	ctx            sdk.Context
	filetreeKeeper *keeper.Keeper
	storageKeeper  *storagekeeper.Keeper
	queryClient    types.QueryClient
	msgSrvr        types.MsgServer
}

func (suite *UpgradeTestKeeper) SetupSuite() {
	suite.reset()
}

func (suite *UpgradeTestKeeper) reset() {
	storageKeeper, filetreeKeeper, _, _, encCfg, ctx := SetUpKeepers(suite.T())

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, encCfg.InterfaceRegistry)
	types.RegisterQueryServer(queryHelper, filetreeKeeper)
	queryClient := types.NewQueryClient(queryHelper)

	suite.ctx = ctx
	suite.filetreeKeeper = filetreeKeeper
	suite.storageKeeper = storageKeeper
	suite.cdc = encCfg.Codec
	suite.queryClient = queryClient
	suite.msgSrvr = keeper.NewMsgServerImpl(*suite.filetreeKeeper)
}

func setupMsgServer(suite *UpgradeTestKeeper) {
	k := suite.filetreeKeeper
	filetree.InitGenesis(suite.ctx, *k, *types.DefaultGenesis())
}

func TestFiletreeTestSuite(t *testing.T) {
	suite.Run(t, new(UpgradeTestKeeper))
}

// trackMockBalances sets up expected calls on the Mock BankKeeper, and also
// locally tracks accounts balances (not modules balances).
func trackMockBalances(bankKeeper *storagetestutil.MockBankKeeper) {
	balances := make(map[string]sdk.Coins)

	// We don't track module account balances.
	bankKeeper.EXPECT().MintCoins(gomock.Any(), minttypes.ModuleName, gomock.Any()).AnyTimes()
	bankKeeper.EXPECT().BurnCoins(gomock.Any(), types.ModuleName, gomock.Any()).DoAndReturn(func(_ sdk.Context, moduleName string, coins sdk.Coins) error {
		newBalance, negative := balances[modAccount.String()].SafeSub(coins)
		if negative {
			return fmt.Errorf("not enough balance")
		}
		balances[modAccount.String()] = newBalance
		return nil
	}).AnyTimes()
	bankKeeper.EXPECT().SendCoinsFromModuleToModule(gomock.Any(), minttypes.ModuleName, types.ModuleName, gomock.Any()).AnyTimes()

	// But we do track normal account balances.
	bankKeeper.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(_ sdk.Context, sender sdk.AccAddress, _ string, coins sdk.Coins) error {
		newBalance, negative := balances[sender.String()].SafeSub(coins) // in v0.46, this method is variadic
		if negative {
			return fmt.Errorf("not enough balance")
		}
		balances[sender.String()] = newBalance
		return nil
	}).AnyTimes()
	bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(_ sdk.Context, module string, rcpt sdk.AccAddress, coins sdk.Coins) error {
		balances[rcpt.String()] = balances[rcpt.String()].Add(coins...)
		return nil
	}).AnyTimes()
	bankKeeper.EXPECT().GetAllBalances(gomock.Any(), gomock.Any()).DoAndReturn(func(_ sdk.Context, addr sdk.AccAddress) sdk.Coins {
		return balances[addr.String()]
	}).AnyTimes()
	bankKeeper.EXPECT().GetBalance(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(_ sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin {
		amt := balances[addr.String()].AmountOf(denom)
		return sdk.NewCoin(denom, amt)
	}).AnyTimes()
}
