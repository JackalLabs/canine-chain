package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/jklmining/types"
)

// SetSaveRequests set a specific saveRequests in the store from its index
func (k Keeper) SetSaveRequests(ctx sdk.Context, saveRequests types.SaveRequests) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SaveRequestsKeyPrefix))
	b := k.cdc.MustMarshal(&saveRequests)
	store.Set(types.SaveRequestsKey(
		saveRequests.Index,
	), b)
}

// GetSaveRequests returns a saveRequests from its index
func (k Keeper) GetSaveRequests(
	ctx sdk.Context,
	index string,

) (val types.SaveRequests, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SaveRequestsKeyPrefix))

	b := store.Get(types.SaveRequestsKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSaveRequests removes a saveRequests from the store
func (k Keeper) RemoveSaveRequests(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SaveRequestsKeyPrefix))
	store.Delete(types.SaveRequestsKey(
		index,
	))
}

// GetAllSaveRequests returns all saveRequests
func (k Keeper) GetAllSaveRequests(ctx sdk.Context) (list []types.SaveRequests) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SaveRequestsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SaveRequests
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
