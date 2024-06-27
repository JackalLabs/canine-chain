package keeper_test

import (
	"github.com/jackalLabs/canine-chain/v4/testutil"
	"github.com/jackalLabs/canine-chain/v4/x/notifications/types"
)

func (suite *KeeperTestSuite) TestMsgCreateNotifications() {
	suite.SetupSuite()
	msgSrvr, context := setupMsgServer(suite)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 3)
	suite.Require().NoError(err)

	alice := testAddresses[0]
	bob := testAddresses[1]
	charlie := testAddresses[2]

	suite.notificationsKeeper.SetBlock(suite.ctx, types.Block{
		Address:        bob,
		BlockedAddress: charlie,
	})

	cases := []struct {
		preRun    func() *types.MsgCreateNotification
		expErr    bool
		expErrMsg string
		name      string
	}{
		{
			preRun: func() *types.MsgCreateNotification {
				return types.NewMsgCreateNotification(
					alice,
					bob,
					"{\"msg\":\"hey bob it's alice here\"}",
					make([]byte, 0),
				)
			},
			expErr: false,
			name:   "alice successfully sends a notification to bob",
		},
		{
			preRun: func() *types.MsgCreateNotification {
				return types.NewMsgCreateNotification(
					alice,
					bob,
					"hey bob it's alice here",
					make([]byte, 0),
				)
			},
			expErr:    true,
			name:      "cannot parse json",
			expErrMsg: "contents are not valid `hey bob it's alice here`: failed to unmarshal JSON bytes",
		},
		{
			preRun: func() *types.MsgCreateNotification {
				return types.NewMsgCreateNotification(
					charlie,
					bob,
					"{\"msg\": \"hey bob it's charlie here\"}",
					make([]byte, 0),
				)
			},
			expErr:    true,
			name:      "charlie is blocked",
			expErrMsg: "you are blocked from sending this user notifications: unauthorized",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			_, err = msgSrvr.CreateNotification(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}
