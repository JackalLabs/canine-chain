package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func (suite *KeeperTestSuite) TestBuyStorage() {
	suite.SetupSuite()
	msgSrvr, _, ctx := setupMsgServer(suite)

	// Create test account
	testAccount, err := sdk.AccAddressFromBech32("cosmos17j2hkm7n9fz9dpntyj2kxgxy5pthzd289nvlfl")
	suite.Require().NoError(err)
	coins := sdk.NewCoins(sdk.NewCoin("ujkl", sdk.NewInt(100000000))) // Send some coins to their account
	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, testAccount, coins)
	suite.Require().NoError(err)

	cases := []struct {
		testName  string
		msg       types.MsgBuyStorage
		expErr    bool
		expErrMsg string
	}{
		{
			testName: "buy 1gb for 2 month",
			msg: types.MsgBuyStorage{
				Creator:      testAccount.String(),
				ForAddress:   "testAccount.String()",
				Duration:     "432000",
				Bytes:        "2000000000",
				PaymentDenom: "ujkl",
			},
			expErr:    false,
			expErrMsg: "",
		},
		{
			testName: "buy less than a gb",
			msg: types.MsgBuyStorage{
				Creator:      testAccount.String(),
				ForAddress:   testAccount.String(),
				Duration:     "432000",
				Bytes:        "12345",
				PaymentDenom: "ujkl",
			},
			expErr:    true,
			expErrMsg: "cannot buy less than a gb",
		},
		{
			testName: "buy less than a month",
			msg: types.MsgBuyStorage{
				Creator:      testAccount.String(),
				ForAddress:   testAccount.String(),
				Duration:     "431999",
				Bytes:        "1000000000",
				PaymentDenom: "ujkl",
			},
			expErr:    true,
			expErrMsg: "cannot buy less than a month",
		},
		{
			testName: "payment with uatom",
			msg: types.MsgBuyStorage{
				Creator:      testAccount.String(),
				ForAddress:   testAccount.String(),
				Duration:     "432000",
				Bytes:        "1000000000",
				PaymentDenom: "uatom",
			},
			expErr:    true,
			expErrMsg: "cannot pay with anything other than ujkl: invalid coins",
		},
		{
			testName: "invalid address",
			msg: types.MsgBuyStorage{
				Creator:      "invalid_address",
				ForAddress:   "invalid_address",
				Duration:     "432000",
				Bytes:        "1000000000",
				PaymentDenom: "ujkl",
			},
			expErr:    true,
			expErrMsg: "decoding bech32 failed: invalid separator index -1",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.testName, func() {
			_, err := msgSrvr.BuyStorage(ctx, &tc.msg)
			if tc.expErr {
				suite.Require().EqualError(err, tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}
