package keeper_test

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jackalLabs/canine-chain/v4/testutil"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
)

func (suite *KeeperTestSuite) TestFindFile() {
	suite.SetupSuite()

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
		CheckWindow:            10,
	})

	merkle := []byte("merkle")

	f := types.UnifiedFile{
		Merkle:        merkle,
		Owner:         testAccount,
		Start:         0,
		Expires:       0,
		FileSize:      1024,
		ProofInterval: 400,
		ProofType:     0,
		Proofs:        make([]string, 0),
		MaxProofs:     3,
		Note:          "test",
	}

	suite.storageKeeper.SetFile(suite.ctx, f)

	suite.storageKeeper.SetProviders(suite.ctx, types.Providers{
		Address:         providerAccount,
		Ip:              "http://localhost:3333",
		Totalspace:      "10000000000",
		BurnedContracts: "0",
		Creator:         providerAccount,
		KeybaseIdentity: "",
		AuthClaimers:    nil,
	})

	f.AddProver(suite.ctx, suite.storageKeeper, providerAccount)

	pg := query.PageRequest{
		Offset:  0,
		Reverse: false,
	}

	res, err := suite.queryClient.AllFiles(context.Background(), &types.QueryAllFiles{
		Pagination: &pg,
	})
	suite.Require().NoError(err)

	suite.Require().Equal(1, len(res.Files))

	mres, err := suite.queryClient.AllFilesByMerkle(context.Background(), &types.QueryAllFilesByMerkle{
		Pagination: &pg,
		Merkle:     merkle,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(1, len(mres.Files))

	ffres, err := suite.queryClient.FindFile(context.Background(), &types.QueryFindFile{Merkle: merkle})
	suite.Require().NoError(err)

	suite.Require().Equal(1, len(ffres.ProviderIps))

	suite.reset()
}
