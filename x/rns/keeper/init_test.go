package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/rns/types"
)

// Combine testing of init.go and msg_server_init.go

// testing the msg server
func (suite *KeeperTestSuite) TestMsgInit() {
	suite.SetupSuite()

	msgSrvr, _, context := setupMsgServer(suite)

	// Need to add mock random addresses
	user, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	cases := map[string]struct {
		preRun    func() *types.MsgInit
		expErr    bool
		expErrMsg string
	}{
		"all good": {
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
