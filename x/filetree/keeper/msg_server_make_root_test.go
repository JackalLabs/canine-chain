package keeper_test

import (
	"github.com/jackalLabs/canine-chain/v3/testutil"
	"github.com/jackalLabs/canine-chain/v3/x/filetree/types"
)

func (suite *KeeperTestSuite) TestMsgMakeRoot() {
	suite.SetupSuite()
	msgSrvr, context := setupMsgServer(suite)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 1)
	suite.Require().NoError(err)

	alice := testAddresses[0]

	msg, err := types.CreateMsgMakeRoot(alice)
	suite.Require().NoError(err)

	cases := []struct {
		preRun    func() *types.MsgMakeRootV2
		expErr    bool
		expErrMsg string
		name      string
	}{
		{
			preRun: func() *types.MsgMakeRootV2 {
				return msg
			},
			expErr: false,
			name:   "make root success",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.MakeRootV2(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgMakeRootResponse{}, *res)

			}
		})
	}
}
