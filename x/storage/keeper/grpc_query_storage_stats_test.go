package keeper_test

import (
	"context"
	"math/rand"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/testutil"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
)

func (suite *KeeperTestSuite) TestStorageStats() {
	for i := 0; i < 100; i++ {

		mm := rand.Int63n(50_000_000_000)
		fs := rand.Int63n(1_000_000_000_000)

		suite.SetupSuite()

		testAddresses, err := testutil.CreateTestAddresses("cosmos", 4)
		suite.Require().NoError(err)

		testAccount := testAddresses[0]
		otherAccount := testAddresses[2]
		deadAccount := testAddresses[3]

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
			PricePerTbPerMonth:     8,
			CollateralPrice:        2,
			CheckWindow:            11,
			ReferralCommission:     25,
			PolRatio:               40,
		})

		suite.storageKeeper.SetStoragePaymentInfo(suite.ctx, types.StoragePaymentInfo{
			Start:          time.Date(2023, 10, 10, 0, 0, 0, 0, time.UTC),
			End:            time.Date(time.Now().Year()+1, 10, 10, 0, 0, 0, 0, time.UTC),
			SpaceAvailable: mm,
			SpaceUsed:      0,
			Address:        testAccount,
			Coins:          nil,
		})

		suite.storageKeeper.SetStoragePaymentInfo(suite.ctx, types.StoragePaymentInfo{ // dead account plan (counts for unique but not active)
			Start:          time.Date(2023, 10, 10, 0, 0, 0, 0, time.UTC),
			End:            time.Date(2023, 12, 10, 0, 0, 0, 0, time.UTC),
			SpaceAvailable: 5_000_000_000,
			SpaceUsed:      0,
			Address:        deadAccount,
			Coins:          nil,
		})

		merkle := []byte("merkle")

		suite.storageKeeper.SetFile(suite.ctx, types.UnifiedFile{
			Merkle:        merkle,
			Owner:         testAccount,
			Start:         0,
			Expires:       0,
			FileSize:      fs,
			ProofInterval: 400,
			ProofType:     0,
			Proofs:        make([]string, 0),
			MaxProofs:     3,
			Note:          "test",
		})

		suite.storageKeeper.SetFile(suite.ctx, types.UnifiedFile{
			Merkle:        merkle,
			Owner:         otherAccount,
			Start:         10,
			Expires:       100,
			FileSize:      fs,
			ProofInterval: 400,
			ProofType:     0,
			Proofs:        make([]string, 0),
			MaxProofs:     3,
			Note:          "test",
		})

		res, err := suite.queryClient.StorageStats(context.Background(), &types.QueryStorageStats{})
		suite.Require().NoError(err)

		suite.Require().Equal(uint64(2), res.ActiveUsers)
		suite.Require().Equal(uint64(3), res.UniqueUsers)
		suite.Require().Equal(uint64(fs*6), res.Used)

		suite.Require().Equal(fs*3+mm, int64(res.Purchased))

		suite.reset()
	}
}
