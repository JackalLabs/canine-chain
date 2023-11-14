package keeper

import (
	"fmt"
	"strconv"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

func (k Keeper) burnContract(ctx sdk.Context, providerAddress string) {
	prov, found := k.GetProviders(ctx, providerAddress)
	if !found {
		return
	}

	burned, err := strconv.ParseInt(prov.BurnedContracts, 10, 64)
	if err != nil {
		ctx.Logger().Error("cannot parse providers burn count")
		return
	}

	prov.BurnedContracts = fmt.Sprintf("%d", burned+1)
	k.SetProviders(ctx, prov)
}

func (k Keeper) manageProofs(ctx sdk.Context, sizeTracker *map[string]int64, file *types.UnifiedFile, proofKey string) {
	st := *sizeTracker

	pks := strings.Split(proofKey, "/")
	providerAddress := pks[0]

	proof, found := k.GetProofWithBuiltKey(ctx, []byte(proofKey))
	if !found {
		ctx.Logger().Info(fmt.Sprintf("cannot find proof: %s", proofKey))
		file.RemoveProverWithKey(ctx, k, proofKey)
		return
	}

	currentHeight := ctx.BlockHeight()

	proven := file.Proven(ctx, k, currentHeight, providerAddress)

	if !proven { // if file has not been proven yet
		ctx.Logger().Info(fmt.Sprintf("proof has not been proven within the last window at %d", currentHeight))
		file.RemoveProverWithKey(ctx, k, proofKey)
		k.burnContract(ctx, providerAddress)
		return
	}

	st[proof.Prover] += file.FileSize
}

// TODO: Completely change the way this is done in Econ v2
func (k Keeper) rewardProviders(ctx sdk.Context, totalSize int64, sizeTracker *map[string]int64) {
	networkValue := sdk.NewDec(totalSize)

	storageWallet := k.accountkeeper.GetModuleAddress(types.ModuleName)

	tokens := k.bankkeeper.GetBalance(ctx, storageWallet, "ujkl")
	tokenAmountDec := tokens.Amount.ToDec()

	for prover, worth := range *sizeTracker {

		providerValue := sdk.NewDec(worth)

		networkPercentage := providerValue.Quo(networkValue)

		tokensValueOwed := networkPercentage.Mul(tokenAmountDec).TruncateInt64()

		coin := sdk.NewInt64Coin("ujkl", tokensValueOwed)
		coins := sdk.NewCoins(coin)

		pAddress, err := sdk.AccAddressFromBech32(prover)
		if err != nil {
			ctx.Logger().Error(sdkerrors.Wrapf(err, "failed to convert prover address %s to bech32", prover).Error())
			continue
		}
		err = k.bankkeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, pAddress, coins)
		if err != nil {
			ctx.Logger().Error(sdkerrors.Wrapf(err, "failed to send %d tokens to %s", tokensValueOwed, prover).Error())
			continue
		}
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

		s := file.FileSize * int64(len(file.Proofs))
		totalSize += s

		k.removeFileIfDeserved(ctx, &file) // delete file if it meets the conditions to be deleted

		for _, proof := range file.Proofs {
			k.manageProofs(ctx, sizeTracker, &file, proof)
		}

		return false
	})

	k.rewardProviders(ctx, totalSize, sizeTracker)
}

func (k Keeper) RunRewardBlock(ctx sdk.Context) {
	DayBlocks := k.GetParams(ctx).ProofWindow // TODO: Change this window to 14400

	if ctx.BlockHeight()%DayBlocks > 0 { // runs once each window (usually a full days worth of blocks)
		ctx.Logger().Debug("skipping reward handling for this block")
		return
	}

	k.ManageRewards(ctx)
}
