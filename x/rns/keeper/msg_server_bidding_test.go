package keeper_test

// testing msg server files for: bid, accept_bid, cancel_bid

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/testutil"
	"github.com/jackalLabs/canine-chain/v4/x/rns/types"
)

const nuggieName = "Nuggie.jkl"

func (suite *KeeperTestSuite) TestMsgAddBid() {
	suite.SetupSuite()
	msgSrvr, _, context := setupMsgServer(suite)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	suite.Require().NoError(err)

	nameOwner, err := sdk.AccAddressFromBech32(testAddresses[0])
	suite.Require().NoError(err)

	bidder, err := sdk.AccAddressFromBech32(testAddresses[1])
	suite.Require().NoError(err)

	coin := sdk.NewCoin("ujkl", sdk.NewInt(100000000))
	coins := sdk.NewCoins(coin)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, nameOwner, coins)
	suite.Require().NoError(err)

	suite.rnsKeeper.SetInit(suite.ctx, types.Init{Address: nameOwner.String(), Complete: true})
	err = suite.rnsKeeper.RegisterName(suite.ctx, nameOwner.String(), nuggieName, "{}", 2)
	suite.Require().NoError(err)

	_, _ = msgSrvr.List(sdk.WrapSDKContext(suite.ctx), &types.MsgList{Creator: nameOwner.String(), Name: nuggieName, Price: sdk.NewInt64Coin("ujkl", 200)})

	coin = sdk.NewCoin("ujkl", sdk.NewInt(10000))
	coins = sdk.NewCoins(coin)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, bidder, coins)
	suite.Require().NoError(err)

	cases := []struct {
		preRun    func() *types.MsgBid
		expErr    bool
		expErrMsg string
		name      string
	}{
		{
			preRun: func() *types.MsgBid {
				return types.NewMsgBid(
					bidder.String(),
					nuggieName,
					sdk.NewInt64Coin("ujkl", 1000),
				)
			},
			expErr: false,
			name:   "bid successfully posted",
		},
		{
			preRun: func() *types.MsgBid {
				return types.NewMsgBid(
					bidder.String(),
					nuggieName,
					sdk.NewInt64Coin("ujkl", 100000000),
				)
			},
			expErr:    true,
			expErrMsg: "not enough balance",
			name:      "you don't have enough money",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.Bid(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgBidResponse{}, *res)

			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgAcceptOneBid() {
	suite.SetupSuite()
	msgSrvr, _, context := setupMsgServer(suite)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	suite.Require().NoError(err)

	nameOwner, err := sdk.AccAddressFromBech32(testAddresses[0])
	suite.Require().NoError(err)

	bidder, err := sdk.AccAddressFromBech32(testAddresses[1])
	suite.Require().NoError(err)

	coin := sdk.NewCoin("ujkl", sdk.NewInt(100000000))
	coins := sdk.NewCoins(coin)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, nameOwner, coins)
	suite.Require().NoError(err)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, bidder, coins)
	suite.Require().NoError(err)

	err = suite.rnsKeeper.RegisterName(suite.ctx, nameOwner.String(), TestName, "{}", 2)
	suite.Require().NoError(err)

	err = suite.rnsKeeper.AddBid(suite.ctx, bidder.String(), TestName, "1000ujkl")
	suite.Require().NoError(err)

	cases := []struct {
		preRun    func() *types.MsgAcceptBid
		expErr    bool
		expErrMsg string
		name      string
	}{
		{
			preRun: func() *types.MsgAcceptBid {
				return &types.MsgAcceptBid{
					Creator: nameOwner.String(),
					Name:    TestName,
					From:    bidder.String(),
				}
			},
			expErr: false,
			name:   "big successfully accepted",
		},
		{
			preRun: func() *types.MsgAcceptBid {
				return &types.MsgAcceptBid{
					Creator: nameOwner.String(),
					Name:    TestName,
					From:    bidder.String(),
				}
			},
			expErr:    true,
			expErrMsg: "You are not the owner of that name.: unauthorized",
			name:      "bid failed to be accepted",
		},
		{
			preRun: func() *types.MsgAcceptBid {
				freeName := "freeBi.jkl"
				blockHeight := suite.ctx.BlockHeight()
				err := suite.rnsKeeper.RegisterName(suite.ctx, nameOwner.String(), freeName, "{}", 2)
				suite.Require().NoError(err)
				name, _ := suite.rnsKeeper.GetNames(suite.ctx, "freeBi", "jkl")
				name.Locked = blockHeight + 1

				suite.rnsKeeper.SetNames(suite.ctx, name)
				err1 := suite.rnsKeeper.AddBid(suite.ctx, bidder.String(), freeName, "2000ujkl")
				suite.Require().NoError(err1)

				return types.NewMsgAcceptBid(
					nameOwner.String(),
					"freeBi.jkl",
					bidder.String(),
				)
			},
			expErr:    true,
			expErrMsg: "cannot transfer free name: unauthorized",
			name:      "cannot transfer free name",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.AcceptBid(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgBidResponse{}, *res)

			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgCancelOneBid() {
	suite.SetupSuite()
	msgSrvr, _, context := setupMsgServer(suite)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	suite.Require().NoError(err)

	nameOwner, err := sdk.AccAddressFromBech32(testAddresses[0])
	suite.Require().NoError(err)

	bidder, err := sdk.AccAddressFromBech32(testAddresses[1])
	suite.Require().NoError(err)

	coin := sdk.NewCoin("ujkl", sdk.NewInt(100000000))
	coins := sdk.NewCoins(coin)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, nameOwner, coins)
	suite.Require().NoError(err)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, bidder, coins)
	suite.Require().NoError(err)

	err = suite.rnsKeeper.RegisterName(suite.ctx, nameOwner.String(), TestName, "{}", 2)
	suite.Require().NoError(err)

	err = suite.rnsKeeper.AddBid(suite.ctx, bidder.String(), TestName, "1000ujkl")
	suite.Require().NoError(err)

	cases := []struct {
		preRun    func() *types.MsgCancelBid
		expErr    bool
		expErrMsg string
		name      string
	}{
		{
			preRun: func() *types.MsgCancelBid {
				return &types.MsgCancelBid{
					Creator: bidder.String(),
					Name:    TestName,
				}
			},
			expErr: false,
			name:   "bid successfully canceled",
		},
		{
			preRun: func() *types.MsgCancelBid {
				return &types.MsgCancelBid{
					Creator: bidder.String(),
					Name:    TestName,
				}
			},
			expErr:    true,
			expErrMsg: "Bid does not exist or has expired.: not found",
			name:      "bid unsuccessfully canceled",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.CancelBid(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgBidResponse{}, *res)

			}
		})
	}
}
