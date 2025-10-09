package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
)

func (k Keeper) SetRewardTracker(ctx sdk.Context, tracker types.RewardTracker) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RewardTrackerKeyPrefix))
	b := k.cdc.MustMarshal(&tracker)
	store.Set(types.RewardTrackerPrimaryKey(
		tracker.Provider,
	), b)
}

func (k Keeper) GetRewardTracker(
	ctx sdk.Context,
	provider string,
) (val types.RewardTracker, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RewardTrackerKeyPrefix))

	b := store.Get(types.RewardTrackerPrimaryKey(
		provider,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetAllRewardTrackers(ctx sdk.Context) (list []types.RewardTracker) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RewardTrackerKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RewardTracker
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
