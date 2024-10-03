package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/testutil"
	"github.com/jackalLabs/canine-chain/v4/x/rns/keeper"
	"github.com/jackalLabs/canine-chain/v4/x/rns/types"
)

func (suite *KeeperTestSuite) TestMsgTransfer() {
	suite.SetupSuite()
	suite.setupNames()

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	suite.Require().NoError(err)

	address, err := sdk.AccAddressFromBech32(testAddresses[0])
	suite.Require().NoError(err)
	receiver, err := sdk.AccAddressFromBech32(testAddresses[1])
	suite.Require().NoError(err)

	name := "test.jkl"

	coin := sdk.NewCoin("ujkl", sdk.NewInt(1000000000))
	coins := sdk.NewCoins(coin)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, address, coins)
	suite.Require().NoError(err)

	beforebal := suite.bankKeeper.GetAllBalances(suite.ctx, address)
	amt := beforebal.AmountOf("ujkl")

	err = suite.rnsKeeper.RegisterRNSName(suite.ctx, address.String(), name, "{}", 2, true)
	suite.Require().NoError(err)

	nameReq := types.QueryName{
		Name: name,
	}

	afterbal := suite.bankKeeper.GetAllBalances(suite.ctx, address)
	newamt := afterbal.AmountOf("ujkl")

	n, t, err := keeper.GetNameAndTLD(name)
	suite.Require().NoError(err)

	cost, err := keeper.GetCostOfName(n, t)
	suite.Require().NoError(err)

	newamt = amt.Sub(newamt)
	leftover := cost * 2
	suite.Require().Equal(newamt.Int64(), leftover) // cost them the amount they bid

	_, err = suite.queryClient.Name(suite.ctx.Context(), &nameReq)
	suite.Require().NoError(err)

	err = suite.rnsKeeper.TransferName(suite.ctx, address.String(), receiver.String(), name) // will pass as the user owns the name
	suite.Require().NoError(err)

	res, err := suite.queryClient.Name(suite.ctx.Context(), &nameReq)
	suite.Require().NoError(err)
	suite.Require().Equal(res.Name.Value, receiver.String())

	err = suite.rnsKeeper.TransferName(suite.ctx, address.String(), receiver.String(), name) // should fail sending a name from an address that doesn't have ownership
	suite.Require().Error(err)

	res, err = suite.queryClient.Name(suite.ctx.Context(), &nameReq)
	suite.Require().NoError(err)
	suite.Require().Equal(res.Name.Value, receiver.String())
}
