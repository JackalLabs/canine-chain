package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
	"github.com/tendermint/tendermint/libs/rand"
)

const (
	Rounds = 20
)

// GetActiveProviders returns a list of recently active providers in a random order
func (k Keeper) GetActiveProviders(ctx sdk.Context) []types.ActiveProviders {
	providers := k.GetAllActiveProviders(ctx)

	size := len(providers)

	rounds := Rounds * size

	i64Size := int64(size)

	rand.Seed(ctx.BlockTime().UnixNano())

	for i := 0; i < rounds; i++ {
		x := rand.Int63n(i64Size)
		y := rand.Int63n(i64Size)

		providers[x], providers[y] = providers[y], providers[x]
	}

	return providers
}

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

// SetActiveProviders set a specific providers in the store from its index
func (k Keeper) SetActiveProviders(ctx sdk.Context, providers types.ActiveProviders) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveProvidersKeyPrefix))
	b := k.cdc.MustMarshal(&providers)
	store.Set(types.ActiveProvidersKey(
		providers.Address,
	), b)
}

// RemoveAllActiveProviders removes all active providers
func (k Keeper) RemoveAllActiveProviders(
	ctx sdk.Context,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveProvidersKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ActiveProviders
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		store.Delete(types.ActiveProvidersKey(
			val.Address,
		))
	}
}

// GetAllActiveProviders returns all providers
func (k Keeper) GetAllActiveProviders(ctx sdk.Context) (list []types.ActiveProviders) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveProvidersKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ActiveProviders
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return list
}
