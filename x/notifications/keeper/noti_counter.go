package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/notifications/types"
)

// SetNotiCounter set a specific notiCounter in the store from its index
func (k Keeper) SetNotiCounter(ctx sdk.Context, notiCounter types.NotiCounter) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NotiCounterKeyPrefix))
	b := k.cdc.MustMarshal(&notiCounter)
	store.Set(types.NotiCounterKey(
		notiCounter.Address,
	), b)
}

// GetNotiCounter returns a notiCounter from its index
func (k Keeper) GetNotiCounter(
	ctx sdk.Context,
	address string,
) (val types.NotiCounter, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NotiCounterKeyPrefix))

	b := store.Get(types.NotiCounterKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNotiCounter removes a notiCounter from the store
func (k Keeper) RemoveNotiCounter(
	ctx sdk.Context,
	address string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NotiCounterKeyPrefix))
	store.Delete(types.NotiCounterKey(
		address,
	))
}

// GetAllNotiCounter returns all notiCounter
func (k Keeper) GetAllNotiCounter(ctx sdk.Context) (list []types.NotiCounter) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NotiCounterKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NotiCounter
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
