package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

//testing providers.go file

func (suite *KeeperTestSuite) TestSetProviders() {
	suite.SetupSuite()
	user, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	provider := types.Providers{
		Address:         user.String(),
		Ip:              "192.158.1.38",
		Totalspace:      "9000",
		BurnedContracts: "0",
		Creator:         user.String(),
	}

	suite.storageKeeper.SetProviders(suite.ctx, provider)
	suite.Require().NoError(err)

	providerRequest := types.QueryGetProvidersRequest{
		Address: user.String(),
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
	user, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	provider := types.Providers{
		Address:         user.String(),
		Ip:              "192.158.1.38",
		Totalspace:      "9000",
		BurnedContracts: "0",
		Creator:         user.String(),
	}

	suite.storageKeeper.SetProviders(suite.ctx, provider)
	suite.Require().NoError(err)

	foundProvider, found := suite.storageKeeper.GetProviders(suite.ctx, user.String())
	suite.Require().NoError(err)
	suite.Require().Equal(found, true)
	suite.Require().Equal(foundProvider.Address, provider.Address)
	suite.Require().Equal(foundProvider.Ip, provider.Ip)
	suite.Require().Equal(foundProvider.Totalspace, provider.Totalspace)
	suite.Require().Equal(foundProvider.BurnedContracts, provider.BurnedContracts)
	suite.Require().Equal(foundProvider.Creator, provider.Creator)

}

func (suite *KeeperTestSuite) TestRemoveProviders() {
	suite.SetupSuite()
	user, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	provider := types.Providers{
		Address:         user.String(),
		Ip:              "192.158.1.38",
		Totalspace:      "9000",
		BurnedContracts: "0",
		Creator:         user.String(),
	}

	suite.storageKeeper.SetProviders(suite.ctx, provider)
	suite.Require().NoError(err)

	suite.storageKeeper.RemoveProviders(suite.ctx, user.String())
	suite.Require().NoError(err)

	foundProvider, found := suite.storageKeeper.GetProviders(suite.ctx, user.String())
	suite.Require().NoError(err)
	suite.Require().Equal(found, false)
	suite.Require().Equal(nil, foundProvider)

}
