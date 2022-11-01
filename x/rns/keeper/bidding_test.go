package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/rns/types"
)

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
