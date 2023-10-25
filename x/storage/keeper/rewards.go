package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

func (k Keeper) manageProofs(ctx sdk.Context, sizeTracker *map[string]int64, file *types.UnifiedFile, proofKey string) {
	st := *sizeTracker

	proof, found := k.GetProofWithBuiltKey(ctx, []byte(proofKey))
	if !found {
		ctx.Logger().Info(fmt.Sprintf("cannot find proof: %s", proofKey))
		file.RemoveProverWithKey(ctx, k, proofKey)
	}

	currentHeight := ctx.BlockHeight()

	windowStart := currentHeight - file.ProofInterval

	if windowStart > proof.LastProven { // if the last time this file was proven was outside the proof window, burn their stake in the file
		ctx.Logger().Info(fmt.Sprintf("proof has not been proven within the last window: %d > %d", windowStart, proof.LastProven))
		file.RemoveProverWithKey(ctx, k, proofKey)
		return
	}

	st[proof.Prover] += file.FileSize
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
		if file.Start+file.ProofInterval < ctx.BlockHeight() {
			k.RemoveFile(ctx, file.Merkle, file.Owner, file.Start)
		}
	}
}

// ManageRewards loops through every file on the network and manages it in some way.
func (k Keeper) ManageRewards(ctx sdk.Context) {
	var totalSize int64
	s := make(map[string]int64)
	sizeTracker := &s

	k.IterateFilesByMerkle(ctx, false, func(key []byte, val []byte) bool {
		var file types.UnifiedFile
		k.cdc.MustUnmarshal(val, &file)

		s := file.FileSize * file.MaxProofs
		totalSize += s

		k.removeFileIfDeserved(ctx, &file) // delete file if it meets the conditions to be deleted

		for _, proof := range file.Proofs {
			k.manageProofs(ctx, sizeTracker, &file, proof)
		}

		return false
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
