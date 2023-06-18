package keeper_test

import (
	"github.com/jackalLabs/canine-chain/testutil"
	"github.com/jackalLabs/canine-chain/x/notifications/types"
)

func (suite *KeeperTestSuite) TestMsgSetCounter() {
	suite.SetupSuite()
	msgSrvr, context := setupMsgServer(suite)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	suite.Require().NoError(err)

	alice := testAddresses[0]

	cases := []struct {
		preRun    func() *types.MsgSetCounter
		expErr    bool
		expErrMsg string
		name      string
	}{
		{ // alice sets her notiCounter
			preRun: func() *types.MsgSetCounter {
				return types.NewMsgSetCounter(
					alice,
				)
			},
			expErr: false,
			name:   "alice successfully sets a notiCounter for herself",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.SetCounter(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {

				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgSetCounterResponse{}, *res)

			}
		})
	}
}
