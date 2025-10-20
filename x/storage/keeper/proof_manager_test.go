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
