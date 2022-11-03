package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/rns/types"
)

// testing msg server files for: init, register, and transfer

func (suite *KeeperTestSuite) TestMsgInit() {
	suite.SetupSuite()

	msgSrvr, _, context := setupMsgServer(suite)

	user, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	cases := map[string]struct {
		preRun    func() *types.MsgInit
		expErr    bool
		expErrMsg string
	}{
		"successful init": {
			preRun: func() *types.MsgInit {
				return types.NewMsgInit(
					user.String(),
				)
			},
			expErr: false,
		},
		"cannot init twice": {
			preRun: func() *types.MsgInit {
				return types.NewMsgInit(
					user.String(),
				)
			},
			expErr:    true,
			expErrMsg: "cannot initialize more than once: invalid request",
		},
	}

	for name, tc := range cases {
		suite.Run(name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.Init(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgInitResponse{}, *res)

			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgRegister() {
	suite.SetupSuite()

	msgSrvr, _, context := setupMsgServer(suite)

	user, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)
	//suite.rnsKeeper.SetInit(suite.ctx, types.Init{Address: user.String(), Complete: true})

	coin := sdk.NewCoin("ujkl", sdk.NewInt(100000000))
	coins := sdk.NewCoins(coin)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, user, coins)
	suite.Require().NoError(err)

	cases := map[string]struct {
		preRun    func() *types.MsgRegister
		expErr    bool
		expErrMsg string
	}{
		"successful register": {
			preRun: func() *types.MsgRegister {
				return types.NewMsgRegister(
					user.String(),
					"BiPhan.jkl",
					"2",
					"{}",
				)
			},
			expErr: false,
		},
		"invalid address": {
			preRun: func() *types.MsgRegister {
				return types.NewMsgRegister(
					"invalid address",
					"BiPhan.jkl",
					"2",
					"{}",
				)
			},
			expErr:    true,
			expErrMsg: "cannot parse sender: decoding bech32 failed: invalid character in string:",
		},
		"invalid name": {
			preRun: func() *types.MsgRegister {
				return types.NewMsgRegister(
					user.String(),
					"BiPhan.LUNC",
					"2",
					"{}",
				)
			},
			expErr:    true,
			expErrMsg: "could not extract the tld from the name provided",
		},
		"invalid years": {
			preRun: func() *types.MsgRegister {
				return types.NewMsgRegister(
					user.String(),
					"BiPhan.jkl",
					"s",
					"{}",
				)
			},
			expErr:    true,
			expErrMsg: "cannot parse years: invalid height",
		},
	}

	for name, tc := range cases {
		suite.Run(name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.Register(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgInitResponse{}, *res)

			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgTrasnfer() {
	suite.SetupSuite()

	msgSrvr, _, context := setupMsgServer(suite)

	owner, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)
	receiver, err := sdk.AccAddressFromBech32("cosmos1xetrp5dwjplsn4lev5r2cu8en5qsq824vza9nu")
	suite.Require().NoError(err)

	coin := sdk.NewCoin("ujkl", sdk.NewInt(100000000))
	coins := sdk.NewCoins(coin)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, owner, coins)
	suite.Require().NoError(err)

	cases := map[string]struct {
		preRun    func() *types.MsgTransfer
		expErr    bool
		expErrMsg string
	}{
		"successful transfer": {
			preRun: func() *types.MsgTransfer {
				return types.NewMsgTransfer(
					owner.String(),
					"BiPhan.jkl",
					receiver.String(),
				)
			},
			expErr: false,
		},
	}

	for name, tc := range cases {
		suite.Run(name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.Transfer(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgInitResponse{}, *res)

			}
		})
	}
}

/*
	receiver, err := sdk.AccAddressFromBech32("cosmos1xetrp5dwjplsn4lev5r2cu8en5qsq824vza9nu")
	suite.Require().NoError(err)
*/
