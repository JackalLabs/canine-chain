package keeper_test

import (
	"encoding/json"

	"github.com/jackalLabs/canine-chain/testutil"
	"github.com/jackalLabs/canine-chain/x/notifications/types"
)

func (suite *KeeperTestSuite) TestMsgBlockSenders() {
	suite.SetupSuite()
	msgSrvr, _, context := setupMsgServer(suite)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 3)
	suite.Require().NoError(err)

	alice := testAddresses[0]
	bob := testAddresses[1]
	charlie := testAddresses[2]

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

	var senderIds []string
	senderIds = append(senderIds, alice, charlie)
	jsonSenders, err := json.Marshal(senderIds)
	suite.Require().NoError(err)

	cases := []struct {
		preRun    func() *types.MsgBlockSenders
		expErr    bool
		expErrMsg string
		name      string
	}{
		{ // bob blocks some senders
			preRun: func() *types.MsgBlockSenders {
				return types.NewMsgBlockSenders(
					bob,
					string(jsonSenders),
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
			res, err := msgSrvr.BlockSenders(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {

				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgBlockSendersResponse{}, *res)
				notiCounterReq := types.QueryGetNotiCounterRequest{
					Address: bob,
				}
				res, err := suite.queryClient.NotiCounter(suite.ctx.Context(), &notiCounterReq)
				suite.Require().NoError(err)

				blockedSenders := res.NotiCounter.BlockedSenders
				blockedSendersMap := make([]string, 0, 1000)
				error := json.Unmarshal([]byte(blockedSenders), &blockedSendersMap)
				suite.Require().NoError(error)

				suite.Require().EqualValues(len(blockedSendersMap), 2)
				suite.Require().EqualValues(blockedSendersMap[0], alice)
				suite.Require().EqualValues(blockedSendersMap[1], charlie)

			}
		})
	}
}
