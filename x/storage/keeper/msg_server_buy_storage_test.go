package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func (suite *KeeperTestSuite) TestBuyStorage() {
	suite.SetupSuite()
	msgSrvr, k, ctx := setupMsgServer(suite)

	// Create test account
	testAccount, err := sdk.AccAddressFromBech32("cosmos17j2hkm7n9fz9dpntyj2kxgxy5pthzd289nvlfl")
	suite.Require().NoError(err)

	depoAccount, err := sdk.AccAddressFromBech32("cosmos1arsaayyj5tash86mwqudmcs2fd5jt5zgp07gl8")
	suite.Require().NoError(err)

	coins := sdk.NewCoins(sdk.NewCoin("ujkl", sdk.NewInt(100000000000))) // Send some coins to their account
	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, testAccount, coins)
	suite.Require().NoError(err)

	suite.storageKeeper.SetParams(suite.ctx, types.Params{
		DepositAccount: depoAccount.String(),
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
		tokens    int64
		expErrMsg string
	}{
		{
			testName: "buy 3gb which is less than current usage of 5gb",
			msg: types.MsgBuyStorage{
				Creator:      testAccount.String(),
				ForAddress:   testAccount.String(),
				Duration:     "1",
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
				Duration:     "1",
				Bytes:        "6000000000",
				PaymentDenom: "ujkl",
			},
			expErr:    false,
			tokens:    200000,
			expErrMsg: "",
		},
		{
			testName: "buy 1tb for 3 month",
			msg: types.MsgBuyStorage{
				Creator:      testAccount.String(),
				ForAddress:   testAccount.String(),
				Duration:     "3",
				Bytes:        "1000000000000",
				PaymentDenom: "ujkl",
			},
			expErr:    false,
			tokens:    100000000,
			expErrMsg: "",
		},
		{
			testName: "buy less than a gb",
			msg: types.MsgBuyStorage{
				Creator:      testAccount.String(),
				ForAddress:   testAccount.String(),
				Duration:     "1",
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
				Duration:     "0",
				Bytes:        "1000000000",
				PaymentDenom: "ujkl",
			},
			expErr:    true,
			expErrMsg: "duration can't be less than 1 month",
		},
		{
			// TODO: update this when we allow alt payments
			testName: "payment with uatom",
			msg: types.MsgBuyStorage{
				Creator:      testAccount.String(),
				ForAddress:   testAccount.String(),
				Duration:     "1",
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
				Duration:     "1",
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
				Duration:     "1",
				Bytes:        "1000000000",
				PaymentDenom: "ujkl",
			},
			expErr:    true,
			expErrMsg: "decoding bech32 failed: invalid separator index -1",
		},
	}

	dep := k.GetParams(suite.ctx).DepositAccount
	fmt.Println(dep)
	add, err := sdk.AccAddressFromBech32(dep)
	suite.Require().NoError(err)
	amt := suite.bankKeeper.GetBalance(suite.ctx, add, "ujkl").Amount.Int64()
	fmt.Println(amt)
	for _, tc := range cases {
		suite.Run(tc.testName, func() {
			_, err := msgSrvr.BuyStorage(ctx, &tc.msg)

			bal := suite.bankKeeper.GetBalance(suite.ctx, add, "ujkl")
			diff := bal.Amount.Int64() - amt
			fmt.Println(diff)
			amt = bal.Amount.Int64()
			if tc.expErr {
				suite.Require().Equal(int64(0), diff)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.tokens, diff)

			}
		})
	}
	suite.reset()
}
