package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/testutil"
	"github.com/jackalLabs/canine-chain/v3/x/rns/types"
)

const TestName = "test.jkl"

func (suite *KeeperTestSuite) TestMsgAcceptBid() {
	suite.SetupSuite()
	suite.setupNames()

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	suite.Require().NoError(err)

	addr := testAddresses[0]
	nameAddr := testAddresses[1]

	address, err := sdk.AccAddressFromBech32(addr)
	suite.Require().NoError(err)

	nameAddress, err := sdk.AccAddressFromBech32(nameAddr)
	suite.Require().NoError(err)

	coin := sdk.NewCoin("ujkl", sdk.NewInt(100000000))
	coins := sdk.NewCoins(coin)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, address, coins)
	suite.Require().NoError(err)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, nameAddress, coins)
	suite.Require().NoError(err)

	err = suite.rnsKeeper.RegisterName(suite.ctx, nameAddress.String(), TestName, "{}", 2)
	suite.Require().NoError(err)

	bidderBalBefore := suite.bankKeeper.GetAllBalances(suite.ctx, address)
	biddeeBalBefore := suite.bankKeeper.GetAllBalances(suite.ctx, nameAddress)

	err = suite.rnsKeeper.AddBid(suite.ctx, address.String(), TestName, "1000ujkl")
	suite.Require().NoError(err)

	err = suite.rnsKeeper.AcceptOneBid(suite.ctx, nameAddress.String(), TestName, address.String())
	suite.Require().NoError(err)

	bidderBalAfter := suite.bankKeeper.GetAllBalances(suite.ctx, address)
	biddeeBalAfter := suite.bankKeeper.GetAllBalances(suite.ctx, nameAddress)

	rbb := bidderBalBefore.AmountOf("ujkl")
	ebb := biddeeBalBefore.AmountOf("ujkl")

	rba := bidderBalAfter.AmountOf("ujkl")
	eba := biddeeBalAfter.AmountOf("ujkl")

	ramt := rbb.Sub(rba)
	eamt := ebb.Sub(eba)

	var bidder int64 = 1000
	var biddee int64 = -1000
	suite.Require().Equal(ramt.Int64(), bidder) // cost them the amount they bid
	suite.Require().Equal(eamt.Int64(), biddee) // cost them the amount they bid

	nameReq := types.QueryNameRequest{
		Index: TestName,
	}

	res, err := suite.queryClient.Names(suite.ctx.Context(), &nameReq)
	suite.Require().NoError(err)
	suite.Require().Equal(res.Names.Value, address.String())
}

func (suite *KeeperTestSuite) TestMsgMakeBid() {
	suite.SetupSuite()
	suite.setupNames()

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 1)
	suite.Require().NoError(err)

	addr := testAddresses[0]
	address, err := sdk.AccAddressFromBech32(addr)
	suite.Require().NoError(err)

	coin := sdk.NewCoin("ujkl", sdk.NewInt(100000))
	coins := sdk.NewCoins(coin)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, address, coins)
	suite.Require().NoError(err)

	beforebal := suite.bankKeeper.GetAllBalances(suite.ctx, address)
	amt := beforebal.AmountOf("ujkl")

	err = suite.rnsKeeper.AddBid(suite.ctx, address.String(), TestName, "1000ujkl")
	suite.Require().NoError(err)

	bidReq := types.QueryBidRequest{
		Index: fmt.Sprintf("%s%s", address.String(), TestName),
	}

	afterbal := suite.bankKeeper.GetAllBalances(suite.ctx, address)
	newamt := afterbal.AmountOf("ujkl")

	newamt = amt.Sub(newamt)
	var leftover int64 = 1000
	suite.Require().Equal(newamt.Int64(), leftover) // cost them the amount they bid

	_, err = suite.queryClient.Bids(suite.ctx.Context(), &bidReq)
	suite.Require().NoError(err)
}

func (suite *KeeperTestSuite) TestMsgCancelBid() {
	suite.SetupSuite()
	suite.setupNames()

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 1)
	suite.Require().NoError(err)

	addr := testAddresses[0]
	address, err := sdk.AccAddressFromBech32(addr)
	suite.Require().NoError(err)

	coin := sdk.NewCoin("ujkl", sdk.NewInt(100000))
	coins := sdk.NewCoins(coin)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, address, coins)
	suite.Require().NoError(err)

	beforebal := suite.bankKeeper.GetAllBalances(suite.ctx, address)
	amt := beforebal.AmountOf("ujkl")

	err = suite.rnsKeeper.AddBid(suite.ctx, address.String(), TestName, "1000ujkl")
	suite.Require().NoError(err)

	bidReq := types.QueryBidRequest{
		Index: fmt.Sprintf("%s%s", address.String(), TestName),
	}

	afterbal := suite.bankKeeper.GetAllBalances(suite.ctx, address)
	newamt := afterbal.AmountOf("ujkl")

	newamt = amt.Sub(newamt)
	var leftover int64 = 1000                       // they spent 1000ujkl so they should have 1000ujkl less
	suite.Require().Equal(newamt.Int64(), leftover) // cost them the amount they bid

	_, err = suite.queryClient.Bids(suite.ctx.Context(), &bidReq)
	suite.Require().NoError(err)

	err = suite.rnsKeeper.CancelOneBid(suite.ctx, address.String(), TestName)
	suite.Require().NoError(err)

	afterbal = suite.bankKeeper.GetAllBalances(suite.ctx, address)
	newamt = afterbal.AmountOf("ujkl")

	newamt = amt.Sub(newamt)
	leftover = 0                                    // they cancelled the bid and thus should receive their money back
	suite.Require().Equal(newamt.Int64(), leftover) // cost them the amount they bid

	_, err = suite.queryClient.Bids(suite.ctx.Context(), &bidReq)
	suite.Require().Error(err)
}
