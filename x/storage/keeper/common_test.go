package keeper_test

import (
	"fmt"
	"testing"

	oracletypes "github.com/jackalLabs/canine-chain/x/oracle/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/cosmos/cosmos-sdk/baseapp"

	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	moduletestutil "github.com/jackalLabs/canine-chain/types/module/testutil" // when importing from sdk,'go mod tidy' keeps trying to import from v0.46.

	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/golang/mock/gomock"
	canineglobaltestutil "github.com/jackalLabs/canine-chain/testutil"
	"github.com/jackalLabs/canine-chain/x/storage/keeper"
	storagetestutil "github.com/jackalLabs/canine-chain/x/storage/testutil"
	types "github.com/jackalLabs/canine-chain/x/storage/types"
)

// setupStorageKeeper creates a storageKeeper as well as all its dependencies.
func setupStorageKeeper(t *testing.T) (
	*keeper.Keeper,
	*storagetestutil.MockBankKeeper,
	*storagetestutil.MockAccountKeeper,
	moduletestutil.TestEncodingConfig,
	sdk.Context,
) {
	key := sdk.NewKVStoreKey(types.StoreKey)
	// memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)
	tkey := sdk.NewTransientStoreKey("transient_test")
	testCtx := canineglobaltestutil.DefaultContextWithDB(t, key, tkey)
	ctx := testCtx.Ctx.WithBlockHeader(tmproto.Header{Time: tmtime.Now()})

	encCfg := moduletestutil.MakeTestEncodingConfig()
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
	accountKeeper.EXPECT().GetModuleAddress(types.ModuleName).Return(modAccount).AnyTimes()

	oracleKeeper.EXPECT().GetFeed(gomock.Any(), gomock.Any()).Return(oracletypes.Feed{
		Data:  `{"price":"0.24","24h_change":"0"}`,
		Name:  "jklprice",
		Owner: "cosmos1arsaayyj5tash86mwqudmcs2fd5jt5zgp07gl8",
	}, true).AnyTimes()

	paramsSubspace := typesparams.NewSubspace(encCfg.Codec,
		types.Amino,
		key,
		tkey,
		"StorageParams",
	)

	// storage keeper initializations
	storageKeeper := keeper.NewKeeper(encCfg.Codec, key, paramsSubspace, bankKeeper, accountKeeper, oracleKeeper)
	storageKeeper.SetParams(ctx, types.DefaultParams())

	// Register all handlers for the MegServiceRouter.
	msr.SetInterfaceRegistry(encCfg.InterfaceRegistry)
	types.RegisterMsgServer(msr, keeper.NewMsgServerImpl(*storageKeeper))
	banktypes.RegisterMsgServer(msr, nil) // Nil is fine here as long as we never execute the proposal's Msgs.

	return storageKeeper, bankKeeper, accountKeeper, encCfg, ctx
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
	bankKeeper.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), gomock.Any(), types.ModuleName, gomock.Any()).DoAndReturn(func(_ sdk.Context, sender sdk.AccAddress, _ string, coins sdk.Coins) error {
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
