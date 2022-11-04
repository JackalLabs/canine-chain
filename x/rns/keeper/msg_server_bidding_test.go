package keeper_test

// testing msg server files for: bid, accept_bid, cancel_bid

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/rns/types"
)

const nuggieName = "Nuggie.jkl"

func (suite *KeeperTestSuite) TestMsgAddBid() {
	suite.SetupSuite()
	msgSrvr, _, context := setupMsgServer(suite)

	nameOwner, err := sdk.AccAddressFromBech32("cosmos17j2hkm7n9fz9dpntyj2kxgxy5pthzd289nvlfl")
	suite.Require().NoError(err)

	coin := sdk.NewCoin("ujkl", sdk.NewInt(100000000))
	coins := sdk.NewCoins(coin)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, nameOwner, coins)
	suite.Require().NoError(err)

	suite.rnsKeeper.SetInit(suite.ctx, types.Init{Address: nameOwner.String(), Complete: true})
	err = suite.rnsKeeper.RegisterName(suite.ctx, nameOwner.String(), nuggieName, "{}", "2")
	suite.Require().NoError(err)

	_, _ = msgSrvr.List(sdk.WrapSDKContext(suite.ctx), &types.MsgList{Creator: nameOwner.String(), Name: nuggieName, Price: "200ujkl"})

	bidder, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

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
					"1000ujkl",
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
					"100000000ujkl",
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

	nameOwner, err := sdk.AccAddressFromBech32("cosmos17j2hkm7n9fz9dpntyj2kxgxy5pthzd289nvlfl")
	suite.Require().NoError(err)

	bidder, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	coin := sdk.NewCoin("ujkl", sdk.NewInt(100000000))
	coins := sdk.NewCoins(coin)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, nameOwner, coins)
	suite.Require().NoError(err)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, bidder, coins)
	suite.Require().NoError(err)

	err = suite.rnsKeeper.RegisterName(suite.ctx, nameOwner.String(), TestName, "{}", "2")
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

	nameOwner, err := sdk.AccAddressFromBech32("cosmos17j2hkm7n9fz9dpntyj2kxgxy5pthzd289nvlfl")
	suite.Require().NoError(err)

	bidder, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	coin := sdk.NewCoin("ujkl", sdk.NewInt(100000000))
	coins := sdk.NewCoins(coin)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, nameOwner, coins)
	suite.Require().NoError(err)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, bidder, coins)
	suite.Require().NoError(err)

	err = suite.rnsKeeper.RegisterName(suite.ctx, nameOwner.String(), TestName, "{}", "2")
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
