package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	testutil "github.com/jackalLabs/canine-chain/v4/testutil"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
)

func (suite *KeeperTestSuite) TestUpgradeStorage() {
	suite.SetupSuite()
	msgSrvr, k, ctx := setupMsgServer(suite)

	// Create test account
	testAddresses, err := testutil.CreateTestAddresses("cosmos", 1)
	suite.Require().NoError(err)

	testAccount := testAddresses[0]

	coins := sdk.NewCoins(sdk.NewCoin("ujkl", sdk.NewInt(100000000000))) // Send some coins to their account
	testAcc, _ := sdk.AccAddressFromBech32(testAccount)
	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, testAcc, coins)
	suite.Require().NoError(err)

	suite.storageKeeper.SetParams(suite.ctx, types.Params{
		DepositAccount:         testAccount,
		ProofWindow:            50,
		ChunkSize:              1024,
		PriceFeed:              "jklprice",
		MissesToBurn:           3,
		MaxContractAgeInBlocks: 100,
		PricePerTbPerMonth:     8,
		CollateralPrice:        2,
		CheckWindow:            10,
	})

	cases := []struct {
		testName  string
		preRun    func()
		msg       types.MsgBuyStorage
		expErr    bool
		expErrMsg string
	}{
		{
			testName: "upgrade to 6gb for 2 month",
			preRun: func() {
				// Set a 3 months plan of 5GB starts 2 months ago, ends in a month
				initialPayInfo := types.StoragePaymentInfo{
					Start:          suite.ctx.BlockTime().AddDate(0, 0, -60),
					End:            suite.ctx.BlockTime().AddDate(0, 0, 30),
					SpaceAvailable: 5000000000,
					SpaceUsed:      4000000000,
					Address:        testAccount,
				}
				k.SetStoragePaymentInfo(suite.ctx, initialPayInfo)
			},
			msg: types.MsgBuyStorage{
				Creator:      testAccount,
				ForAddress:   testAccount,
				DurationDays: 30,
				Bytes:        6000000000,
				PaymentDenom: "ujkl",
			},
			expErr:    false,
			expErrMsg: "",
		},
		{
			testName: "downgrading with refund higher than new cost",
			preRun: func() {
				// Set a 3 months plan of 5GB starts 2 months ago, ends in a month
				initialPayInfo := types.StoragePaymentInfo{
					Start:          suite.ctx.BlockTime().AddDate(0, 0, -60),
					End:            suite.ctx.BlockTime().AddDate(0, 0, 30),
					SpaceAvailable: 5000000000,
					SpaceUsed:      4000000000,
					Address:        testAccount,
				}
				k.SetStoragePaymentInfo(suite.ctx, initialPayInfo)
			},
			msg: types.MsgBuyStorage{
				Creator:      testAccount,
				ForAddress:   testAccount,
				DurationDays: 30,
				Bytes:        4000000000,
				PaymentDenom: "ujkl",
			},
			expErr:    true,
			expErrMsg: "cannot downgrade until current plan expires: invalid request",
		},
		{
			testName: "downgrading from 10GB to 8GB ",
			preRun: func() {
				// Set a 3 months plan of 5GB starts 2 months ago, ends in a month
				initialPayInfo := types.StoragePaymentInfo{
					Start:          suite.ctx.BlockTime().AddDate(0, 0, -60),
					End:            suite.ctx.BlockTime().AddDate(0, 0, 30),
					SpaceAvailable: 10000000000,
					SpaceUsed:      4000000000,
					Address:        testAccount,
				}
				k.SetStoragePaymentInfo(suite.ctx, initialPayInfo)
			},
			msg: types.MsgBuyStorage{
				Creator:      testAccount,
				ForAddress:   testAccount,
				DurationDays: 30 * 4,
				Bytes:        8000000000,
				PaymentDenom: "ujkl",
			},
			expErr:    false,
			expErrMsg: "",
		},
		{
			testName: "upgrading expired plan", // upgrading an expired plan is the same as buying a new plan
			preRun: func() {
				// Set a 3 months plan of 5GB starts 2 months ago, ends a month ago
				initialPayInfo := types.StoragePaymentInfo{
					Start:          suite.ctx.BlockTime().AddDate(0, 0, -60),
					End:            suite.ctx.BlockTime().AddDate(0, 0, -30),
					SpaceAvailable: 10000000000,
					SpaceUsed:      4000000000,
					Address:        testAccount,
				}
				k.SetStoragePaymentInfo(suite.ctx, initialPayInfo)
			},
			msg: types.MsgBuyStorage{
				Creator:      testAccount,
				ForAddress:   testAccount,
				DurationDays: 30,
				Bytes:        8000000000,
				PaymentDenom: "ujkl",
			},
			expErr: false,
		},
		{
			testName: "downgrading to buy less gb than current usage",
			preRun: func() {
				initialPayInfo := types.StoragePaymentInfo{
					Start:          suite.ctx.BlockTime().AddDate(0, 0, -60),
					End:            suite.ctx.BlockTime().AddDate(0, 0, 30),
					SpaceAvailable: 10000000000,
					SpaceUsed:      4000000000,
					Address:        testAccount,
				}
				k.SetStoragePaymentInfo(suite.ctx, initialPayInfo)
			},
			msg: types.MsgBuyStorage{
				Creator:      testAccount,
				ForAddress:   testAccount,
				DurationDays: 30,
				Bytes:        3000000000,
				PaymentDenom: "ujkl",
			},
			expErr:    true,
			expErrMsg: "cannot buy less than your current gb usage", // fail a downgrade
		},
	}

	for _, tcs := range cases {
		tc := tcs
		suite.Run(tc.testName, func() {
			if tc.preRun != nil {
				tc.preRun()
			}
			_, err := msgSrvr.BuyStorage(ctx, &tc.msg)
			if tc.expErr {
				suite.Require().ErrorContains(err, tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
			}
			k.RemoveStoragePaymentInfo(suite.ctx, testAccount)
		})
	}
	suite.reset()
}
