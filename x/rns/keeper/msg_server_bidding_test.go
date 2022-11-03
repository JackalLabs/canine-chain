package keeper_test

// msg server tests for all the bidding msg server files

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/rns/types"
)

func (suite *KeeperTestSuite) TestMsgAddBid() {
	suite.SetupSuite()
	msgSrvr, _, context := setupMsgServer(suite)

	nameOwner, err := sdk.AccAddressFromBech32("cosmos17j2hkm7n9fz9dpntyj2kxgxy5pthzd289nvlfl")
	suite.Require().NoError(err)

	rnsName := "Nuggie.jkl"

	suite.rnsKeeper.SetInit(suite.ctx, types.Init{Address: nameOwner.String(), Complete: true})
	err = suite.rnsKeeper.RegisterName(suite.ctx, nameOwner.String(), rnsName, "{}", "2")
	suite.Require().NoError(err)

	_, _ = msgSrvr.List(sdk.WrapSDKContext(suite.ctx), &types.MsgList{Creator: nameOwner.String(), Name: rnsName, Price: "200ujkl"})

	bidder, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	coin := sdk.NewCoin("ujkl", sdk.NewInt(100000))
	coins := sdk.NewCoins(coin)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, bidder, coins)
	suite.Require().NoError(err)

	cases := map[string]struct {
		preRun    func() *types.MsgBid
		expErr    bool
		expErrMsg string
	}{
		"bid successfully posted": {
			preRun: func() *types.MsgBid {
				return types.NewMsgBid(
					bidder.String(),
					"Nuggie.jkl",
					"1000ujkl",
				)
			},
			expErr: false,
		},
		"you don't have enough money": {
			preRun: func() *types.MsgBid {
				return types.NewMsgBid(
					bidder.String(),
					"Nuggie.jkl",
					"1000000ujkl",
				)
			},
			expErr:    true,
			expErrMsg: "not enough balance",
		},
	}

	for name, tc := range cases {
		suite.Run(name, func() {
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

	cases := map[string]struct {
		preRun    func() *types.MsgAcceptBid
		expErr    bool
		expErrMsg string
	}{
		"bid successfully accepted": {
			preRun: func() *types.MsgAcceptBid {
				return &types.MsgAcceptBid{
					Creator: nameOwner.String(),
					Name:    TestName,
					From:    bidder.String(),
				}
			},
			expErr: false,
		},
		"bid failed to be accepted": {
			preRun: func() *types.MsgAcceptBid {
				return &types.MsgAcceptBid{
					Creator: nameOwner.String(),
					Name:    TestName,
					From:    bidder.String(),
				}
			},
			expErr: true,
		},
	}

	for name, tc := range cases {
		suite.Run(name, func() {
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

// Cancel bid goes here
