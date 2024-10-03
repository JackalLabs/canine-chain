package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	testutil "github.com/jackalLabs/canine-chain/v4/testutil"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
)

func (suite *KeeperTestSuite) TestPostFile() {
	suite.SetupSuite()
	msgSrvr, k, ctx := setupMsgServer(suite)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 3)
	suite.Require().NoError(err)

	testAccount := testAddresses[0]
	depoAccount := testAddresses[1]
	providerAccount := testAddresses[2]

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
		PricePerTbPerMonth:     8,
		CollateralPrice:        2,
		CheckWindow:            11,
		ReferralCommission:     25,
		PolRatio:               40,
	})

	suite.storageKeeper.SetProviders(suite.ctx, types.Providers{
		Address:         providerAccount,
		Ip:              "http://localhost:3333",
		Totalspace:      "10000000000",
		BurnedContracts: "0",
		Creator:         providerAccount,
		KeybaseIdentity: "",
		AuthClaimers:    nil,
	})

	initialPayInfo := types.StoragePaymentInfo{
		Start:          suite.ctx.BlockTime().AddDate(0, 0, -60),
		End:            suite.ctx.BlockTime().AddDate(0, 0, 30),
		SpaceAvailable: 100_000_000_000,
		SpaceUsed:      5_000_000_000,
		Address:        testAccount,
	}
	k.SetStoragePaymentInfo(suite.ctx, initialPayInfo)

	cases := []struct {
		testName  string
		msg       types.MsgPostFile
		expErr    bool
		tokens    int64
		expErrMsg string
	}{
		{
			testName: "post file",
			msg: types.MsgPostFile{
				Creator:       testAccount,
				Merkle:        []byte("merkle"),
				FileSize:      10,
				ProofInterval: 10,
				ProofType:     0,
				MaxProofs:     3,
				Expires:       0,
				Note:          "{}",
			},
			expErr: false,
		},
	}

	for _, tcs := range cases {
		tc := tcs
		suite.Run(tc.testName, func() {
			res, err := msgSrvr.PostFile(ctx, &tc.msg)
			if tc.expErr {
				suite.Require().ErrorContains(err, tc.expErrMsg)
			} else {
				suite.Require().NoError(err)

				suite.Require().Equal(0, len(res.ProviderIps)) // we do not pre-populate the provider IPs in v4 since the econ change. Change this to > 1 if ever the econ changes and required pre-population
			}
		})
	}
	suite.reset()
}
