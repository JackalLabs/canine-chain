package keeper_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/testutil"
	"github.com/jackalLabs/canine-chain/v4/x/rns/keeper"
	"github.com/jackalLabs/canine-chain/v4/x/rns/types"
)

func (suite *KeeperTestSuite) TestMsgRegisterName() {
	suite.SetupSuite()
	suite.setupNames()

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 1)
	suite.Require().NoError(err)

	address, err := sdk.AccAddressFromBech32(testAddresses[0])
	suite.Require().NoError(err)

	name := "test.jkl"
	capname := "Test.jkl"

	n, t, err := keeper.GetNameAndTLD(name)
	suite.Require().NoError(err)

	cost, err := keeper.GetCostOfName(n, t)
	suite.Require().NoError(err)

	coin := sdk.NewCoin("ujkl", sdk.NewInt(10000000000000000))
	coins := sdk.NewCoins(coin)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, address, coins)
	suite.Require().NoError(err)

	beforebal := suite.bankKeeper.GetAllBalances(suite.ctx, address)
	amt := beforebal.AmountOf("ujkl")

	err = suite.rnsKeeper.RegisterRNSName(suite.ctx, address.String(), name, "{}", 2, true)
	suite.Require().NoError(err)

	primName, f := suite.rnsKeeper.GetPrimaryName(suite.ctx, address.String())
	suite.Require().Equal(true, f)
	suite.Require().Equal(name, primName.GetDisplay())

	nameReq := types.QueryName{
		Name: name,
	}

	afterbal := suite.bankKeeper.GetAllBalances(suite.ctx, address)
	newamt := afterbal.AmountOf("ujkl")

	newamt = amt.Sub(newamt)
	leftover := 2 * cost
	suite.Require().Equal(leftover, newamt.Int64()) // cost them the price of the registration

	_, err = suite.queryClient.Name(suite.ctx.Context(), &nameReq)
	suite.Require().NoError(err)

	err = suite.rnsKeeper.RegisterRNSName(suite.ctx, address.String(), capname, "{}", 2, true) // adding time to registration
	suite.Require().NoError(err)

	afterbal = suite.bankKeeper.GetAllBalances(suite.ctx, address)
	newamt = afterbal.AmountOf("ujkl")
	leftover = cost * 4
	newamt = amt.Sub(newamt)
	suite.Require().Equal(leftover, newamt.Int64()) // cost them the price of the registration

	_, err = suite.queryClient.Name(suite.ctx.Context(), &nameReq)
	suite.Require().NoError(err)

	for i := 0; i < 100; i++ {
		err = suite.rnsKeeper.RegisterRNSName(suite.ctx, address.String(), fmt.Sprintf("mrpumpkinman%d.jkl", i), "{}", 1, false) // adding time to registration
		suite.Require().NoError(err)
	}

	r := types.QueryListOwnedNames{
		Address: address.String(),
	}

	res, err := suite.queryClient.ListOwnedNames(suite.ctx.Context(), &r)
	suite.Require().NoError(err)

	suite.Require().Equal(101, len(res.Names))

}
