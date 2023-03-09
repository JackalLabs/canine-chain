package keeper_test

import (
	"encoding/json"

	"github.com/jackalLabs/canine-chain/testutil"
	"github.com/jackalLabs/canine-chain/x/notifications/types"
)

func (suite *KeeperTestSuite) TestMsgCreateNotifications() {
	suite.SetupSuite()
	msgSrvr, _, context := setupMsgServer(suite)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 3)
	suite.Require().NoError(err)

	alice := testAddresses[0]
	bob := testAddresses[1]
	charlie := testAddresses[2]

	// set noti counter for bob

	blockedSendersMap := make([]string, 0, 1000)
	blockedSendersMap = append(blockedSendersMap, charlie)

	marshalledBlockedSenders, err := json.Marshal(blockedSendersMap)
	suite.Require().NoError(err)
	BlockedSenders := string(marshalledBlockedSenders)

	notiCounter := types.NotiCounter{
		Address:        bob,
		Counter:        0,
		BlockedSenders: BlockedSenders,
	}
	suite.Require().NoError(err)
	suite.notificationsKeeper.SetNotiCounter(suite.ctx, notiCounter)

	cases := []struct {
		preRun    func() *types.MsgCreateNotifications
		expErr    bool
		expErrMsg string
		name      string
	}{
		{ // alice sends a notification to bob
			preRun: func() *types.MsgCreateNotifications {
				return types.NewMsgCreateNotifications(
					alice,
					"hey bob it's alice here",
					bob,
				)
			},
			expErr: false,
			name:   "alice successfully sends a notification to bob",
		},
		{ // alice can't notify casper if he doesn't have his notiCounter set
			preRun: func() *types.MsgCreateNotifications {
				return types.NewMsgCreateNotifications(
					alice,
					"hey casper it's alice here",
					"casper",
				)
			},
			expErr:    true,
			name:      "alice cannot notify casper",
			expErrMsg: "User's notiCounter not set",
		},
		{ // charlie cannot notify bob
			preRun: func() *types.MsgCreateNotifications {
				return types.NewMsgCreateNotifications(
					charlie,
					"hey bob it's charlie",
					bob,
				)
			},
			expErr:    true,
			name:      "charlie cannot notify bob",
			expErrMsg: "You are a blocked sender",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.CreateNotifications(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {

				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgCreateNotificationsResponse{NotiCounter: 1}, *res)

			}
		})
	}
}
