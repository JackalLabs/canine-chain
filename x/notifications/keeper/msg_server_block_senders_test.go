package keeper_test

import (
	"github.com/jackalLabs/canine-chain/v4/testutil"
	"github.com/jackalLabs/canine-chain/v4/x/notifications/types"
)

func (suite *KeeperTestSuite) TestMsgBlockSenders() {
	suite.SetupSuite()
	msgSrvr, context := setupMsgServer(suite)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 3)
	suite.Require().NoError(err)

	alice := testAddresses[0]
	bob := testAddresses[1]
	charlie := testAddresses[2]

	cases := []struct {
		preRun    func() *types.MsgBlockSenders
		expErr    bool
		expErrMsg string
		name      string
	}{
		{
			preRun: func() *types.MsgBlockSenders {
				return types.NewMsgBlockSenders(
					bob,
					alice,
					charlie,
				)
			},
			expErr: false,
			name:   "bob blocks some senders",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			_, err := msgSrvr.BlockSenders(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().True(suite.notificationsKeeper.IsBlocked(suite.ctx, bob, alice))
				suite.Require().True(suite.notificationsKeeper.IsBlocked(suite.ctx, bob, charlie))
			}
		})
	}
}
