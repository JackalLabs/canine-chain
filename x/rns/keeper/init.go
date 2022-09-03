package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/rns/types"
)

// SetInit set a specific init in the store from its index
func (k Keeper) SetInit(ctx sdk.Context, init types.Init) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InitKeyPrefix))
	b := k.cdc.MustMarshal(&init)
	store.Set(types.InitKey(
		init.Address,
	), b)
}

// GetInit returns a init from its index
func (k Keeper) GetInit(
	ctx sdk.Context,
	address string,

) (val types.Init, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InitKeyPrefix))

	b := store.Get(types.InitKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveInit removes a init from the store
func (k Keeper) RemoveInit(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InitKeyPrefix))
	store.Delete(types.InitKey(
		address,
	))
}

// GetAllInit returns all init
func (k Keeper) GetAllInit(ctx sdk.Context) (list []types.Init) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InitKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Init
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
