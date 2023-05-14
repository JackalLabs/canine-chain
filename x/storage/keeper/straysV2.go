package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

// SetStrays set a specific strays in the store from its index
func (k Keeper) SetStrays(ctx sdk.Context, strays types.StrayV2) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StraysV2KeyPrefix))
	b := k.cdc.MustMarshal(&strays)
	store.Set(types.StraysKey(
		strays.Cid,
	), b)
}

// GetStrays returns a strays from its index
func (k Keeper) GetStrays(
	ctx sdk.Context,
	cid string,
) (val types.StrayV2, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StraysV2KeyPrefix))

	b := store.Get(types.StraysKey(
		cid,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStrays removes a strays from the store
func (k Keeper) RemoveStrays(
	ctx sdk.Context,
	cid string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StraysV2KeyPrefix))
	store.Delete(types.StraysKey(
		cid,
	))
}

// GetAllStrays returns all strays
func (k Keeper) GetAllStrays(ctx sdk.Context) (list []types.StrayV2) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StraysV2KeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.StrayV2
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
