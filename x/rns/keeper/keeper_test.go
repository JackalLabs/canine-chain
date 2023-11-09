package keeper_test

import (
	"fmt"
	"testing"

	gocontext "context"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/jackalLabs/canine-chain/v3/testutil"
	rns "github.com/jackalLabs/canine-chain/v3/x/rns"
	"github.com/jackalLabs/canine-chain/v3/x/rns/keeper"
	rnstestutil "github.com/jackalLabs/canine-chain/v3/x/rns/testutil"
	"github.com/jackalLabs/canine-chain/v3/x/rns/types"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite

	cdc         codec.Codec
	ctx         sdk.Context
	rnsKeeper   *keeper.Keeper
	bankKeeper  *rnstestutil.MockBankKeeper
	queryClient types.QueryClient
	msgSrvr     types.MsgServer
}

func (suite *KeeperTestSuite) SetupSuite() {
	suite.reset()
}

func (suite *KeeperTestSuite) reset() {
	rnsKeeper, bankKeeper, encCfg, ctx := setupRNSKeeper(suite.T())

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, encCfg.InterfaceRegistry)
	types.RegisterQueryServer(queryHelper, rnsKeeper)
	queryClient := types.NewQueryClient(queryHelper)

	coins := sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(1000000000)))
	err := bankKeeper.MintCoins(ctx, minttypes.ModuleName, coins)
	suite.NoError(err)
	err = bankKeeper.SendCoinsFromModuleToModule(ctx, minttypes.ModuleName, types.ModuleName, coins)
	suite.NoError(err)

	suite.ctx = ctx
	suite.rnsKeeper = rnsKeeper
	suite.bankKeeper = bankKeeper
	suite.cdc = encCfg.Codec
	suite.queryClient = queryClient
	suite.msgSrvr = keeper.NewMsgServerImpl(*suite.rnsKeeper)
}

func setupMsgServer(suite *KeeperTestSuite) (types.MsgServer, keeper.Keeper, gocontext.Context) {
	k := suite.rnsKeeper
	rns.InitGenesis(suite.ctx, *k, *types.DefaultGenesis())
	ctx := sdk.WrapSDKContext(suite.ctx)
	return keeper.NewMsgServerImpl(*k), *k, ctx
}

func (suite *KeeperTestSuite) setupNames() {
	testAddresses, err := testutil.CreateTestAddresses("cosmos", 1)
	suite.Require().NoError(err)

	address, err := sdk.AccAddressFromBech32(testAddresses[0])
	suite.Require().NoError(err)

	account := authtypes.NewBaseAccountWithAddress(address)

	for i := 0; i < 10; i++ {
		name := types.Names{
			Name:       fmt.Sprintf("name%d", i),
			Expires:    1000,
			Value:      account.Address,
			Data:       "{}",
			Subdomains: nil,
			Tld:        "jkl",
			Locked:     0,
		}

		suite.rnsKeeper.SetNames(suite.ctx, name)
	}
}

func (suite *KeeperTestSuite) TestMakeBid() {
	suite.SetupSuite()
	suite.setupNames()

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 1)
	suite.Require().NoError(err)

	address, err := sdk.AccAddressFromBech32(testAddresses[0])
	suite.Require().NoError(err)

	name := "name1.jkl"
	bid := types.Bids{
		Index:  fmt.Sprintf("%s%s", address.String(), name),
		Name:   name,
		Bidder: address.String(),
		Price:  "1000ujkl",
	}

	suite.rnsKeeper.SetBids(suite.ctx, bid)

	bidReq := types.QueryBid{
		Index: fmt.Sprintf("%s%s", address.String(), name),
	}

	_, err = suite.queryClient.Bid(suite.ctx.Context(), &bidReq)
	suite.Require().NoError(err)
}

func (suite *KeeperTestSuite) TestUpdateName() {
	suite.SetupSuite()

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 1)
	suite.Require().NoError(err)

	address, err := sdk.AccAddressFromBech32(testAddresses[0])
	suite.Require().NoError(err)

	account := authtypes.NewBaseAccountWithAddress(address)

	name := types.Names{
		Name:       "validname",
		Expires:    1000,
		Value:      account.Address,
		Data:       "{}",
		Subdomains: nil,
		Tld:        "jkl",
		Locked:     0,
	}

	suite.rnsKeeper.SetNames(suite.ctx, name)

	nameReq := types.QueryName{
		Index: "validname.jkl",
	}

	_, err = suite.queryClient.Name(suite.ctx.Context(), &nameReq)
	suite.Require().NoError(err)

	newData := "{\"A\":\"192.168.0.1\"}"
	name.Data = newData
	suite.rnsKeeper.SetNames(suite.ctx, name)

	res, err := suite.queryClient.Name(suite.ctx.Context(), &nameReq)
	suite.Require().NoError(err)
	suite.Require().Equal(res.Name.Data, newData)
}

func (suite *KeeperTestSuite) TestRemoveName() {
	suite.SetupSuite()

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 1)
	suite.Require().NoError(err)

	address, err := sdk.AccAddressFromBech32(testAddresses[0])
	suite.Require().NoError(err)

	account := authtypes.NewBaseAccountWithAddress(address)

	name := types.Names{
		Name:       "validname",
		Expires:    1000,
		Value:      account.Address,
		Data:       "{}",
		Subdomains: nil,
		Tld:        "jkl",
		Locked:     0,
	}

	suite.rnsKeeper.SetNames(suite.ctx, name)

	nameReq := types.QueryName{
		Index: "validname.jkl",
	}

	_, err = suite.queryClient.Name(suite.ctx.Context(), &nameReq)
	suite.Require().NoError(err)

	suite.rnsKeeper.RemoveNames(suite.ctx, "validname", "jkl")

	_, err = suite.queryClient.Name(suite.ctx.Context(), &nameReq)
	suite.Require().Error(err)
}

func (suite *KeeperTestSuite) TestSetName() {
	suite.SetupSuite()

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 1)
	suite.Require().NoError(err)

	address, err := sdk.AccAddressFromBech32(testAddresses[0])
	suite.Require().NoError(err)

	account := authtypes.NewBaseAccountWithAddress(address)

	name := types.Names{
		Name:       "validname",
		Expires:    1000,
		Value:      account.Address,
		Data:       "{}",
		Subdomains: nil,
		Tld:        "jkl",
		Locked:     0,
	}

	suite.rnsKeeper.SetNames(suite.ctx, name)

	nameReq := types.QueryName{
		Index: "validname.jkl",
	}

	_, err = suite.queryClient.Name(suite.ctx.Context(), &nameReq)
	suite.Require().NoError(err)

	badname := types.Names{
		Name:       "badname",
		Expires:    1000,
		Value:      account.Address,
		Data:       "{}",
		Subdomains: nil,
		Tld:        "non",
		Locked:     0,
	}
	suite.rnsKeeper.SetNames(suite.ctx, badname)

	nameReq = types.QueryName{
		Index: "badname.jkl",
	}
	_, err = suite.queryClient.Name(suite.ctx.Context(), &nameReq)
	suite.Require().Error(err)
}

func (suite *KeeperTestSuite) TestGRPCParams() {
	suite.SetupSuite()
	params, err := suite.queryClient.Params(gocontext.Background(), &types.QueryParams{})
	suite.Require().NoError(err)
	suite.Require().Equal(params.Params, suite.rnsKeeper.GetParams(suite.ctx))
}

func TestRnsTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
