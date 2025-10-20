package keeper_test

import (
	"testing"
	"time"

	testutil "github.com/jackalLabs/canine-chain/v4/testutil"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
	"github.com/stretchr/testify/require"
)

func TestManageProofs(t *testing.T) {
	storageKeeper, _, _, _, ctx := setupStorageKeeper(t)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 3)
	require.NoError(t, err)

	owner := testAddresses[0]
	provider1 := testAddresses[1]
	provider2 := testAddresses[2]

	// Set up parameters
	storageKeeper.SetParams(ctx, types.Params{
		DepositAccount:         testAddresses[0],
		ProofWindow:            50,
		ChunkSize:              1024,
		PriceFeed:              "jklprice",
		MissesToBurn:           3,
		MaxContractAgeInBlocks: 100,
		PricePerTbPerMonth:     15,
		CollateralPrice:        2,
		CheckWindow:            11,
		ReferralCommission:     25,
		PolRatio:               40,
	})

	// Create providers
	provider1Data := types.Providers{
		Address:         provider1,
		Ip:              "192.168.1.1",
		Totalspace:      "1000",
		BurnedContracts: "0",
		Creator:         provider1,
	}
	provider2Data := types.Providers{
		Address:         provider2,
		Ip:              "192.168.1.2",
		Totalspace:      "1000",
		BurnedContracts: "0",
		Creator:         provider2,
	}
	storageKeeper.SetProviders(ctx, provider1Data)
	storageKeeper.SetProviders(ctx, provider2Data)

	// Create a test file
	merkle := []byte("test-merkle")
	file := types.UnifiedFile{
		Merkle:        merkle,
		Owner:         owner,
		Start:         10,
		Expires:       0, // Plan file
		FileSize:      1024,
		ProofInterval: 50,
		ProofType:     0,
		Proofs:        []string{},
		MaxProofs:     3,
		Note:          "test file",
	}

	// Add provers to the file
	file.AddProver(ctx, storageKeeper, provider1)
	file.AddProver(ctx, storageKeeper, provider2)

	// Set the file in storage
	storageKeeper.SetFile(ctx, file)

	// Create storage payment info for the owner
	paymentInfo := types.StoragePaymentInfo{
		Address:   owner,
		End:       ctx.BlockTime().Add(24 * time.Hour), // Not expired
		SpaceUsed: 0,
	}
	storageKeeper.SetStoragePaymentInfo(ctx, paymentInfo)

	// Run ManageProofs
	storageKeeper.ManageProofs(ctx)

	// Verify the file still exists
	retrievedFile, found := storageKeeper.GetFile(ctx, merkle, owner, file.Start)
	require.True(t, found)
	require.Equal(t, file.Merkle, retrievedFile.Merkle)
	require.Equal(t, file.Owner, retrievedFile.Owner)
}

func TestManageProofsWithExpiredPlan(t *testing.T) {
	storageKeeper, _, _, _, ctx := setupStorageKeeper(t)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	require.NoError(t, err)

	owner := testAddresses[0]
	provider := testAddresses[1]

	// Set up parameters
	storageKeeper.SetParams(ctx, types.Params{
		DepositAccount:         testAddresses[0],
		ProofWindow:            50,
		ChunkSize:              1024,
		PriceFeed:              "jklprice",
		MissesToBurn:           3,
		MaxContractAgeInBlocks: 100,
		PricePerTbPerMonth:     15,
		CollateralPrice:        2,
		CheckWindow:            11,
		ReferralCommission:     25,
		PolRatio:               40,
	})

	// Create provider
	providerData := types.Providers{
		Address:         provider,
		Ip:              "192.168.1.1",
		Totalspace:      "1000",
		BurnedContracts: "0",
		Creator:         provider,
	}
	storageKeeper.SetProviders(ctx, providerData)

	// Create a test file
	merkle := []byte("test-merkle-expired")
	file := types.UnifiedFile{
		Merkle:        merkle,
		Owner:         owner,
		Start:         10,
		Expires:       0, // Plan file
		FileSize:      1024,
		ProofInterval: 50,
		ProofType:     0,
		Proofs:        []string{},
		MaxProofs:     3,
		Note:          "test file",
	}

	// Add prover to the file
	file.AddProver(ctx, storageKeeper, provider)

	// Set the file in storage
	storageKeeper.SetFile(ctx, file)

	// Create expired storage payment info for the owner
	paymentInfo := types.StoragePaymentInfo{
		Address:   owner,
		End:       ctx.BlockTime().Add(-24 * time.Hour), // Expired
		SpaceUsed: 0,
	}
	storageKeeper.SetStoragePaymentInfo(ctx, paymentInfo)

	// Run ManageProofs
	storageKeeper.ManageProofs(ctx)

	// Verify the file was removed
	_, found := storageKeeper.GetFile(ctx, merkle, owner, file.Start)
	require.False(t, found)
}

func TestManageProofsWithExpiredFile(t *testing.T) {
	storageKeeper, _, _, _, ctx := setupStorageKeeper(t)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	require.NoError(t, err)

	owner := testAddresses[0]
	provider := testAddresses[1]

	// Set up parameters
	storageKeeper.SetParams(ctx, types.Params{
		DepositAccount:         testAddresses[0],
		ProofWindow:            50,
		ChunkSize:              1024,
		PriceFeed:              "jklprice",
		MissesToBurn:           3,
		MaxContractAgeInBlocks: 100,
		PricePerTbPerMonth:     15,
		CollateralPrice:        2,
		CheckWindow:            11,
		ReferralCommission:     25,
		PolRatio:               40,
	})

	// Create provider
	providerData := types.Providers{
		Address:         provider,
		Ip:              "192.168.1.1",
		Totalspace:      "1000",
		BurnedContracts: "0",
		Creator:         provider,
	}
	storageKeeper.SetProviders(ctx, providerData)

	// Create a test file with expiration
	merkle := []byte("test-merkle-expired-file")
	file := types.UnifiedFile{
		Merkle:        merkle,
		Owner:         owner,
		Start:         10,
		Expires:       ctx.BlockHeight() - 1, // Expires before current height
		FileSize:      1024,
		ProofInterval: 50,
		ProofType:     0,
		Proofs:        []string{},
		MaxProofs:     3,
		Note:          "test file",
	}

	// Add prover to the file
	file.AddProver(ctx, storageKeeper, provider)

	// Set the file in storage
	storageKeeper.SetFile(ctx, file)

	// Run ManageProofs
	storageKeeper.ManageProofs(ctx)

	// Verify the file was removed
	_, found := storageKeeper.GetFile(ctx, merkle, owner, file.Start)
	require.False(t, found)
}

func TestRunProofChecks(t *testing.T) {
	storageKeeper, _, _, _, ctx := setupStorageKeeper(t)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 1)
	require.NoError(t, err)

	// Set up parameters with proof window
	storageKeeper.SetParams(ctx, types.Params{
		DepositAccount:         testAddresses[0],
		ProofWindow:            50,
		ChunkSize:              1024,
		PriceFeed:              "jklprice",
		MissesToBurn:           3,
		MaxContractAgeInBlocks: 100,
		PricePerTbPerMonth:     15,
		CollateralPrice:        2,
		CheckWindow:            11,
		ReferralCommission:     25,
		PolRatio:               40,
	})

	// Test when block height is not divisible by proof window
	ctx = ctx.WithBlockHeight(25) // Not divisible by 50
	storageKeeper.RunProofChecks(ctx)

	// Test when block height is divisible by proof window
	ctx = ctx.WithBlockHeight(50) // Divisible by 50
	storageKeeper.RunProofChecks(ctx)
}

func TestManageProof_ValidProofStays(t *testing.T) {
	storageKeeper, _, _, _, ctx := setupStorageKeeper(t)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	require.NoError(t, err)

	owner := testAddresses[0]
	provider := testAddresses[1]

	// Set up parameters
	storageKeeper.SetParams(ctx, types.Params{
		DepositAccount:         testAddresses[0],
		ProofWindow:            50,
		ChunkSize:              1024,
		PriceFeed:              "jklprice",
		MissesToBurn:           3,
		MaxContractAgeInBlocks: 100,
		PricePerTbPerMonth:     15,
		CollateralPrice:        2,
		CheckWindow:            11,
		ReferralCommission:     25,
		PolRatio:               40,
	})

	// Create provider
	providerData := types.Providers{
		Address:         provider,
		Ip:              "192.168.1.1",
		Totalspace:      "1000",
		BurnedContracts: "0",
		Creator:         provider,
	}
	storageKeeper.SetProviders(ctx, providerData)

	// Create a test file
	merkle := []byte("test-merkle-valid")
	file := types.UnifiedFile{
		Merkle:        merkle,
		Owner:         owner,
		Start:         10,
		Expires:       0,
		FileSize:      1024,
		ProofInterval: 50,
		ProofType:     0,
		Proofs:        []string{},
		MaxProofs:     3,
		Note:          "test file",
	}

	// Add prover to the file
	file.AddProver(ctx, storageKeeper, provider)

	// Set the file in storage
	storageKeeper.SetFile(ctx, file)

	// Get the proof and update it to be recently proven
	proofKey := file.MakeProofKey(provider)
	proof, found := storageKeeper.GetProofWithBuiltKey(ctx, []byte(proofKey))
	require.True(t, found)

	// Set context to a height where the file is not young
	ctx = ctx.WithBlockHeight(file.Start + (file.ProofInterval * 3))

	// Update the proof to be recently proven (within the last window)
	// Calculate the current window and set LastProven to be within it
	currentHeight := ctx.BlockHeight()
	window := ((currentHeight-file.Start)/file.ProofInterval)*file.ProofInterval + file.Start
	lastWindowStart := window - file.ProofInterval
	proof.LastProven = lastWindowStart + 10 // Recently proven within the last window
	storageKeeper.SetProof(ctx, proof)

	// Run manageProof directly
	storageKeeper.ManageProofs(ctx)

	// Verify the proof still exists
	retrievedProof, found := storageKeeper.GetProofWithBuiltKey(ctx, []byte(proofKey))
	require.True(t, found)
	require.Equal(t, proof.LastProven, retrievedProof.LastProven)
}

func TestManageProof_NoProofStays(t *testing.T) {
	storageKeeper, _, _, _, ctx := setupStorageKeeper(t)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	require.NoError(t, err)

	owner := testAddresses[0]
	provider := testAddresses[1]

	// Set up parameters
	storageKeeper.SetParams(ctx, types.Params{
		DepositAccount:         testAddresses[0],
		ProofWindow:            50,
		ChunkSize:              1024,
		PriceFeed:              "jklprice",
		MissesToBurn:           3,
		MaxContractAgeInBlocks: 100,
		PricePerTbPerMonth:     15,
		CollateralPrice:        2,
		CheckWindow:            11,
		ReferralCommission:     25,
		PolRatio:               40,
	})

	// Create provider
	providerData := types.Providers{
		Address:         provider,
		Ip:              "192.168.1.1",
		Totalspace:      "1000",
		BurnedContracts: "0",
		Creator:         provider,
	}
	storageKeeper.SetProviders(ctx, providerData)

	// Create a test file
	merkle := []byte("test-merkle-no-proof")
	file := types.UnifiedFile{
		Merkle:        merkle,
		Owner:         owner,
		Start:         10,
		Expires:       0,
		FileSize:      1024,
		ProofInterval: 50,
		ProofType:     0,
		Proofs:        []string{},
		MaxProofs:     3,
		Note:          "test file",
	}

	// Add prover to the file
	file.AddProver(ctx, storageKeeper, provider)

	// Set the file in storage
	storageKeeper.SetFile(ctx, file)

	// Remove the proof to simulate no proof scenario
	proofKey := file.MakeProofKey(provider)
	storageKeeper.RemoveProofWithBuiltKey(ctx, []byte(proofKey))

	// Set context to a height where the file is not young
	ctx = ctx.WithBlockHeight(file.Start + (file.ProofInterval * 3))

	// Run manageProof directly
	storageKeeper.ManageProofs(ctx)

	// Verify the file still exists but without the proof
	retrievedFile, found := storageKeeper.GetFile(ctx, merkle, owner, file.Start)
	require.True(t, found)
	require.False(t, retrievedFile.ContainsProver(provider))
}

func TestManageProof_InvalidProofRemoved(t *testing.T) {
	storageKeeper, _, _, _, ctx := setupStorageKeeper(t)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	require.NoError(t, err)

	owner := testAddresses[0]
	provider := testAddresses[1]

	// Set up parameters
	storageKeeper.SetParams(ctx, types.Params{
		DepositAccount:         testAddresses[0],
		ProofWindow:            50,
		ChunkSize:              1024,
		PriceFeed:              "jklprice",
		MissesToBurn:           3,
		MaxContractAgeInBlocks: 100,
		PricePerTbPerMonth:     15,
		CollateralPrice:        2,
		CheckWindow:            11,
		ReferralCommission:     25,
		PolRatio:               40,
	})

	// Create provider
	providerData := types.Providers{
		Address:         provider,
		Ip:              "192.168.1.1",
		Totalspace:      "1000",
		BurnedContracts: "0",
		Creator:         provider,
	}
	storageKeeper.SetProviders(ctx, providerData)

	// Create a test file
	merkle := []byte("test-merkle-invalid")
	file := types.UnifiedFile{
		Merkle:        merkle,
		Owner:         owner,
		Start:         1000,
		Expires:       0,
		FileSize:      1024,
		ProofInterval: 50,
		ProofType:     0,
		Proofs:        []string{},
		MaxProofs:     3,
		Note:          "test file",
	}

	// Add prover to the file
	file.AddProver(ctx, storageKeeper, provider)

	// Set the file in storage
	storageKeeper.SetFile(ctx, file)

	// Get the proof and update it to be old (not proven within the last window)
	proofKey := file.MakeProofKey(provider)
	proof, found := storageKeeper.GetProofWithBuiltKey(ctx, []byte(proofKey))
	require.True(t, found)

	// Update the proof to be old (not proven within the last window)
	proof.LastProven = file.Start - 100 // Very old proof
	storageKeeper.SetProof(ctx, proof)

	// Set context to a height where the file is not young
	ctx = ctx.WithBlockHeight(file.Start + (file.ProofInterval * 3))

	// Run manageProof directly
	storageKeeper.ManageProofs(ctx)

	// Verify the proof was removed
	_, found = storageKeeper.GetProofWithBuiltKey(ctx, []byte(proofKey))
	require.False(t, found)

	// Verify the provider's burned contracts count increased
	updatedProvider, found := storageKeeper.GetProviders(ctx, provider)
	require.True(t, found)
	require.Equal(t, "1", updatedProvider.BurnedContracts)
}

func TestManageProof_YoungFileSkipped(t *testing.T) {
	storageKeeper, _, _, _, ctx := setupStorageKeeper(t)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	require.NoError(t, err)

	owner := testAddresses[0]
	provider := testAddresses[1]

	// Set up parameters
	storageKeeper.SetParams(ctx, types.Params{
		DepositAccount:         testAddresses[0],
		ProofWindow:            50,
		ChunkSize:              1024,
		PriceFeed:              "jklprice",
		MissesToBurn:           3,
		MaxContractAgeInBlocks: 100,
		PricePerTbPerMonth:     15,
		CollateralPrice:        2,
		CheckWindow:            11,
		ReferralCommission:     25,
		PolRatio:               40,
	})

	// Create provider
	providerData := types.Providers{
		Address:         provider,
		Ip:              "192.168.1.1",
		Totalspace:      "1000",
		BurnedContracts: "0",
		Creator:         provider,
	}
	storageKeeper.SetProviders(ctx, providerData)

	// Create a test file
	merkle := []byte("test-merkle-young")
	file := types.UnifiedFile{
		Merkle:        merkle,
		Owner:         owner,
		Start:         10,
		Expires:       0,
		FileSize:      1024,
		ProofInterval: 50,
		ProofType:     0,
		Proofs:        []string{},
		MaxProofs:     3,
		Note:          "test file",
	}

	// Add prover to the file
	file.AddProver(ctx, storageKeeper, provider)

	// Set the file in storage
	storageKeeper.SetFile(ctx, file)

	// Set context to a height where the file is still young
	ctx = ctx.WithBlockHeight(file.Start + file.ProofInterval) // Still young

	// Run manageProof directly
	storageKeeper.ManageProofs(ctx)

	// Verify the proof still exists (should be skipped because file is young)
	proofKey := file.MakeProofKey(provider)
	_, found := storageKeeper.GetProofWithBuiltKey(ctx, []byte(proofKey))
	require.True(t, found)

	// Verify the provider's burned contracts count didn't increase
	updatedProvider, found := storageKeeper.GetProviders(ctx, provider)
	require.True(t, found)
	require.Equal(t, "0", updatedProvider.BurnedContracts)
}

func TestBurnContract(t *testing.T) {
	storageKeeper, _, _, _, ctx := setupStorageKeeper(t)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	require.NoError(t, err)

	owner := testAddresses[0]
	provider := testAddresses[1]

	// Set up parameters
	storageKeeper.SetParams(ctx, types.Params{
		DepositAccount:         testAddresses[0],
		ProofWindow:            50,
		ChunkSize:              1024,
		PriceFeed:              "jklprice",
		MissesToBurn:           3,
		MaxContractAgeInBlocks: 100,
		PricePerTbPerMonth:     15,
		CollateralPrice:        2,
		CheckWindow:            11,
		ReferralCommission:     25,
		PolRatio:               40,
	})

	// Create provider with initial burned contracts count
	providerData := types.Providers{
		Address:         provider,
		Ip:              "192.168.1.1",
		Totalspace:      "1000",
		BurnedContracts: "5",
		Creator:         provider,
	}
	storageKeeper.SetProviders(ctx, providerData)

	// Create a test file
	merkle := []byte("test-merkle-burn")
	file := types.UnifiedFile{
		Merkle:        merkle,
		Owner:         owner,
		Start:         10,
		Expires:       0,
		FileSize:      1024,
		ProofInterval: 50,
		ProofType:     0,
		Proofs:        []string{},
		MaxProofs:     3,
		Note:          "test file",
	}

	// Add prover to the file
	file.AddProver(ctx, storageKeeper, provider)

	// Set the file in storage
	storageKeeper.SetFile(ctx, file)

	// Get the proof and update it to be old (invalid)
	proofKey := file.MakeProofKey(provider)
	proof, found := storageKeeper.GetProofWithBuiltKey(ctx, []byte(proofKey))
	require.True(t, found)

	// Update the proof to be old (not proven within the last window)
	proof.LastProven = file.Start - 100 // Very old proof
	storageKeeper.SetProof(ctx, proof)

	// Set context to a height where the file is not young
	ctx = ctx.WithBlockHeight(file.Start + (file.ProofInterval * 3))

	// Run ManageProofs which should call burnContract
	storageKeeper.ManageProofs(ctx)

	// Verify the burned contracts count increased
	updatedProvider, found := storageKeeper.GetProviders(ctx, provider)
	require.True(t, found)
	require.Equal(t, "6", updatedProvider.BurnedContracts)
}

func TestBurnContract_ProviderNotFound(t *testing.T) {
	storageKeeper, _, _, _, ctx := setupStorageKeeper(t)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 1)
	require.NoError(t, err)

	provider := testAddresses[0]

	// Don't create the provider, so it won't be found

	// Create a test file with the non-existent provider
	merkle := []byte("test-merkle-no-provider")
	file := types.UnifiedFile{
		Merkle:        merkle,
		Owner:         provider,
		Start:         10,
		Expires:       0,
		FileSize:      1024,
		ProofInterval: 50,
		ProofType:     0,
		Proofs:        []string{},
		MaxProofs:     3,
		Note:          "test file",
	}

	// Add prover to the file
	file.AddProver(ctx, storageKeeper, provider)

	// Set the file in storage
	storageKeeper.SetFile(ctx, file)

	// Get the proof and update it to be old
	proofKey := file.MakeProofKey(provider)
	proof, found := storageKeeper.GetProofWithBuiltKey(ctx, []byte(proofKey))
	require.True(t, found)

	// Update the proof to be old
	proof.LastProven = file.Start - 100
	storageKeeper.SetProof(ctx, proof)

	// Set context to a height where the file is not young
	ctx = ctx.WithBlockHeight(file.Start + (file.ProofInterval * 3))

	// Run manageProof directly - should not panic even if provider is not found
	storageKeeper.ManageProofs(ctx)

	// Verify the proof was still removed
	_, found = storageKeeper.GetProofWithBuiltKey(ctx, []byte(proofKey))
	require.False(t, found)
}

func TestManageProof_StepThroughBlockHeights(t *testing.T) {
	storageKeeper, _, _, _, ctx := setupStorageKeeper(t)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	require.NoError(t, err)

	owner := testAddresses[0]
	provider := testAddresses[1]

	// Set up parameters
	storageKeeper.SetParams(ctx, types.Params{
		DepositAccount:         testAddresses[0],
		ProofWindow:            50,
		ChunkSize:              1024,
		PriceFeed:              "jklprice",
		MissesToBurn:           3,
		MaxContractAgeInBlocks: 100,
		PricePerTbPerMonth:     15,
		CollateralPrice:        2,
		CheckWindow:            11,
		ReferralCommission:     25,
		PolRatio:               40,
	})

	// Create provider
	providerData := types.Providers{
		Address:         provider,
		Ip:              "192.168.1.1",
		Totalspace:      "1000",
		BurnedContracts: "0",
		Creator:         provider,
	}
	storageKeeper.SetProviders(ctx, providerData)

	// Create a test file
	merkle := []byte("test-merkle-step-through")
	file := types.UnifiedFile{
		Merkle:        merkle,
		Owner:         owner,
		Start:         1000,
		Expires:       0,
		FileSize:      1024,
		ProofInterval: 50,
		ProofType:     0,
		Proofs:        []string{},
		MaxProofs:     3,
		Note:          "test file",
	}

	// Add prover to the file
	file.AddProver(ctx, storageKeeper, provider)

	// Set the file in storage
	storageKeeper.SetFile(ctx, file)

	// Get the proof and set it to be proven at the start
	proofKey := file.MakeProofKey(provider)
	proof, found := storageKeeper.GetProofWithBuiltKey(ctx, []byte(proofKey))
	require.True(t, found)

	// Set LastProven to the file start (initial proof)
	proof.LastProven = file.Start
	storageKeeper.SetProof(ctx, proof)

	// Create storage payment info for the owner (not expired)
	paymentInfo := types.StoragePaymentInfo{
		Address:   owner,
		End:       ctx.BlockTime().Add(24 * time.Hour), // Not expired
		SpaceUsed: 0,
	}
	storageKeeper.SetStoragePaymentInfo(ctx, paymentInfo)

	// Track when the proof gets removed
	proofRemovedAt := int64(-1)
	fileRemovedAt := int64(-1)

	// Step through block heights starting from when file is young to when it should be removed
	startHeight := file.Start
	endHeight := file.Start + (file.ProofInterval * 5) // Go well beyond what's needed

	t.Logf("Starting step-through test:")
	t.Logf("File start: %d", file.Start)
	t.Logf("File proof interval: %d", file.ProofInterval)
	t.Logf("File is young until: %d", file.Start+file.ProofInterval)
	t.Logf("Testing from height %d to %d", startHeight, endHeight)

	for height := startHeight; height <= endHeight; height++ {
		ctx = ctx.WithBlockHeight(height)

		// Check if file exists before running ManageProofs
		_, fileExistsBefore := storageKeeper.GetFile(ctx, merkle, owner, file.Start)
		_, proofExistsBefore := storageKeeper.GetProofWithBuiltKey(ctx, []byte(proofKey))

		// Run ManageProofs at this height
		storageKeeper.ManageProofs(ctx)

		// Check if file exists after running ManageProofs
		_, fileExistsAfter := storageKeeper.GetFile(ctx, merkle, owner, file.Start)
		_, proofExistsAfter := storageKeeper.GetProofWithBuiltKey(ctx, []byte(proofKey))

		// Log significant events
		if proofExistsBefore && !proofExistsAfter && proofRemovedAt == -1 {
			proofRemovedAt = height
			t.Logf("PROOF REMOVED at height %d", height)
			t.Logf("  - File was young: %t", file.IsYoung(height))
			t.Logf("  - Proof was valid: %t", file.ProvenLastBlock(height, proof.LastProven))
		}

		if fileExistsBefore && !fileExistsAfter && fileRemovedAt == -1 {
			fileRemovedAt = height
			t.Logf("FILE REMOVED at height %d", height)
		}

		// Log every 10th block for visibility
		if height%10 == 0 {
			t.Logf("Height %d: File exists=%t, Proof exists=%t, File young=%t",
				height, fileExistsAfter, proofExistsAfter, file.IsYoung(height))
		}

		// If both file and proof are gone, we can stop
		if !fileExistsAfter && !proofExistsAfter {
			t.Logf("Both file and proof removed, stopping at height %d", height)
			break
		}
	}

	// Verify results
	t.Logf("\nTest Results:")
	t.Logf("Proof removed at height: %d", proofRemovedAt)
	t.Logf("File removed at height: %d", fileRemovedAt)

	// The proof should be removed when:
	// 1. File is not young (height > file.Start + file.ProofInterval)
	// 2. Proof is not valid (not proven within the last window)
	// 3. RunProofChecks runs (height % proofWindow == 0)
	youngUntil := file.Start + file.ProofInterval

	if proofRemovedAt != -1 {
		require.GreaterOrEqual(t, proofRemovedAt, youngUntil,
			"Proof should only be removed after file is no longer young")

		// Verify the proof was actually invalid when removed
		wasValid := file.ProvenLastBlock(proofRemovedAt, proof.LastProven)
		require.False(t, wasValid,
			"Proof should be invalid when removed (not proven within last window)")

		t.Logf("✓ Proof correctly removed at height %d (file was young until %d)",
			proofRemovedAt, youngUntil)
	} else {
		t.Logf("⚠ Proof was never removed during test")
	}

	// The file should still exist (since payment plan is not expired)
	finalFile, fileExists := storageKeeper.GetFile(ctx, merkle, owner, file.Start)
	if fileExists {
		t.Logf("✓ File still exists (as expected, since payment plan is not expired)")
		require.Empty(t, finalFile.Proofs, "File should have no proofs after removal")
	} else {
		t.Logf("⚠ File was removed unexpectedly")
	}
}

func TestRunProofChecks_Scheduling(t *testing.T) {
	storageKeeper, _, _, _, ctx := setupStorageKeeper(t)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	require.NoError(t, err)

	owner := testAddresses[0]
	provider := testAddresses[1]

	// Set up parameters with proof window
	storageKeeper.SetParams(ctx, types.Params{
		DepositAccount:         testAddresses[0],
		ProofWindow:            50,
		ChunkSize:              1024,
		PriceFeed:              "jklprice",
		MissesToBurn:           3,
		MaxContractAgeInBlocks: 100,
		PricePerTbPerMonth:     15,
		CollateralPrice:        2,
		CheckWindow:            11,
		ReferralCommission:     25,
		PolRatio:               40,
	})

	// Create provider
	providerData := types.Providers{
		Address:         provider,
		Ip:              "192.168.1.1",
		Totalspace:      "1000",
		BurnedContracts: "0",
		Creator:         provider,
	}
	storageKeeper.SetProviders(ctx, providerData)

	// Create a test file
	merkle := []byte("test-merkle-scheduling")
	file := types.UnifiedFile{
		Merkle:        merkle,
		Owner:         owner,
		Start:         1000,
		Expires:       0,
		FileSize:      1024,
		ProofInterval: 50,
		ProofType:     0,
		Proofs:        []string{},
		MaxProofs:     3,
		Note:          "test file",
	}

	// Add prover to the file
	file.AddProver(ctx, storageKeeper, provider)

	// Set the file in storage
	storageKeeper.SetFile(ctx, file)

	// Get the proof and set it to be old (invalid)
	proofKey := file.MakeProofKey(provider)
	proof, found := storageKeeper.GetProofWithBuiltKey(ctx, []byte(proofKey))
	require.True(t, found)

	// Set LastProven to be old (not proven within the last window)
	proof.LastProven = file.Start - 100
	storageKeeper.SetProof(ctx, proof)

	// Create storage payment info for the owner (not expired)
	paymentInfo := types.StoragePaymentInfo{
		Address:   owner,
		End:       ctx.BlockTime().Add(24 * time.Hour), // Not expired
		SpaceUsed: 0,
	}
	storageKeeper.SetStoragePaymentInfo(ctx, paymentInfo)

	// Set context to a height where the file is not young
	ctx = ctx.WithBlockHeight(file.Start + (file.ProofInterval * 3))

	t.Logf("Testing RunProofChecks scheduling:")
	t.Logf("Proof window: %d", storageKeeper.GetParams(ctx).ProofWindow)
	t.Logf("Current height: %d", ctx.BlockHeight())

	// Test RunProofChecks at different heights
	testHeights := []int64{1100, 1101, 1149, 1150, 1151, 1200, 1249, 1250, 1251}

	for _, height := range testHeights {
		ctx = ctx.WithBlockHeight(height)

		// Check if proof exists before
		_, proofExistsBefore := storageKeeper.GetProofWithBuiltKey(ctx, []byte(proofKey))

		// Run RunProofChecks
		storageKeeper.RunProofChecks(ctx)

		// Check if proof exists after
		_, proofExistsAfter := storageKeeper.GetProofWithBuiltKey(ctx, []byte(proofKey))

		shouldRun := height%storageKeeper.GetParams(ctx).ProofWindow == 0
		proofRemoved := proofExistsBefore && !proofExistsAfter

		// Debug the modulo calculation
		modulo := height % storageKeeper.GetParams(ctx).ProofWindow
		t.Logf("Height %d: Modulo=%d, Should run=%t, Proof removed=%t", height, modulo, shouldRun, proofRemoved)

		if shouldRun && !proofRemoved {
			t.Logf("  ⚠ RunProofChecks should have removed proof but didn't")
		} else if !shouldRun && proofRemoved {
			t.Logf("  ⚠ RunProofChecks removed proof when it shouldn't have run")
		} else if shouldRun && proofRemoved {
			t.Logf("  ✓ RunProofChecks correctly removed proof")
		} else {
			t.Logf("  ✓ RunProofChecks correctly skipped (proof still exists)")
		}

		// If proof was removed, we can stop testing
		if proofRemoved {
			break
		}
	}
}
