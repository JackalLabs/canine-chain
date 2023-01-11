package keeper_test

import (
	testutil "github.com/jackalLabs/canine-chain/testutil"
	"github.com/jackalLabs/canine-chain/x/storage/types"
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

	res, err := suite.queryClient.Providers(suite.ctx.Context(), &providerRequest)
	suite.Require().NoError(err)
	suite.Require().Equal(res.Providers.Address, provider.Address)
	suite.Require().Equal(res.Providers.Ip, provider.Ip)
	suite.Require().Equal(res.Providers.Totalspace, provider.Totalspace)
	suite.Require().Equal(res.Providers.BurnedContracts, provider.BurnedContracts)
	suite.Require().Equal(res.Providers.Creator, provider.Creator)
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
