package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/testutil"
	types "github.com/jackalLabs/canine-chain/v4/x/rns/types"
)

func (suite *KeeperTestSuite) TestBuyMsg() {
	suite.SetupSuite()
	msgSrvr, ctx := setupMsgServer(suite)
	keeper := suite.rnsKeeper

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 3)
	suite.Require().NoError(err)

	nameOwner, err := sdk.AccAddressFromBech32(testAddresses[0])
	suite.Require().NoError(err)

	buyer, err := sdk.AccAddressFromBech32(testAddresses[1])
	suite.Require().NoError(err)

	brokeBuyer, err := sdk.AccAddressFromBech32(testAddresses[2])
	suite.Require().NoError(err)

	coins := sdk.NewCoins(sdk.NewCoin("ujkl", sdk.NewInt(10000000000)))
	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, nameOwner, coins)
	suite.Require().NoError(err)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, buyer, coins)
	suite.Require().NoError(err)

	// Init rns account and register rns
	rnsName := "Nuggie"
	rnsTLD := "jkl"
	fullName := rnsName + "." + rnsTLD // "Nuggie.jkl"

	keeper.SetInit(suite.ctx, types.Init{Address: nameOwner.String(), Complete: true})
	err = suite.rnsKeeper.RegisterRNSName(suite.ctx, nameOwner.String(), fullName, "{}", 2, true)
	suite.Require().NoError(err)
	originalNames, found := keeper.GetNames(suite.ctx, rnsName, rnsTLD)
	suite.Require().True(found)

	// Put it up for sale
	salePrice := sdk.NewInt64Coin("ujkl", 1000000)
	msgList := types.MsgList{
		Creator: nameOwner.String(),
		Name:    fullName,
		Price:   salePrice,
	}
	_, err = msgSrvr.List(ctx, &msgList)
	suite.Require().NoError(err)

	// Test cases
	cases := []struct {
		testName  string
		preRun    func() *types.MsgBuy
		postRun   func()
		expErr    bool
		expErrMsg string
	}{
		{
			testName: "name_not_for_sale",
			preRun: func() *types.MsgBuy {
				return &types.MsgBuy{
					Name:    "not_for_sale.jkl",
					Creator: buyer.String(),
				}
			},
			expErr:    true,
			expErrMsg: "Name not for sale.",
		},
		{
			testName: "name_not_found",
			preRun: func() *types.MsgBuy {
				// Delete name from KVStore
				names, found := keeper.GetNames(suite.ctx, rnsName, rnsTLD)
				if !found {
					return nil
				}
				keeper.RemoveNames(suite.ctx, names.Name, names.Tld)
				_, found = keeper.GetNames(suite.ctx, rnsName, rnsTLD)
				suite.Require().False(found)
				return &types.MsgBuy{
					Name:    fullName,
					Creator: buyer.String(),
				}
			},
			postRun: func() {
				// add back name
				keeper.SetNames(suite.ctx, originalNames)
				_, found = keeper.GetNames(suite.ctx, rnsName, rnsTLD)
				suite.Require().True(found)
			},
			expErr:    true,
			expErrMsg: "Name does not exist",
		},

		{
			testName: "name_expired",
			preRun: func() *types.MsgBuy {
				// Make the name expired
				names, found := keeper.GetNames(suite.ctx, rnsName, rnsTLD)
				suite.Require().True(found)
				names.Expires = suite.ctx.BlockHeight() - 1
				keeper.SetNames(suite.ctx, names)
				return &types.MsgBuy{
					Name:    fullName,
					Creator: buyer.String(),
				}
			},
			postRun: func() {
				names, found := keeper.GetNames(suite.ctx, rnsName, rnsTLD)
				suite.Require().True(found)
				names.Expires++
				keeper.SetNames(suite.ctx, names)
			},
			expErr:    true,
			expErrMsg: "expired",
		},

		{
			testName: "owner_cannot_buy_owners_name",
			preRun: func() *types.MsgBuy {
				return &types.MsgBuy{
					Creator: nameOwner.String(),
					Name:    fullName,
				}
			},
			expErr:    true,
			expErrMsg: "cannot buy your own name",
		},
		{
			testName: "not_enough_balance",
			preRun: func() *types.MsgBuy {
				return &types.MsgBuy{
					Creator: brokeBuyer.String(),
					Name:    fullName,
				}
			},
			expErr:    true,
			expErrMsg: "not enough balance",
		},
		{
			testName: "successful_sale",
			preRun: func() *types.MsgBuy {
				return &types.MsgBuy{
					Name:    "Nuggie.jkl",
					Creator: buyer.String(),
				}
			},
			expErr: false,
		},
	}

	for _, tc := range cases {
		suite.Run(tc.testName, func() {
			msg := tc.preRun()

			err := keeper.BuyName(suite.ctx, msg.Creator, msg.Name)
			if tc.expErr {
				suite.Require().ErrorContains(err, tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
			}

			if tc.postRun != nil {
				tc.postRun()
			}
		})
	}
}
