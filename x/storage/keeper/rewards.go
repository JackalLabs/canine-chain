package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

func (k Keeper) manageProofs(ctx sdk.Context, sizeTracker *map[string]int64, file *types.UnifiedFile, prover string) {
	st := *sizeTracker

	proof, found := k.GetProof(ctx, prover, file.Merkle, file.Owner, file.Start)
	if !found {
		file.RemoveProver(ctx, k, prover)
	}

	currentHeight := ctx.BlockHeight()

	windowStart := currentHeight - file.ProofInterval

	if windowStart > proof.LastProven { // if the last time this file was proven was outside the proof window, burn their stake in the file
		file.RemoveProver(ctx, k, prover)
		return
	}

	st[prover] += file.FileSize
}

func (k Keeper) rewardProviders(totalSize int64, sizeTracker *map[string]int64) {
	networkValue := sdk.NewDec(totalSize)

	providersVault := sdk.NewDec(8000000) // TODO: Change this to the actual amount of tokens in the vault

	for prover, worth := range *sizeTracker {

		providerValue := sdk.NewDec(worth)

		networkPercentage := providerValue.Quo(networkValue)

		tokensValueOwed := networkPercentage.Mul(providersVault).TruncateInt64()

		_ = tokensValueOwed // TODO: send actual tokens from some vault address to the providers
		_ = prover
	}
}

func (k Keeper) removeFileIfDeserved(ctx sdk.Context, file *types.UnifiedFile) {
	if len(file.Proofs) == 0 { // remove file if it
		k.RemoveFile(ctx, file.Merkle, file.Owner, file.Start)
	}
}

// ManageRewards loops through every file on the network and manages it in some way.
func (k Keeper) ManageRewards(ctx sdk.Context) {
	var totalSize int64
	s := make(map[string]int64)
	sizeTracker := &s

	k.IterateFilesByMerkle(ctx, func(key []byte, val []byte) {
		var file types.UnifiedFile
		k.cdc.MustUnmarshal(val, &file)

		s := file.FileSize * file.MaxProofs
		totalSize += s

		k.removeFileIfDeserved(ctx, &file) // delete file if it meets the conditions to be deleted

		for _, proof := range file.Proofs {
			k.manageProofs(ctx, sizeTracker, &file, proof)
		}
	})

	k.rewardProviders(totalSize, sizeTracker)
}

func (k Keeper) RunRewardBlock(ctx sdk.Context) {
	DayBlocks := k.GetParams(ctx).ProofWindow // TODO: Change this window to 14400

	if ctx.BlockHeight()%DayBlocks > 0 { // runs once each window (usually a full days worth of blocks)
		ctx.Logger().Debug("skipping reward handling for this block")
		return
	}

	k.ManageRewards(ctx)
}
