package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetMinerUsing(ctx sdk.Context, miner string) int64 {
	allDeals := k.GetAllActiveDeals(ctx)

	var space int64 = 0
	for i := 0; i < len(allDeals); i++ {
		deal := allDeals[i]
		if deal.Miner != miner {
			continue
		}
		size, ok := sdk.NewIntFromString(deal.Filesize)
		if !ok {
			continue
		}

		space += size.Int64()

	}

	return space
}
