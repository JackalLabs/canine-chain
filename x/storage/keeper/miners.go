package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/storage/types"
)

// SetMiners set a specific miners in the store from its index
func (k Keeper) SetMiners(ctx sdk.Context, miners types.Miners) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MinersKeyPrefix))
	b := k.cdc.MustMarshal(&miners)
	store.Set(types.MinersKey(
		miners.Address,
	), b)
}

// GetMiners returns a miners from its index
func (k Keeper) GetMiners(
	ctx sdk.Context,
	address string,

) (val types.Miners, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MinersKeyPrefix))

	b := store.Get(types.MinersKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMiners removes a miners from the store
func (k Keeper) RemoveMiners(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MinersKeyPrefix))
	store.Delete(types.MinersKey(
		address,
	))
}

// GetAllMiners returns all miners
func (k Keeper) GetAllMiners(ctx sdk.Context) (list []types.Miners) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MinersKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Miners
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
