package keeper_test

import (
	"encoding/json"

	"github.com/jackalLabs/canine-chain/testutil"
	"github.com/jackalLabs/canine-chain/x/notifications/types"
)

func (suite *KeeperTestSuite) TestMsgDeleteNotifications() {
	suite.SetupSuite()
	msgSrvr, context := setupMsgServer(suite)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	suite.Require().NoError(err)

	alice := testAddresses[0]
	bob := testAddresses[1]

	// set noti counter for bob

	placeholderMap := make([]string, 0, 1000)
	marshalledBlockedSenders, err := json.Marshal(placeholderMap)
	suite.Require().NoError(err)
	BlockedSenders := string(marshalledBlockedSenders)

	notiCounter := types.NotiCounter{
		Address:        bob,
		Counter:        0,
		BlockedSenders: BlockedSenders,
	}
	suite.Require().NoError(err)
	suite.notificationsKeeper.SetNotiCounter(suite.ctx, notiCounter)

	notification := types.Notifications{
		Count:        notiCounter.Counter,
		Notification: "hey bob it's alice",
		Address:      bob,
		Sender:       alice,
	}
	suite.Require().NoError(err)
	suite.notificationsKeeper.SetNotifications(suite.ctx, notification, bob)

	// Increase the notiCounter just as it happens in the keeper code

	notiCounter.Counter++
	suite.Require().NoError(err)
	suite.notificationsKeeper.SetNotiCounter(suite.ctx, notiCounter)

	cases := []struct {
		preRun    func() *types.MsgDeleteNotifications
		expErr    bool
		expErrMsg string
		name      string
	}{
		{ // bob deletes his latest notification
			preRun: func() *types.MsgDeleteNotifications {
				return types.NewMsgDeleteNotifications(
					bob,
				)
			},
			expErr: false,
			name:   "bob deletes his latest notification",
		},
		{ // alice has nothing to delete
			preRun: func() *types.MsgDeleteNotifications {
				return types.NewMsgDeleteNotifications(
					alice,
				)
			},
			expErr:    true,
			name:      "alice has nothing to delete",
			expErrMsg: "User's notiCounter not set",
		},
		{ // no more notifications for bob to delete
			preRun: func() *types.MsgDeleteNotifications {
				return types.NewMsgDeleteNotifications(
					bob,
				)
			},
			expErr:    true,
			name:      "no more notifications for bob to delete",
			expErrMsg: "Notification does not exist",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.DeleteNotifications(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {

				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgDeleteNotificationsResponse{}, *res)

			}
		})
	}
}
