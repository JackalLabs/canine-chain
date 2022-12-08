package keeper_test

import (
	"testing"

	gocontext "context"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/jackalLabs/canine-chain/x/oracle/keeper"
	oracletestutil "github.com/jackalLabs/canine-chain/x/oracle/testutil"
	"github.com/jackalLabs/canine-chain/x/oracle/types"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite

	cdc          codec.Codec
	ctx          sdk.Context
	oracleKeeper *keeper.Keeper
	bankKeeper   *oracletestutil.MockBankKeeper
	queryClient  types.QueryClient
	msgSrvr      types.MsgServer
}

func (suite *KeeperTestSuite) SetupSuite() {
	suite.reset()
}

func (suite *KeeperTestSuite) reset() {
	oracleKeeper, bankKeeper, encCfg, ctx := setupOracleKeeper(suite.T())

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, encCfg.InterfaceRegistry)
	types.RegisterQueryServer(queryHelper, oracleKeeper)
	queryClient := types.NewQueryClient(queryHelper)

	coins := sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(1000000000)))
	err := bankKeeper.MintCoins(ctx, minttypes.ModuleName, coins)
	suite.NoError(err)
	err = bankKeeper.SendCoinsFromModuleToModule(ctx, minttypes.ModuleName, types.ModuleName, coins)
	suite.NoError(err)

	suite.ctx = ctx
	suite.oracleKeeper = oracleKeeper
	suite.bankKeeper = bankKeeper
	suite.cdc = encCfg.Codec
	suite.queryClient = queryClient
	suite.msgSrvr = keeper.NewMsgServerImpl(*suite.oracleKeeper)
}

// TODO: add msgServer tests
// func setupMsgServer(suite *KeeperTestSuite) (types.MsgServer, keeper.Keeper, gocontext.Context) {
//	 k := suite.oracleKeeper
//	 oracle.InitGenesis(suite.ctx, *k, *types.DefaultGenesis())
//	 ctx := sdk.WrapSDKContext(suite.ctx)
//	 return keeper.NewMsgServerImpl(*k), *k, ctx
// }

func (suite *KeeperTestSuite) TestGRPCParams() {
	suite.SetupSuite()
	params, err := suite.queryClient.Params(gocontext.Background(), &types.QueryParamsRequest{})
	suite.Require().NoError(err)
	suite.Require().Equal(params.Params, suite.oracleKeeper.GetParams(suite.ctx))
}

func TestOracleTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
