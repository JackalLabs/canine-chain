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

	cases := []struct {
		preRun    func() *types.MsgInit
		expErr    bool
		expErrMsg string
		name      string
	}{
		{
			preRun: func() *types.MsgInit {
				return types.NewMsgInit(
					user.String(),
				)
			},
			expErr: false,
			name:   "init success",
		},
		{
			preRun: func() *types.MsgInit {
				return types.NewMsgInit(
					user.String(),
				)
			},
			expErr:    true,
			expErrMsg: "cannot initialize more than once: invalid request",
			name:      "cannot init twice",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
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

	coin := sdk.NewCoin("ujkl", sdk.NewInt(100000000))
	coins := sdk.NewCoins(coin)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, user, coins)
	suite.Require().NoError(err)

	cases := []struct {
		preRun    func() *types.MsgRegister
		expErr    bool
		expErrMsg string
		name      string
	}{
		{
			preRun: func() *types.MsgRegister {
				return types.NewMsgRegister(
					user.String(),
					"BiPhan.jkl",
					"2",
					"{}",
				)
			},
			expErr: false,
			name:   "successful register",
		},
		{
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
			name:      "invalid address",
		},
		{
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
			name:      "invalid name",
		},
		{
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
			name:      "invalid years",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.Register(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgRegisterResponse{}, *res)

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

	successfulName := "BiPhan.jkl"

	coin := sdk.NewCoin("ujkl", sdk.NewInt(100000000))
	coins := sdk.NewCoins(coin)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, owner, coins)
	suite.Require().NoError(err)

	err = suite.rnsKeeper.RegisterName(suite.ctx, owner.String(), successfulName, "{}", "2")
	suite.Require().NoError(err)

	cases := []struct {
		preRun    func() *types.MsgTransfer
		expErr    bool
		expErrMsg string
		name      string
	}{
		{
			preRun: func() *types.MsgTransfer {
				return types.NewMsgTransfer(
					owner.String(),
					successfulName,
					receiver.String(),
				)
			},
			expErr: false,
			name:   "successful transfer",
		},
		{
			preRun: func() *types.MsgTransfer {
				return types.NewMsgTransfer(
					owner.String(),
					successfulName,
					receiver.String(),
				)
			},
			expErr:    true,
			expErrMsg: "You are not the owner of that name.: unauthorized",
			name:      "failed transfer",
		},
		{
			preRun: func() *types.MsgTransfer {
				return types.NewMsgTransfer(
					owner.String(),
					"nonExistentName.jkl",
					receiver.String(),
				)
			},
			expErr:    true,
			expErrMsg: "Name does not exist or has expired.: not found",
			name:      "cannot transfer name that doesn't exist",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.Transfer(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgRegisterResponse{}, *res)

			}
		})
	}
}
