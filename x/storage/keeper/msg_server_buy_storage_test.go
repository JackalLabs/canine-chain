package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func (suite *KeeperTestSuite) TestBuyStorage() {
	suite.SetupSuite()
	msgSrvr, k, ctx := setupMsgServer(suite)

	// Create test account
	testAccount, err := sdk.AccAddressFromBech32("cosmos17j2hkm7n9fz9dpntyj2kxgxy5pthzd289nvlfl")
	suite.Require().NoError(err)
	coins := sdk.NewCoins(sdk.NewCoin("ujkl", sdk.NewInt(100000000))) // Send some coins to their account
	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, testAccount, coins)
	suite.Require().NoError(err)

	suite.storageKeeper.SetParams(suite.ctx, types.Params{
		DepositAccount: testAccount.String(),
	})

	// Set user current SpaceUsed to 5GB
	initialPayInfo := types.StoragePaymentInfo{
		SpaceUsed: 5000000000,
		Address:   testAccount.String(),
	}
	k.SetStoragePaymentInfo(suite.ctx, initialPayInfo)

	cases := []struct {
		testName  string
		msg       types.MsgBuyStorage
		expErr    bool
		expErrMsg string
	}{
		{
			testName: "buy 3gb which is less than current usage of 5gb",
			msg: types.MsgBuyStorage{
				Creator:      testAccount.String(),
				ForAddress:   testAccount.String(),
				Duration:     "1440h",
				Bytes:        "3000000000",
				PaymentDenom: "ujkl",
			},
			expErr:    true,
			expErrMsg: "cannot buy less than your current gb usage",
		},
		{
			testName: "buy 6gb for 1 month",
			msg: types.MsgBuyStorage{
				Creator:      testAccount.String(),
				ForAddress:   testAccount.String(),
				Duration:     "1440h",
				Bytes:        "6000000000",
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
				Duration:     "1440h",
				Bytes:        "-1",
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
				Duration:     "1h",
				Bytes:        "1000000000",
				PaymentDenom: "ujkl",
			},
			expErr:    true,
			expErrMsg: "cannot buy less than a month",
		},
		{
			// TODO: update this when we allow alt payments
			testName: "payment with uatom",
			msg: types.MsgBuyStorage{
				Creator:      testAccount.String(),
				ForAddress:   testAccount.String(),
				Duration:     "400000h",
				Bytes:        "1000000000",
				PaymentDenom: "uatom",
			},
			expErr:    true,
			expErrMsg: "cannot pay with anything other than ujkl: invalid coins",
		},
		{
			testName: "invalid creator address",
			msg: types.MsgBuyStorage{
				Creator:      "invalid_address",
				ForAddress:   testAccount.String(),
				Duration:     "400000h",
				Bytes:        "1000000000",
				PaymentDenom: "ujkl",
			},
			expErr:    true,
			expErrMsg: "decoding bech32 failed: invalid separator index -1",
		},
		{
			testName: "invalid for address",
			msg: types.MsgBuyStorage{
				Creator:      testAccount.String(),
				ForAddress:   "invalid_address",
				Duration:     "432000h",
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
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
	suite.reset()
}
