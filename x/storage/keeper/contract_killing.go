package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const MaxContractAgeInBlocks = 100

func (k Keeper) KillOldContracts(ctx sdk.Context) {
	contracts := k.GetAllContracts(ctx)

	for _, contract := range contracts {
		if contract.Age+MaxContractAgeInBlocks < ctx.BlockHeight() {
			k.RemoveContracts(ctx, contract.Cid)
		}
	}
}
