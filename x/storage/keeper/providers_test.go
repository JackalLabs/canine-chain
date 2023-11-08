package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	testutil "github.com/jackalLabs/canine-chain/v3/testutil"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

// testing providers.go file
func (suite *KeeperTestSuite) TestSetProviders() {
	suite.SetupSuite()

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 1)
	suite.Require().NoError(err)

	user := testAddresses[0]

	provider := types.Providers{
		Address:         user,
		Ip:              "192.158.1.38",
		Totalspace:      "9000",
		BurnedContracts: "0",
		Creator:         user,
	}

	suite.storageKeeper.SetProviders(suite.ctx, provider)
	suite.Require().NoError(err)

	providerRequest := types.QueryProviderRequest{
		Address: user,
	}

	res, err := suite.queryClient.Provider(suite.ctx.Context(), &providerRequest)
	suite.Require().NoError(err)
	suite.Require().Equal(res.Provider.Address, provider.Address)
	suite.Require().Equal(res.Provider.Ip, provider.Ip)
	suite.Require().Equal(res.Provider.Totalspace, provider.Totalspace)
	suite.Require().Equal(res.Provider.BurnedContracts, provider.BurnedContracts)
	suite.Require().Equal(res.Provider.Creator, provider.Creator)
}

// testing providers.go file
func (suite *KeeperTestSuite) TestInitProviders() {
	suite.SetupSuite()
	msgSrvr, _, ctx := setupMsgServer(suite)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 1)
	suite.Require().NoError(err)

	user := testAddresses[0]

	deposit := sdk.NewCoin("ujkl", sdk.NewInt(10_000_000_000))

	coins := sdk.NewCoins(deposit) // Send some coins to their account
	userAcc, _ := sdk.AccAddressFromBech32(user)
	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, userAcc, coins)

	suite.Require().NoError(err)

	initMsg := types.MsgInitProvider{
		Creator:    user,
		Ip:         "192.158.1.38",
		Keybase:    "",
		TotalSpace: "9000",
	}

	_, err = msgSrvr.InitProvider(ctx, &initMsg)
	suite.Require().NoError(err)

	providerRequest := types.QueryProviderRequest{
		Address: user,
	}

	res, err := suite.queryClient.Provider(suite.ctx.Context(), &providerRequest)
	suite.Require().NoError(err)
	suite.Require().Equal(res.Provider.Address, user)
	suite.Require().Equal(res.Provider.Ip, initMsg.Ip)
	suite.Require().Equal(res.Provider.Totalspace, initMsg.TotalSpace)
	suite.Require().Equal(res.Provider.BurnedContracts, "0")
	suite.Require().Equal(res.Provider.Creator, initMsg.Creator)

	coin := suite.bankKeeper.GetBalance(suite.ctx, userAcc, "ujkl")

	suite.Require().Equal(sdk.NewInt(0), coin.Amount)

	shutdownMsg := types.MsgShutdownProvider{
		Creator: user,
	}
	_, err = msgSrvr.ShutdownProvider(ctx, &shutdownMsg)
	suite.Require().NoError(err)

	coin = suite.bankKeeper.GetBalance(suite.ctx, userAcc, "ujkl")

	suite.Require().Equal(deposit.Amount, coin.Amount)
}

func (suite *KeeperTestSuite) TestGetProviders() {
	suite.SetupSuite()

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 1)
	suite.Require().NoError(err)

	user := testAddresses[0]

	provider := types.Providers{
		Address:         user,
		Ip:              "192.158.1.38",
		Totalspace:      "9000",
		BurnedContracts: "0",
		Creator:         user,
	}

	suite.storageKeeper.SetProviders(suite.ctx, provider)
	suite.Require().NoError(err)

	foundProvider, found := suite.storageKeeper.GetProviders(suite.ctx, user)
	suite.Require().NoError(err)
	suite.Require().Equal(found, true)
	suite.Require().Equal(foundProvider.Address, provider.Address)
	suite.Require().Equal(foundProvider.Ip, provider.Ip)
	suite.Require().Equal(foundProvider.Totalspace, provider.Totalspace)
	suite.Require().Equal(foundProvider.BurnedContracts, provider.BurnedContracts)
	suite.Require().Equal(foundProvider.Creator, provider.Creator)
}

func (suite *KeeperTestSuite) TestGetAllProviders() {
	suite.SetupSuite()

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	suite.Require().NoError(err)

	alice := testAddresses[0]
	bob := testAddresses[1]

	provider := types.Providers{
		Address:         alice,
		Ip:              "192.158.1.38",
		Totalspace:      "9000",
		BurnedContracts: "0",
		Creator:         alice,
	}

	allProvidersbefore := suite.storageKeeper.GetAllProviders(suite.ctx)
	suite.Require().Equal(0, len(allProvidersbefore))

	suite.storageKeeper.SetProviders(suite.ctx, provider)

	provider1 := types.Providers{
		Address:         bob,
		Ip:              "127.159.2.39",
		Totalspace:      "18000",
		BurnedContracts: "0",
		Creator:         bob,
	}

	suite.storageKeeper.SetProviders(suite.ctx, provider1)

	allProviders := suite.storageKeeper.GetAllProviders(suite.ctx)
	suite.Require().Equal(2, len(allProviders))
}

func (suite *KeeperTestSuite) TestRemoveProviders() {
	suite.SetupSuite()

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 1)
	suite.Require().NoError(err)

	user := testAddresses[0]

	provider := types.Providers{
		Address:         user,
		Ip:              "192.158.1.38",
		Totalspace:      "9000",
		BurnedContracts: "0",
		Creator:         user,
	}

	suite.storageKeeper.SetProviders(suite.ctx, provider)
	suite.Require().NoError(err)

	suite.storageKeeper.RemoveProviders(suite.ctx, user)
	suite.Require().NoError(err)

	foundProvider, found := suite.storageKeeper.GetProviders(suite.ctx, user)
	suite.Require().NoError(err)
	suite.Require().Equal(found, false)

	ghostProvider := types.Providers{
		Address:         "",
		Ip:              "",
		Totalspace:      "",
		BurnedContracts: "",
		Creator:         "",
	}

	suite.Require().Equal(foundProvider, ghostProvider)
}

func (suite *KeeperTestSuite) TestActiveProviders() {
	suite.SetupSuite()

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 1)
	suite.Require().NoError(err)

	user := testAddresses[0]

	realProvider := types.Providers{
		Address: user,
		Ip:      "https://test.com",
	}

	suite.storageKeeper.SetProviders(suite.ctx, realProvider)

	provider := types.ActiveProviders{
		Address: user,
	}

	suite.storageKeeper.SetActiveProviders(suite.ctx, provider)

	foundProvider := suite.storageKeeper.GetActiveProviders(suite.ctx, "")
	suite.Require().Equal(1, len(foundProvider))

	foundProvider = suite.storageKeeper.GetActiveProviders(suite.ctx, "https://test.com")
	suite.Require().Equal(0, len(foundProvider))

	suite.storageKeeper.RemoveAllActiveProviders(suite.ctx)

	foundProvider = suite.storageKeeper.GetActiveProviders(suite.ctx, "")
	suite.Require().Equal(0, len(foundProvider))
}
