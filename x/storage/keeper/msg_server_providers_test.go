package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

// testing msg server files for: ...

func (suite *KeeperTestSuite) TestMsgInitProvider() {
	suite.SetupSuite()

	msgSrvr, _, context := setupMsgServer(suite)

	user, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	cases := []struct {
		preRun    func() *types.MsgInitProvider
		expErr    bool
		expErrMsg string
		name      string
	}{
		{
			preRun: func() *types.MsgInitProvider {
				return types.NewMsgInitProvider(
					user.String(),
					"127.0.0.1",
					"1000000000",
				)
			},
			expErr: false,
			name:   "init provider success",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.InitProvider(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgInitProviderResponse{}, *res)

			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgSetProviderIP() {
	suite.SetupSuite()

	msgSrvr, _, context := setupMsgServer(suite)

	user, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	cases := []struct {
		preRun    func() *types.MsgSetProviderIP
		expErr    bool
		expErrMsg string
		name      string
	}{
		{
			preRun: func() *types.MsgSetProviderIP {
				return types.NewMsgSetProviderIP(
					user.String(),
					"127.0.0.1",
				)
			},
			expErr: false,
			name:   "set provider ip success",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.SetProviderIP(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgSetProviderIPResponse{}, *res)

			}
		})
	}
}
