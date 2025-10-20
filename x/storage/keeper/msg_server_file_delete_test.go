package keeper_test

import (
	testutil "github.com/jackalLabs/canine-chain/v4/testutil"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
)

func (suite *KeeperTestSuite) TestDeleteFile() {
	suite.SetupSuite()
	msgSrvr, k, ctx := setupMsgServer(suite)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 3)
	suite.Require().NoError(err)

	fileOwner := testAddresses[0]
	otherUser := testAddresses[1]
	depoAccount := testAddresses[2]

	// Create test file data
	merkle := []byte("test-merkle-hash")
	start := int64(100)
	fileSize := int64(1024)

	// Create a file that can be deleted (Expires = 0)
	deletableFile := types.UnifiedFile{
		Merkle:        merkle,
		Owner:         fileOwner,
		Start:         start,
		Expires:       0, // Can be deleted
		FileSize:      fileSize,
		ProofInterval: 50,
		ProofType:     0,
		Proofs:        []string{},
		MaxProofs:     3,
		Note:          "test file",
	}

	// Create a file that cannot be deleted (Expires != 0)
	nonDeletableFile := types.UnifiedFile{
		Merkle:        []byte("non-deletable-merkle"),
		Owner:         fileOwner,
		Start:         start + 1,
		Expires:       1000, // Cannot be deleted
		FileSize:      fileSize,
		ProofInterval: 50,
		ProofType:     0,
		Proofs:        []string{},
		MaxProofs:     3,
		Note:          "non-deletable file",
	}

	cases := []struct {
		testName     string
		msg          types.MsgDeleteFile
		setupFile    bool
		file         *types.UnifiedFile
		setupPayment bool
		expErr       bool
		expErrMsg    string
	}{
		{
			testName: "file not found",
			msg: types.MsgDeleteFile{
				Creator: fileOwner,
				Merkle:  merkle,
				Start:   start,
			},
			setupFile:    false,
			setupPayment: true,
			expErr:       true,
			expErrMsg:    "file not found",
		},
		{
			testName: "file not expired - cannot delete",
			msg: types.MsgDeleteFile{
				Creator: fileOwner,
				Merkle:  nonDeletableFile.Merkle,
				Start:   nonDeletableFile.Start,
			},
			setupFile:    true,
			file:         &nonDeletableFile,
			setupPayment: true,
			expErr:       true,
			expErrMsg:    "can not delete files before they expire",
		},
		{
			testName: "wrong owner - cannot delete",
			msg: types.MsgDeleteFile{
				Creator: otherUser, // Different user trying to delete
				Merkle:  merkle,
				Start:   start,
			},
			setupFile:    true,
			file:         &deletableFile,
			setupPayment: true,
			expErr:       true,
			expErrMsg:    "file not found", // File not found because owner doesn't match
		},
		{
			testName: "payment info not found",
			msg: types.MsgDeleteFile{
				Creator: fileOwner,
				Merkle:  merkle,
				Start:   start,
			},
			setupFile:    true,
			file:         &deletableFile,
			setupPayment: false,
			expErr:       true,
			expErrMsg:    "payment info not found",
		},
		{
			testName: "successful deletion",
			msg: types.MsgDeleteFile{
				Creator: fileOwner,
				Merkle:  merkle,
				Start:   start,
			},
			setupFile:    true,
			file:         &deletableFile,
			setupPayment: true,
			expErr:       false,
		},
	}

	for _, tcs := range cases {
		tc := tcs
		suite.Run(tc.testName, func() {
			// Reset the keeper state for each test
			suite.reset()
			msgSrvr, k, ctx = setupMsgServer(suite)

			// Reapply module params after reset
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

			// Set up the file if needed
			if tc.setupFile && tc.file != nil {
				k.SetFile(suite.ctx, *tc.file)
			}

			// Create fresh payment info for this test case
			var paymentInfo types.StoragePaymentInfo
			if tc.setupPayment {
				paymentInfo = types.StoragePaymentInfo{
					Start:          suite.ctx.BlockTime().AddDate(0, 0, -60),
					End:            suite.ctx.BlockTime().AddDate(0, 0, 30),
					SpaceAvailable: 100_000_000_000,
					SpaceUsed:      fileSize, // Start with some space used
					Address:        tc.msg.Creator,
				}
				k.SetStoragePaymentInfo(suite.ctx, paymentInfo)
			}

			// Execute the delete file message
			res, err := msgSrvr.DeleteFile(ctx, &tc.msg)

			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
				suite.Require().Nil(res)
			} else {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)

				// Verify the file was actually deleted
				_, found := k.GetFile(suite.ctx, tc.msg.Merkle, tc.msg.Creator, tc.msg.Start)
				suite.Require().False(found, "file should be deleted")

				// Verify payment info was updated correctly
				updatedPaymentInfo, found := k.GetStoragePaymentInfo(suite.ctx, tc.msg.Creator)
				suite.Require().True(found, "payment info should still exist")
				suite.Require().Equal(paymentInfo.SpaceUsed-fileSize, updatedPaymentInfo.SpaceUsed, "space used should be reduced by file size")
			}
		})
	}
}

func (suite *KeeperTestSuite) TestDeleteFileSpaceUsageEdgeCases() {
	suite.SetupSuite()
	msgSrvr, k, ctx := setupMsgServer(suite)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	suite.Require().NoError(err)

	fileOwner := testAddresses[0]
	depoAccount := testAddresses[1]

	merkle := []byte("test-merkle-hash")
	start := int64(100)
	fileSize := int64(1024)

	// Create a deletable file
	deletableFile := types.UnifiedFile{
		Merkle:        merkle,
		Owner:         fileOwner,
		Start:         start,
		Expires:       0, // Can be deleted
		FileSize:      fileSize,
		ProofInterval: 50,
		ProofType:     0,
		Proofs:        []string{},
		MaxProofs:     3,
		Note:          "test file",
	}

	cases := []struct {
		testName          string
		initialSpaceUsed  int64
		expectedSpaceUsed int64
	}{
		{
			testName:          "normal space usage reduction",
			initialSpaceUsed:  fileSize * 2,
			expectedSpaceUsed: fileSize, // Should reduce by fileSize
		},
		{
			testName:          "space usage goes to zero",
			initialSpaceUsed:  fileSize,
			expectedSpaceUsed: 0, // Should be capped at 0
		},
		{
			testName:          "space usage would go negative - capped at zero",
			initialSpaceUsed:  fileSize / 2,
			expectedSpaceUsed: 0, // Should be capped at 0
		},
	}

	for _, tcs := range cases {
		tc := tcs
		suite.Run(tc.testName, func() {
			// Reset the keeper state for each test
			suite.reset()
			msgSrvr, k, ctx = setupMsgServer(suite)

			// Reapply module params after reset
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

			// Set up the file
			k.SetFile(suite.ctx, deletableFile)

			// Set up payment info with specific space usage
			paymentInfo := types.StoragePaymentInfo{
				Start:          suite.ctx.BlockTime().AddDate(0, 0, -60),
				End:            suite.ctx.BlockTime().AddDate(0, 0, 30),
				SpaceAvailable: 100_000_000_000,
				SpaceUsed:      tc.initialSpaceUsed,
				Address:        fileOwner,
			}
			k.SetStoragePaymentInfo(suite.ctx, paymentInfo)

			// Execute the delete file message
			msg := types.MsgDeleteFile{
				Creator: fileOwner,
				Merkle:  merkle,
				Start:   start,
			}

			res, err := msgSrvr.DeleteFile(ctx, &msg)
			suite.Require().NoError(err)
			suite.Require().NotNil(res)

			// Verify payment info was updated correctly
			updatedPaymentInfo, found := k.GetStoragePaymentInfo(suite.ctx, fileOwner)
			suite.Require().True(found, "payment info should still exist")
			suite.Require().Equal(tc.expectedSpaceUsed, updatedPaymentInfo.SpaceUsed, "space used should match expected value")
		})
	}
}
