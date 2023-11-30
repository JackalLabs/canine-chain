package keeper_test

import (
	"time"

	"github.com/jackalLabs/canine-chain/v3/testutil"
	"github.com/jackalLabs/canine-chain/v3/x/notifications/types"
)

func (suite *KeeperTestSuite) TestMsgDeleteNotifications() {
	suite.SetupSuite()
	msgSrvr, context := setupMsgServer(suite)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	suite.Require().NoError(err)

	alice := testAddresses[0]
	bob := testAddresses[1]

	t := time.Now()

	notification := types.Notification{
		To:       bob,
		From:     alice,
		Time:     t,
		Contents: "{}",
	}

	suite.notificationsKeeper.SetNotification(suite.ctx, notification)

	cases := []struct {
		preRun    func() *types.MsgDeleteNotification
		expErr    bool
		expErrMsg string
		name      string
	}{
		{
			preRun: func() *types.MsgDeleteNotification {
				return types.NewMsgDeleteNotification(
					bob,
					alice,
					t,
				)
			},
			expErr: false,
			name:   "bob deletes his latest notification",
		},
		{
			preRun: func() *types.MsgDeleteNotification {
				return types.NewMsgDeleteNotification(
					bob,
					alice,
					t,
				)
			},
			expErr: true,
			name:   "cannot find already deleted notification",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.DeleteNotification(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {

				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgDeleteNotificationResponse{}, *res)

			}
		})
	}
}
