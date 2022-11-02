package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	testutil "github.com/jackalLabs/canine-chain/testutil"
	"github.com/jackalLabs/canine-chain/x/rns/types"
)

// testing the msg server
func (suite *KeeperTestSuite) TestMsgInit() {
	logger, logFile := testutil.CreateLogger()
	ctx := sdk.WrapSDKContext(suite.ctx)
	suite.reset()
	//We could use simulate address library, but veeering away from being dependent on the cosmosSDK
	user, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)
	logger.Println("user is", user)
	logger.Println("user as string is", user.String())

	cases := map[string]struct {
		preRun    func() (*types.MsgInit, error)
		expErr    bool
		expErrMsg string
	}{
		"all good": {
			preRun: func() (*types.MsgInit, error) {
				return types.NewMsgInit(
					user.String(),
				)
			},
			expErr: false,
		},
		"cannot init twice": {
			preRun: func() (*types.MsgInit, error) {
				return types.NewMsgInit(
					user.String(),
				)
			},
			expErr:    true,
			expErrMsg: "not permitted to init twice",
		},
	}

	for name, tc := range cases {
		suite.Run(name, func() {
			msg, err := tc.preRun()
			suite.Require().NoError(err)
			res, err := suite.msgSrvr.Init(ctx, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().Nil(res)
			}
		})
	}
	logFile.Close()

}

/*
	suite.SetupSuite()
	err := suite.setupNames()

	suite.Require().NoError(err)


	initMsg := types.NewMsgInit(address.String())

	_, err1 := suite.msgSrvr.Init(sdk.WrapSDKContext(suite.ctx), initMsg)
	suite.Require().NoError(err1)

	initReq := types.QueryGetInitRequest{
		Address: address.String(),
	}

	_, err2 := suite.queryClient.Init(suite.ctx.Context(), &initReq)
	suite.Require().NoError(err2)

	//init again should fail
	_, err3 := suite.msgSrvr.Init(sdk.WrapSDKContext(suite.ctx), initMsg)
	suite.Require().Error(err3)




*/
