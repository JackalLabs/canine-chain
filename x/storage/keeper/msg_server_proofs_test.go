package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func (suite *KeeperTestSuite) TestPostProof() {
	suite.SetupSuite()

	// Create user account
	user, err := sdk.AccAddressFromBech32("cosmos17j2hkm7n9fz9dpntyj2kxgxy5pthzd289nvlfl")
	suite.Require().NoError(err)

	// Create provider account
	testProvider, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	msgSrvr, _, context := setupMsgServer(suite)

	// Init Provider
	_, err = msgSrvr.InitProvider(context, &types.MsgInitProvider{
		Creator:    testProvider.String(),
		Ip:         "198.0.0.1",
		Totalspace: "1_000_000",
	})
	if err != nil {
		fmt.Println(err)
	}

	const CID = "6ef1cf960c0b1e257049645ff13ed890c2d4ef69c62165bf4e090ec480770d67"

	// Post Contract
	_, err = msgSrvr.PostContract(context, &types.MsgPostContract{
		Creator:  testProvider.String(),
		Signee:   user.String(),
		Duration: "1",
		Filesize: "1_000",
		Fid:      "fid",
		Merkle:   "merkle",
	})
	if err != nil {
		fmt.Println(err)
	}

	// Sign Contract for active deal
	_, err = msgSrvr.SignContract(context, &types.MsgSignContract{
		Creator: user.String(),
		Cid:     CID,
	})
	if err != nil {
		fmt.Println(err)
	}

	// item, hashlist, err :=

	// cases := []struct {
	// 	testName  string
	// 	msg       types.MsgPostproof
	// 	expErr    bool
	// 	expErrMsg string
	// }{
	// 	{
	// 		testName: "create 1 proof",
	// 		msg: types.MsgPostproof{
	// 			Creator:  testProvider.String(),
	// 			Cid:      CID,
	// 			Item:     "this is item",
	// 			Hashlist: "put hashlist here",
	// 		},
	// 		expErr:    false,
	// 		expErrMsg: "",
	// 	},
	// }

	// for _, tc := range cases {
	// 	suite.Run(
	// 		tc.testName, func() {
	// 			_, err := msgSrvr.Postproof(context, &tc.msg)
	// 			fmt.Println(err)
	// 		},
	// 	)
	// }
}
