package keeper_test

import (
	gocontext "context"
	"testing"
	"time"

	"github.com/jackalLabs/canine-chain/v5/testutil"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	storage "github.com/jackalLabs/canine-chain/v5/x/storage"
	"github.com/jackalLabs/canine-chain/v5/x/storage/keeper"
	storagetestutil "github.com/jackalLabs/canine-chain/v5/x/storage/testutil"
	"github.com/jackalLabs/canine-chain/v5/x/storage/types"
	"github.com/stretchr/testify/suite"
)

var modAccount = authtypes.NewModuleAddress(types.ModuleName)

type KeeperTestSuite struct {
	suite.Suite

	cdc           codec.Codec
	ctx           sdk.Context
	storageKeeper *keeper.Keeper
	bankKeeper    *storagetestutil.MockBankKeeper
	accountKeeper *storagetestutil.MockAccountKeeper
	queryClient   types.QueryClient
	msgSrvr       types.MsgServer
}

func (suite *KeeperTestSuite) SetupSuite() {
	suite.reset()
}

func (suite *KeeperTestSuite) reset() {
	storageKeeper, bankKeeper, accountKeeper, encCfg, ctx := setupStorageKeeper(suite.T())

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, encCfg.InterfaceRegistry)
	types.RegisterQueryServer(queryHelper, storageKeeper)
	queryClient := types.NewQueryClient(queryHelper)

	// accountKeeper.EXPECT().GetModuleAccount(gomock.Any(), types.ModuleName).Return(authtypes.NewEmptyModuleAccount(types.ModuleName)).AnyTimes()

	coins := sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(1000000000)))
	err := bankKeeper.MintCoins(ctx, minttypes.ModuleName, coins)
	suite.NoError(err)
	err = bankKeeper.SendCoinsFromModuleToModule(ctx, minttypes.ModuleName, types.ModuleName, coins)
	suite.NoError(err)

	suite.ctx = ctx.WithBlockTime(time.Now())
	suite.storageKeeper = storageKeeper
	suite.bankKeeper = bankKeeper
	suite.accountKeeper = accountKeeper
	suite.cdc = encCfg.Codec
	suite.queryClient = queryClient
	suite.msgSrvr = keeper.NewMsgServerImpl(*suite.storageKeeper)
}

func setupMsgServer(suite *KeeperTestSuite) (types.MsgServer, keeper.Keeper, gocontext.Context) {
	k := suite.storageKeeper
	storage.InitGenesis(suite.ctx, *k, *types.DefaultGenesis())
	ctx := sdk.WrapSDKContext(suite.ctx)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 3)
	suite.Require().NoError(err)

	depoAccount := testAddresses[0]

	k.SetParams(suite.ctx, types.Params{
		DepositAccount:         depoAccount,
		ProofWindow:            50,
		ChunkSize:              1024,
		PriceFeed:              "jklprice",
		MissesToBurn:           3,
		MaxContractAgeInBlocks: 100,
		PricePerTbPerMonth:     15,
		CollateralPrice:        2,
		CheckWindow:            11,
		ReferralCommission:     25,
		PolRatio:               40,
	})

	return keeper.NewMsgServerImpl(*k), *k, ctx
}

func TestStorageTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
