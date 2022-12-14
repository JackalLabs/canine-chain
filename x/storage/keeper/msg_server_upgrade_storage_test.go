package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func (suite *KeeperTestSuite) TestUpgradeStorage() {
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

	cases := []struct {
		testName  string
		preRun    func()
		msg       types.MsgUpgradeStorage
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
					Address:        testAccount.String(),
				}
				k.SetStoragePaymentInfo(suite.ctx, initialPayInfo)
			},
			msg: types.MsgUpgradeStorage{
				Creator:      testAccount.String(),
				ForAddress:   testAccount.String(),
				Duration:     "2",
				Bytes:        "6000000000",
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
					Address:        testAccount.String(),
				}
				k.SetStoragePaymentInfo(suite.ctx, initialPayInfo)
			},
			msg: types.MsgUpgradeStorage{
				Creator:      testAccount.String(),
				ForAddress:   testAccount.String(),
				Duration:     "1",
				Bytes:        "4000000000",
				PaymentDenom: "ujkl",
			},
			expErr:    true,
			expErrMsg: "cannot downgrade at the moment, please wait till your subscription is over",
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
					Address:        testAccount.String(),
				}
				k.SetStoragePaymentInfo(suite.ctx, initialPayInfo)
			},
			msg: types.MsgUpgradeStorage{
				Creator:      testAccount.String(),
				ForAddress:   testAccount.String(),
				Duration:     "3",
				Bytes:        "8000000000",
				PaymentDenom: "ujkl",
			},
			expErr:    false,
			expErrMsg: "",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.testName, func() {
			if tc.preRun != nil {
				tc.preRun()
			}
			_, err := msgSrvr.UpgradeStorage(ctx, &tc.msg)
			if tc.expErr {
				suite.Require().ErrorContains(err, tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
			}
			k.RemoveStoragePaymentInfo(suite.ctx, testAccount.String())
		})
	}
	suite.reset()
}
