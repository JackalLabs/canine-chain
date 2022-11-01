package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	testutil "github.com/jackalLabs/canine-chain/testutil"
	"github.com/jackalLabs/canine-chain/x/rns/types"
)

func (suite *KeeperTestSuite) TestMsgRegisterName() {

	logger, logFile := testutil.CreateLogger()

	suite.SetupSuite()

	logger.Println("The suite is:", suite)

	err := suite.setupNames()

	suite.Require().NoError(err)
	address, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)
	name := "test.jkl"

	coin := sdk.NewCoin("ujkl", sdk.NewInt(100000000))
	coins := sdk.NewCoins(coin)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, address, coins)
	suite.Require().NoError(err)

	beforebal := suite.bankKeeper.GetAllBalances(suite.ctx, address)
	amt := beforebal.AmountOf("ujkl")

	err = suite.rnsKeeper.RegisterName(suite.ctx, address.String(), name, "{}", "2")
	suite.Require().NoError(err)

	nameReq := types.QueryGetNamesRequest{
		Index: name,
	}

	afterbal := suite.bankKeeper.GetAllBalances(suite.ctx, address)
	newamt := afterbal.AmountOf("ujkl")

	newamt = amt.Sub(newamt)
	var leftover int64 = 15000000
	suite.Require().Equal(newamt.Int64(), leftover) // cost them the amount they bid

	_, err = suite.queryClient.Names(suite.ctx.Context(), &nameReq)
	suite.Require().NoError(err)
	logFile.Close()
}
