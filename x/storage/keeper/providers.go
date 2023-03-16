package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

// SetProviders set a specific providers in the store from its index
func (k Keeper) SetProviders(ctx sdk.Context, providers types.Providers) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProvidersKeyPrefix))
	b := k.cdc.MustMarshal(&providers)
	store.Set(types.ProvidersKey(
		providers.Address,
	), b)
}

// GetProviders returns a providers from its index
func (k Keeper) GetProviders(
	ctx sdk.Context,
	address string,
) (val types.Providers, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProvidersKeyPrefix))

	b := store.Get(types.ProvidersKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveProviders removes a providers from the store
func (k Keeper) RemoveProviders(
	ctx sdk.Context,
	address string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProvidersKeyPrefix))
	store.Delete(types.ProvidersKey(
		address,
	))
}

// GetAllProviders returns all providers
func (k Keeper) GetAllProviders(ctx sdk.Context) (list []types.Providers) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProvidersKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Providers
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
