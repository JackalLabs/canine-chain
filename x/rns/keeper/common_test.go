package keeper_test

import (
	"fmt"
	"testing"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/cosmos/cosmos-sdk/baseapp"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	moduletestutil "github.com/jackalLabs/canine-chain/types/module/testutil" //when importing from sdk,'go mod tidy' keeps trying to import from v0.46.

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/golang/mock/gomock"
	canineglobaltestutil "github.com/jackalLabs/canine-chain/testutil"
	"github.com/jackalLabs/canine-chain/x/rns/keeper"
	rnstestutil "github.com/jackalLabs/canine-chain/x/rns/testutil"
	types "github.com/jackalLabs/canine-chain/x/rns/types"
)

var (
	_, _, addr = testdata.KeyTestPubAddr()
	rnsAcct    = authtypes.NewModuleAddress(types.ModuleName)
	//1 rns test to go here
)

//1 rns test to go here

// setupRNSKeeper creates a rnsKeeper as well as all its dependencies.
func setupRNSKeeper(t *testing.T) (
	*keeper.Keeper,
	*rnstestutil.MockBankKeeper,
	moduletestutil.TestEncodingConfig,
	sdk.Context,
) {
	key := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)
	testCtx := canineglobaltestutil.DefaultContextWithDB(t, key, sdk.NewTransientStoreKey("transient_test"))
	ctx := testCtx.Ctx.WithBlockHeader(tmproto.Header{Time: tmtime.Now()})

	encCfg := moduletestutil.MakeTestEncodingConfig()
	types.RegisterInterfaces(encCfg.InterfaceRegistry)
	banktypes.RegisterInterfaces(encCfg.InterfaceRegistry)

	// Create MsgServiceRouter, but don't populate it before creating the rns keeper.
	msr := baseapp.NewMsgServiceRouter()

	// gomock initalizations
	ctrl := gomock.NewController(t)
	bankKeeper := rnstestutil.NewMockBankKeeper(ctrl)
	trackMockBalances(bankKeeper)

	paramsSubspace := typesparams.NewSubspace(encCfg.Codec,
		types.Amino,
		key,
		memStoreKey,
		"RNSParams",
	)

	// rns keeper initializations
	rnsKeeper := keeper.NewKeeper(encCfg.Codec, key, memStoreKey, paramsSubspace, bankKeeper)
	rnsKeeper.SetParams(ctx, types.DefaultParams())

	// Register all handlers for the MegServiceRouter.
	msr.SetInterfaceRegistry(encCfg.InterfaceRegistry)
	types.RegisterMsgServer(msr, keeper.NewMsgServerImpl(*rnsKeeper))
	banktypes.RegisterMsgServer(msr, nil) // Nil is fine here as long as we never execute the proposal's Msgs.

	return rnsKeeper, bankKeeper, encCfg, ctx

}

// trackMockBalances sets up expected calls on the Mock BankKeeper, and also
// locally tracks accounts balances (not modules balances).
func trackMockBalances(bankKeeper *rnstestutil.MockBankKeeper) {
	balances := make(map[string]sdk.Coins)

	// We don't track module account balances.
	bankKeeper.EXPECT().MintCoins(gomock.Any(), minttypes.ModuleName, gomock.Any()).AnyTimes()
	bankKeeper.EXPECT().BurnCoins(gomock.Any(), types.ModuleName, gomock.Any()).AnyTimes()
	bankKeeper.EXPECT().SendCoinsFromModuleToModule(gomock.Any(), minttypes.ModuleName, types.ModuleName, gomock.Any()).AnyTimes()

	// But we do track normal account balances.
	bankKeeper.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), gomock.Any(), types.ModuleName, gomock.Any()).DoAndReturn(func(_ sdk.Context, sender sdk.AccAddress, _ string, coins sdk.Coins) error {
		newBalance, negative := balances[sender.String()].SafeSub(coins) //in v0.46, this method is variadic
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
}
