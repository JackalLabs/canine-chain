package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

// SetClientUsage set a specific clientUsage in the store from its index
func (k Keeper) SetClientUsage(ctx sdk.Context, clientUsage types.ClientUsage) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClientUsageKeyPrefix))
	b := k.cdc.MustMarshal(&clientUsage)
	store.Set(types.ClientUsageKey(
		clientUsage.Address,
	), b)
}

// GetClientUsage returns a clientUsage from its index
func (k Keeper) GetClientUsage(
	ctx sdk.Context,
	address string,
) (val types.ClientUsage, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClientUsageKeyPrefix))

	b := store.Get(types.ClientUsageKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveClientUsage removes a clientUsage from the store
func (k Keeper) RemoveClientUsage(
	ctx sdk.Context,
	address string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClientUsageKeyPrefix))
	store.Delete(types.ClientUsageKey(
		address,
	))
}

// GetAllClientUsage returns all clientUsage
func (k Keeper) GetAllClientUsage(ctx sdk.Context) (list []types.ClientUsage) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClientUsageKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ClientUsage
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
