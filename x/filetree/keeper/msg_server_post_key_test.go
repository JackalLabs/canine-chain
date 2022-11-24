package keeper_test

import (
	"fmt"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/filetree/types"
)

func (suite *KeeperTestSuite) TestMsgPostKey() {
	suite.SetupSuite()

	msgSrvr, _, context := setupMsgServer(suite)

	alice, err := sdkTypes.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	privateKey, err := types.MakePrivateKey("alice") // clientCtx.FromName in the CLI will be alice's keyring name (alice), not the full account address
	suite.Require().NoError(err)

	pubKey := privateKey.PublicKey.Bytes(false)

	cases := []struct {
		preRun    func() *types.MsgPostkey
		expErr    bool
		expErrMsg string
		name      string
	}{
		{
			preRun: func() *types.MsgPostkey {
				return types.NewMsgPostkey(
					alice.String(),
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
			res, err := msgSrvr.Postkey(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgPostkeyResponse{}, *res)

			}
		})
	}
}
