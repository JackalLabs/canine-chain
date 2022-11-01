package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/rns/types"
)

func (suite *KeeperTestSuite) TestMsgAcceptBid() {
	suite.SetupSuite()
	err := suite.setupNames()
	suite.Require().NoError(err)
	address, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	nameAddress, err := sdk.AccAddressFromBech32("cosmos17j2hkm7n9fz9dpntyj2kxgxy5pthzd289nvlfl")
	suite.Require().NoError(err)

	name := "test.jkl"

	coin := sdk.NewCoin("ujkl", sdk.NewInt(100000000))
	coins := sdk.NewCoins(coin)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, address, coins)
	suite.Require().NoError(err)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, nameAddress, coins)
	suite.Require().NoError(err)

	err = suite.rnsKeeper.RegisterName(suite.ctx, nameAddress.String(), name, "{}", "2")
	suite.Require().NoError(err)

	bidderBalBefore := suite.bankKeeper.GetAllBalances(suite.ctx, address)
	biddeeBalBefore := suite.bankKeeper.GetAllBalances(suite.ctx, nameAddress)

	err = suite.rnsKeeper.AddBid(suite.ctx, address.String(), name, "1000ujkl")
	suite.Require().NoError(err)

	err = suite.rnsKeeper.AcceptOneBid(suite.ctx, nameAddress.String(), name, address.String())
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
	suite.Require().Equal(ramt.Int64(), bidder) //cost them the amount they bid
	suite.Require().Equal(eamt.Int64(), biddee) //cost them the amount they bid

	nameReq := types.QueryGetNamesRequest{
		Index: name,
	}

	res, err := suite.queryClient.Names(suite.ctx.Context(), &nameReq)
	suite.Require().NoError(err)
	suite.Require().Equal(res.Names.Value, address.String())
}

func (suite *KeeperTestSuite) TestMsgMakeBid() {
	suite.SetupSuite()
	err := suite.setupNames()
	suite.Require().NoError(err)
	address, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)
	name := "test.jkl"

	coin := sdk.NewCoin("ujkl", sdk.NewInt(100000))
	coins := sdk.NewCoins(coin)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, address, coins)
	suite.Require().NoError(err)

	beforebal := suite.bankKeeper.GetAllBalances(suite.ctx, address)
	amt := beforebal.AmountOf("ujkl")

	err = suite.rnsKeeper.AddBid(suite.ctx, address.String(), name, "1000ujkl")
	suite.Require().NoError(err)

	bidReq := types.QueryGetBidsRequest{
		Index: fmt.Sprintf("%s%s", address.String(), name),
	}

	afterbal := suite.bankKeeper.GetAllBalances(suite.ctx, address)
	newamt := afterbal.AmountOf("ujkl")

	newamt = amt.Sub(newamt)
	var leftover int64 = 1000
	suite.Require().Equal(newamt.Int64(), leftover) //cost them the amount they bid

	_, err = suite.queryClient.Bids(suite.ctx.Context(), &bidReq)
	suite.Require().NoError(err)

}
