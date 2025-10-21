package keeper_test

import (
	"testing"

	gocontext "context"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	oracle "github.com/jackalLabs/canine-chain/v5/x/oracle"
	"github.com/jackalLabs/canine-chain/v5/x/oracle/keeper"
	oracletestutil "github.com/jackalLabs/canine-chain/v5/x/oracle/testutil"
	"github.com/jackalLabs/canine-chain/v5/x/oracle/types"
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
	testAccs     []sdk.AccAddress
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
	suite.testAccs = CreateRandomAccounts(3)
}

func setupMsgServer(suite *KeeperTestSuite) (types.MsgServer, gocontext.Context) {
	k := suite.oracleKeeper
	oracle.InitGenesis(suite.ctx, *k, *types.DefaultGenesis())
	ctx := sdk.WrapSDKContext(suite.ctx)
	return keeper.NewMsgServerImpl(*k), ctx
}

func CreateRandomAccounts(numAccs int) []sdk.AccAddress {
	accs := make([]sdk.AccAddress, numAccs)

	for i := 0; i < numAccs; i++ {
		pk := secp256k1.GenPrivKey().PubKey()
		accs[i] = sdk.AccAddress(pk.Address())
	}

	return accs
}

func (suite *KeeperTestSuite) TestGRPCParams() {
	suite.SetupSuite()
	params, err := suite.queryClient.Params(gocontext.Background(), &types.QueryParams{})
	suite.Require().NoError(err)
	suite.Require().Equal(params.Params, suite.oracleKeeper.GetParams(suite.ctx))
}

func TestOracleTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
