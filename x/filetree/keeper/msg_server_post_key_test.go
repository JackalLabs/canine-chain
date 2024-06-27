package keeper_test

import (
	"fmt"

	"github.com/jackalLabs/canine-chain/v4/testutil"
	"github.com/jackalLabs/canine-chain/v4/x/filetree/types"
)

func (suite *KeeperTestSuite) TestMsgPostKey() {
	suite.SetupSuite()

	msgSrvr, context := setupMsgServer(suite)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 1)
	suite.Require().NoError(err)

	alice := testAddresses[0]

	privateKey, err := types.MakePrivateKey("alice") // clientCtx.FromName in the CLI will be alice's keyring ID (alice), not the full account address
	suite.Require().NoError(err)

	pubKey := privateKey.PublicKey.Bytes(false)

	cases := []struct {
		preRun    func() *types.MsgPostKey
		expErr    bool
		expErrMsg string
		name      string
	}{
		{
			preRun: func() *types.MsgPostKey {
				return types.NewMsgPostKey(
					alice,
					fmt.Sprintf("%x", pubKey),
				)
			},
			expErr: false,
			name:   "post key success",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.PostKey(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgPostKeyResponse{}, *res)

			}
		})
	}
}
