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

// manageProof checks the status of a given proof, if the file is too young, we skip it. If it's old enough and the
// prover has either failed to prove it or the proof simply never existed we remove it.
func (k Keeper) manageProof(ctx sdk.Context, sizeTracker *map[string]int64, file *types.UnifiedFile, proofKey string) {
	st := *sizeTracker

	pks := strings.Split(proofKey, "/")
	providerAddress := pks[0]

	proof, found := k.GetProofWithBuiltKey(ctx, []byte(proofKey))
	// If we check the file and there is a proof delegated but the provider hasn't proven it yet we remove it.
	// However, we need to check if the file is new and is being caught by accident
	if !file.IsYoung(ctx.BlockHeight()) { // give first window grace before removal
		if !found {
			ctx.Logger().Info(fmt.Sprintf("cannot find proof: %s", proofKey))
			file.RemoveProverWithKey(ctx, k, proofKey)
			return
		}

		currentHeight := ctx.BlockHeight()

		proven := file.ProvenLastBlock(currentHeight, proof.LastProven)

		if !proven { // if file has not been proven yet
			ctx.Logger().Info(fmt.Sprintf("proof has not been proven within the last window at %d", currentHeight))
			file.RemoveProverWithKey(ctx, k, proofKey)
			k.burnContract(ctx, providerAddress)
			return
		}

		st[proof.Prover] += file.FileSize // only give rewards to providers who have held onto the file for a full window
	}
}

func (k Keeper) rewardProviders(ctx sdk.Context, totalSize int64, sizeTracker *map[string]int64) {
	networkValue := sdk.NewDec(totalSize)

	storageWallet, err := types.GetTokenHolderAccount()
	if err != nil {
		ctx.Logger().Error(err.Error())
		return
	}

	currentTime := ctx.BlockTime()

	totalCoins := make(sdk.Coins, 0)

	k.IterateGauges(ctx, func(pg types.PaymentGauge) {
		if pg.End.Before(currentTime) {
			k.RemoveGauge(ctx, pg.Id)
			return
		}

		if pg.End.Before(pg.Start) {
			k.RemoveGauge(ctx, pg.Id)
			return
		}

		if pg.End.Equal(pg.Start) {
			k.RemoveGauge(ctx, pg.Id)
			return
		}

		totalTime := pg.End.Sub(pg.Start)
		ttM := totalTime.Microseconds()

		used := currentTime.Sub(pg.Start)
		uM := used.Microseconds()

		totalDec := sdk.NewDec(ttM)
		usedDec := sdk.NewDec(uM)

		usedRatio := usedDec.Quo(totalDec)
		uss := usedRatio.String()
		_ = uss
		coinsToAdd := pg.Coins
		for _, coin := range coinsToAdd {
			newAmt := coin.Amount.ToDec().Mul(usedRatio).TruncateInt()
			c := sdk.NewCoin(coin.Denom, newAmt)

			totalCoins = totalCoins.Add(c)
		}
	})

	for prover, worth := range *sizeTracker {

		providerValue := sdk.NewDec(worth)

		networkPercentage := providerValue.Quo(networkValue)

		coins := make(sdk.Coins, 0)

		for _, coin := range totalCoins {
			tokensValueOwed := networkPercentage.Mul(coin.Amount.ToDec()).TruncateInt()
			c := sdk.NewCoin(coin.Denom, tokensValueOwed)
			coins = coins.Add(c)
		}

		pAddress, err := sdk.AccAddressFromBech32(prover)
		if err != nil {
			ctx.Logger().Error(sdkerrors.Wrapf(err, "failed to convert prover address %s to bech32", prover).Error())
			continue
		}
		err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, storageWallet, types.ModuleName, coins)
		if err != nil {
			ctx.Logger().Error(sdkerrors.Wrapf(err, "failed to send %s to %s", coins.String(), types.ModuleName).Error())
			continue
		}
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, pAddress, coins)
		if err != nil {
			ctx.Logger().Error(sdkerrors.Wrapf(err, "failed to send %s to %s", coins.String(), prover).Error())
			continue
		}
	}
}

func (k Keeper) removeFileIfDeserved(ctx sdk.Context, file *types.UnifiedFile) {
	if len(file.Proofs) == 0 { // remove file if it
		if !file.IsYoung(ctx.BlockHeight()) { // give first window grace
			k.RemoveFile(ctx, file.Merkle, file.Owner, file.Start)
		}
	}
}

// ManageRewards loops through every file on the network and manages it in some way.
func (k Keeper) ManageRewards(ctx sdk.Context) {
	var totalSize int64
	s := make(map[string]int64)
	sizeTracker := &s

	k.IterateFilesByMerkle(ctx, false, func(_ []byte, val []byte) bool {
		var file types.UnifiedFile
		k.cdc.MustUnmarshal(val, &file)

		s := file.FileSize * int64(len(file.Proofs))
		totalSize += s

		k.removeFileIfDeserved(ctx, &file) // delete file if it meets the conditions to be deleted

		for _, proof := range file.Proofs { // manage all proofs in proof list
			k.manageProof(ctx, sizeTracker, &file, proof)
		}

		return false
	})

	k.rewardProviders(ctx, totalSize, sizeTracker)
}

func (k Keeper) RunRewardBlock(ctx sdk.Context) {
	DayBlocks := k.GetParams(ctx).CheckWindow // checks more often than proofs take to catch them more frequently

	if ctx.BlockHeight()%DayBlocks > 0 { // runs once each window (usually a full days worth of blocks)
		ctx.Logger().Debug("skipping reward handling for this block")
		return
	}

	k.ManageRewards(ctx)
}
