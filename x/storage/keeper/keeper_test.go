package keeper_test

import (
	"testing"
	gocontext "context"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	storage "github.com/jackalLabs/canine-chain/x/storage"
	"github.com/jackalLabs/canine-chain/x/storage/keeper"
	storagetestutil "github.com/jackalLabs/canine-chain/x/storage/testutil"
	"github.com/jackalLabs/canine-chain/x/storage/types"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite

	cdc           codec.Codec
	ctx           sdk.Context
	storageKeeper *keeper.Keeper
	bankKeeper    *storagetestutil.MockBankKeeper
	queryClient   types.QueryClient
	msgSrvr       types.MsgServer
}

func (suite *KeeperTestSuite) SetupSuite() {
	suite.reset()
}

func (suite *KeeperTestSuite) reset() {
	storageKeeper, bankKeeper, encCfg, ctx := setupStorageKeeper(suite.T())

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, encCfg.InterfaceRegistry)
	types.RegisterQueryServer(queryHelper, storageKeeper)
	queryClient := types.NewQueryClient(queryHelper)

	coins := sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(1000000000)))
	err := bankKeeper.MintCoins(ctx, minttypes.ModuleName, coins)
	suite.NoError(err)
	err = bankKeeper.SendCoinsFromModuleToModule(ctx, minttypes.ModuleName, types.ModuleName, coins)
	suite.NoError(err)

	suite.ctx = ctx
	suite.storageKeeper = storageKeeper
	suite.bankKeeper = bankKeeper
	suite.cdc = encCfg.Codec
	suite.queryClient = queryClient
	suite.msgSrvr = keeper.NewMsgServerImpl(*suite.storageKeeper)
}

func setupMsgServer(suite *KeeperTestSuite) (types.MsgServer, keeper.Keeper, gocontext.Context) {
	k := suite.storageKeeper
	storage.InitGenesis(suite.ctx, *k, *types.DefaultGenesis())
	ctx := sdk.WrapSDKContext(suite.ctx)
	return keeper.NewMsgServerImpl(*k), *k, ctx
}

func TestStorageTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
