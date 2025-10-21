package keeper_test

import (
	gocontext "context"
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	filetree "github.com/jackalLabs/canine-chain/v5/x/filetree"
	"github.com/jackalLabs/canine-chain/v5/x/filetree/keeper"
	"github.com/jackalLabs/canine-chain/v5/x/filetree/types"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite

	cdc            codec.Codec
	ctx            sdk.Context
	filetreeKeeper *keeper.Keeper
	queryClient    types.QueryClient
	msgSrvr        types.MsgServer
}

func (suite *KeeperTestSuite) SetupSuite() {
	suite.reset()
}

func (suite *KeeperTestSuite) reset() {
	filetreeKeeper, encCfg, ctx := SetupFiletreeKeeper(suite.T())

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, encCfg.InterfaceRegistry)
	types.RegisterQueryServer(queryHelper, filetreeKeeper)
	queryClient := types.NewQueryClient(queryHelper)

	suite.ctx = ctx
	suite.filetreeKeeper = filetreeKeeper
	suite.cdc = encCfg.Codec
	suite.queryClient = queryClient
	suite.msgSrvr = keeper.NewMsgServerImpl(*suite.filetreeKeeper)
}

func setupMsgServer(suite *KeeperTestSuite) (types.MsgServer, gocontext.Context) {
	k := suite.filetreeKeeper
	filetree.InitGenesis(suite.ctx, *k, *types.DefaultGenesis())
	ctx := sdk.WrapSDKContext(suite.ctx)
	return keeper.NewMsgServerImpl(*k), ctx
}

func TestFiletreeTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
