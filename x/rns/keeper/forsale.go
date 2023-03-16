package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/rns/types"
)

// SetForsale set a specific forsale in the store from its index
func (k Keeper) SetForsale(ctx sdk.Context, forsale types.Forsale) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ForsaleKeyPrefix))
	b := k.cdc.MustMarshal(&forsale)
	store.Set(types.ForsaleKey(
		forsale.Name,
	), b)
}

// GetForsale returns a forsale from its index
func (k Keeper) GetForsale(
	ctx sdk.Context,
	name string,
) (val types.Forsale, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ForsaleKeyPrefix))

	b := store.Get(types.ForsaleKey(
		name,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveForsale removes a forsale from the store
func (k Keeper) RemoveForsale(
	ctx sdk.Context,
	name string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ForsaleKeyPrefix))
	store.Delete(types.ForsaleKey(
		name,
	))
}

// GetAllForsale returns all forsale
func (k Keeper) GetAllForsale(ctx sdk.Context) (list []types.Forsale) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ForsaleKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Forsale
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
