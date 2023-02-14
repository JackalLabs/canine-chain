package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) KillOldContracts(ctx sdk.Context) {
	maxContractAgeInBlocks := k.GetParams(ctx).MaxContractAgeInBlocks
	contracts := k.GetAllContracts(ctx)

	for _, contract := range contracts {
		if contract.Age+maxContractAgeInBlocks < ctx.BlockHeight() {
			k.RemoveContracts(ctx, contract.Cid)
		}
	}
}
