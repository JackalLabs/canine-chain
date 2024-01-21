package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	testutil "github.com/jackalLabs/canine-chain/v3/testutil"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

func (suite *KeeperTestSuite) TestBuyStorage() {
	suite.SetupSuite()
	msgSrvr, k, ctx := setupMsgServer(suite)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	suite.Require().NoError(err)

	testAccount := testAddresses[0]
	depoAccount := testAddresses[1]

	coins := sdk.NewCoins(sdk.NewCoin("ujkl", sdk.NewInt(100000000000))) // Send some coins to their account
	testAcc, _ := sdk.AccAddressFromBech32(testAccount)
	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, testAcc, coins)
	suite.Require().NoError(err)

	suite.storageKeeper.SetParams(suite.ctx, types.Params{
		DepositAccount:         depoAccount,
		ProofWindow:            50,
		ChunkSize:              1024,
		PriceFeed:              "jklprice",
		MissesToBurn:           3,
		MaxContractAgeInBlocks: 100,
		PricePerTbPerMonth:     15,
		CollateralPrice:        2,
		CheckWindow:            10,
	})

	cases := []struct {
		testName  string
		preRun    func()
		msg       types.MsgBuyStorage
		expErr    bool
		tokens    int64
		expErrMsg string
	}{
		{
			testName: "buy less storage than active while having an active plan",
			preRun: func() {
				initialPayInfo := types.StoragePaymentInfo{
					Start:          suite.ctx.BlockTime().AddDate(0, 0, -60),
					End:            suite.ctx.BlockTime().AddDate(0, 0, 30),
					SpaceAvailable: 100_000_000_000,
					SpaceUsed:      5_000_000_000,
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
			expErr:    true,
			expErrMsg: "cannot downgrade until current plan expires: invalid request",
		},
		{
			testName: "buy 3gb which is less than current usage of 5gb",
			preRun: func() {
				// Set user current SpaceUsed to 5GB
				initialPayInfo := types.StoragePaymentInfo{
					SpaceUsed: 5000000000,
					Address:   testAccount,
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
			expErrMsg: "cannot buy less than your current gb usage",
		},
		{
			testName: "successfully buy 6gb for 1 month",
			preRun: func() {
				initialPayInfo := types.StoragePaymentInfo{
					Start:          suite.ctx.BlockTime().AddDate(0, 0, -60),
					End:            suite.ctx.BlockTime().AddDate(0, 0, -1),
					SpaceAvailable: 100_000_000_000,
					SpaceUsed:      5_000_000_000,
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
			tokens:    43749,
			expErrMsg: "",
		},
		{
			testName: "successfully buy 1tb for 3 month",
			msg: types.MsgBuyStorage{
				Creator:      testAccount,
				ForAddress:   testAccount,
				DurationDays: 90,
				Bytes:        1000000000000,
				PaymentDenom: "ujkl",
			},
			expErr:    false,
			tokens:    21874999,
			expErrMsg: "",
		},
		{
			testName: "successfully buy 1tb for 3 month with referral",
			msg: types.MsgBuyStorage{
				Creator:      testAccount,
				ForAddress:   testAccount,
				DurationDays: 90,
				Bytes:        1000000000000,
				PaymentDenom: "ujkl",
				Referral:     depoAccount,
			},
			expErr:    false,
			tokens:    19687499,
			expErrMsg: "",
		},
		{
			testName: "buy less than a gb",
			msg: types.MsgBuyStorage{
				Creator:      testAccount,
				ForAddress:   testAccount,
				DurationDays: 30,
				Bytes:        -1,
				PaymentDenom: "ujkl",
			},
			expErr:    true,
			expErrMsg: "cannot buy less than a gb",
		},
		{
			testName: "buy less than a month",
			msg: types.MsgBuyStorage{
				Creator:      testAccount,
				ForAddress:   testAccount,
				DurationDays: 1,
				Bytes:        1000000000,
				PaymentDenom: "ujkl",
			},
			expErr:    true,
			expErrMsg: "duration can't be less than 1 month",
		},
		{
			// TODO: update this when we allow alt payments
			testName: "payment with uatom",
			msg: types.MsgBuyStorage{
				Creator:      testAccount,
				ForAddress:   testAccount,
				DurationDays: 30,
				Bytes:        1000000000,
				PaymentDenom: "uatom",
			},
			expErr:    true,
			expErrMsg: "cannot pay with anything other than ujkl: invalid coins",
		},
		{
			testName: "invalid creator address",
			msg: types.MsgBuyStorage{
				Creator:      "invalid_address",
				ForAddress:   testAccount,
				DurationDays: 30,
				Bytes:        1000000000,
				PaymentDenom: "ujkl",
			},
			expErr:    true,
			expErrMsg: "decoding bech32 failed: invalid separator index -1",
		},
		{
			testName: "invalid for address",
			msg: types.MsgBuyStorage{
				Creator:      testAccount,
				ForAddress:   "invalid_address",
				DurationDays: 30,
				Bytes:        1000000000,
				PaymentDenom: "ujkl",
			},
			expErr:    true,
			expErrMsg: "decoding bech32 failed: invalid separator index -1",
		},
	}

	add, err := types.GetTokenHolderAccount()
	suite.Require().NoError(err)
	amt := suite.bankKeeper.GetBalance(suite.ctx, add, "ujkl").Amount.Int64()

	for _, tcs := range cases {
		tc := tcs
		suite.Run(tc.testName, func() {
			if tc.preRun != nil {
				tc.preRun()
			}
			_, err := msgSrvr.BuyStorage(ctx, &tc.msg)

			bal := suite.bankKeeper.GetBalance(suite.ctx, add, "ujkl")
			diff := bal.Amount.Int64() - amt
			amt = bal.Amount.Int64()

			if tc.expErr {
				suite.Require().Equal(int64(0), diff)
				suite.Require().ErrorContains(err, tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.tokens, diff)
			}

			k.RemoveStoragePaymentInfo(suite.ctx, testAccount)
		})
	}
	suite.reset()
}

func (suite *KeeperTestSuite) TestBuyStorageValues() {
	suite.SetupSuite()
	msgSrvr, k, ctx := setupMsgServer(suite)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 3)
	suite.Require().NoError(err)

	testAccount := testAddresses[0]
	depoAccount := testAddresses[1]

	coins := sdk.NewCoins(sdk.NewCoin("ujkl", sdk.NewInt(100000000000))) // Send some coins to their account
	testAcc, _ := sdk.AccAddressFromBech32(testAccount)
	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, testAcc, coins)
	suite.Require().NoError(err)

	suite.storageKeeper.SetParams(suite.ctx, types.Params{
		DepositAccount:         depoAccount,
		ProofWindow:            50,
		ChunkSize:              1024,
		PriceFeed:              "jklprice",
		MissesToBurn:           3,
		MaxContractAgeInBlocks: 100,
		PricePerTbPerMonth:     15,
		CollateralPrice:        2,
		CheckWindow:            10,
	})

	var bytes int64 = 3_000_000_000_000
	var days int64 = 30

	_, err = msgSrvr.BuyStorage(ctx, &types.MsgBuyStorage{
		Creator:      testAccount,
		ForAddress:   testAccount,
		DurationDays: days,
		Bytes:        bytes,
		PaymentDenom: "ujkl",
		Referral:     "",
	})
	suite.Require().NoError(err)

	cost := float64(suite.storageKeeper.GetStorageCost(suite.ctx, bytes/1_000_000_000, days*24).Int64())

	providerAccount, err := types.GetTokenHolderAccount()
	suite.Require().NoError(err)

	bal := suite.bankKeeper.GetBalance(suite.ctx, providerAccount, "ujkl")
	suite.Require().Equal(int64(cost*0.35), bal.Amount.Int64())

	polAccount, err := types.GetPOLAccount()
	suite.Require().NoError(err)

	bal = suite.bankKeeper.GetBalance(suite.ctx, polAccount, "ujkl")
	suite.Require().Equal(int64(cost*0.40), bal.Amount.Int64())

	_ = k

	suite.reset()
}

func (suite *KeeperTestSuite) TestBuyStorageReferralValues() {
	suite.SetupSuite()
	msgSrvr, k, ctx := setupMsgServer(suite)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 3)
	suite.Require().NoError(err)

	testAccount := testAddresses[0]
	depoAccount := testAddresses[1]

	coins := sdk.NewCoins(sdk.NewCoin("ujkl", sdk.NewInt(100000000000))) // Send some coins to their account
	testAcc, _ := sdk.AccAddressFromBech32(testAccount)
	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, testAcc, coins)
	suite.Require().NoError(err)

	suite.storageKeeper.SetParams(suite.ctx, types.Params{
		DepositAccount:         depoAccount,
		ProofWindow:            50,
		ChunkSize:              1024,
		PriceFeed:              "jklprice",
		MissesToBurn:           3,
		MaxContractAgeInBlocks: 100,
		PricePerTbPerMonth:     15,
		CollateralPrice:        2,
		CheckWindow:            10,
	})

	var bytes int64 = 3_000_000_000_000
	var days int64 = 30

	_, err = msgSrvr.BuyStorage(ctx, &types.MsgBuyStorage{
		Creator:      testAccount,
		ForAddress:   testAccount,
		DurationDays: days,
		Bytes:        bytes,
		PaymentDenom: "ujkl",
		Referral:     depoAccount,
	})
	suite.Require().NoError(err)

	cost := float64(suite.storageKeeper.GetStorageCost(suite.ctx, bytes/1_000_000_000, days*24).Int64()) * 0.90

	providerAccount, err := types.GetTokenHolderAccount()
	suite.Require().NoError(err)

	bal := suite.bankKeeper.GetBalance(suite.ctx, providerAccount, "ujkl")
	suite.Require().Equal(int64(cost*0.35), bal.Amount.Int64())

	polAccount, err := types.GetPOLAccount()
	suite.Require().NoError(err)

	bal = suite.bankKeeper.GetBalance(suite.ctx, polAccount, "ujkl")
	suite.Require().Equal(int64(cost*0.30), bal.Amount.Int64())

	_ = k

	suite.reset()
}
