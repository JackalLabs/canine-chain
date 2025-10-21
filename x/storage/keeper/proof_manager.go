package keeper

import (
	"fmt"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v5/x/storage/types"
)

// ManageProofs loops through every file on the network and manages it in some way.
func (k Keeper) ManageProofs(ctx sdk.Context) {
	expiredOwners := make(map[string]bool)
	ownerSize := make(map[string]int64)

	plans := k.GetAllStoragePaymentInfo(ctx)
	for _, plan := range plans {
		expired := plan.End.Before(ctx.BlockTime())
		if expired {
			expiredOwners[plan.Address] = true
		}
	}

	k.IterateAndParseFilesByMerkle(ctx, false, func(key []byte, file types.UnifiedFile) bool {
		if file.Expires == 0 { // for plans
			if expiredOwners[file.Owner] {
				k.RemoveFile(ctx, file.Merkle, file.Owner, file.Start) // remove if plan is expired
				return false
			}
			ownerSize[file.Owner] += file.FileSize
		} else if file.Expires <= ctx.BlockHeight() { // kill the file if it's too old
			k.RemoveFile(ctx, file.Merkle, file.Owner, file.Start) // for plain files
			return false
		}

		for _, proof := range file.Proofs { // run it through the proof remover
			k.manageProof(ctx, &file, proof)
		}

		return false
	})

	for _, plan := range plans { // fix all plan sizes
		size := ownerSize[plan.Address]
		if size != plan.SpaceUsed {
			plan.SpaceUsed = size
			k.SetStoragePaymentInfo(ctx, plan)
		}
	}
}

func (k Keeper) RunProofChecks(ctx sdk.Context) {
	DayBlocks := k.GetParams(ctx).ProofWindow // checks more often than proofs take to catch them more frequently

	if ctx.BlockHeight()%DayBlocks > 0 { // runs once each window (usually a full days worth of blocks)
		ctx.Logger().Debug("skipping reward handling for this block")
		return
	}

	k.ManageProofs(ctx)
}

// manageProof checks the status of a given proof, if the file is too young, we skip it. If it's old enough and the
// prover has either failed to prove it or the proof simply never existed we remove it.
func (k Keeper) manageProof(ctx sdk.Context, file *types.UnifiedFile, proofKey string) {
	providerAddress := strings.Split(proofKey, "/")[0]
	proof, found := k.GetProofWithBuiltKey(ctx, []byte(proofKey))

	currentHeight := ctx.BlockHeight()

	if !found {
		if !file.IsYoung(currentHeight) { // if the file is old, and we can't find the proof, remove the prover
			ctx.Logger().Info(fmt.Sprintf("cannot find proof: %s", proofKey))
			file.RemoveProverWithKey(ctx, k, proofKey)
		}
		return
	}

	proven := file.ProvenLastBlock(currentHeight, proof.LastProven)

	if !proven && !file.IsYoung(currentHeight) { // if file wasn't proven, and is old, we burn it.
		ctx.Logger().Info(fmt.Sprintf("proof has not been proven within the last window at %d | %s", currentHeight, proofKey))
		file.RemoveProverWithKey(ctx, k, proofKey)
		k.burnContract(ctx, providerAddress)
		return
	}
}

func (k Keeper) burnContract(ctx sdk.Context, providerAddress string) {
	prov, found := k.GetProviders(ctx, providerAddress)
	if !found {
		return
	}

	burned, err := strconv.Atoi(prov.BurnedContracts)
	if err != nil {
		ctx.Logger().Error("cannot parse providers burn count")
		return
	}

	prov.BurnedContracts = fmt.Sprintf("%d", burned+1)
	k.SetProviders(ctx, prov)
}
