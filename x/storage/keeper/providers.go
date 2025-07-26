package keeper

import (
	"net/url"
	"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
	"github.com/tendermint/tendermint/libs/rand"
)

const (
	Rounds = 20
)

// GetActiveProviders returns a list of recently active providers in a random order
func (k Keeper) GetActiveProviders(ctx sdk.Context, filterAddress string) []types.ActiveProviders {
	providers := k.GetAllActiveProviders(ctx)
	allowedProviders := make([]types.ActiveProviders, 0)

	filterDomain := ""
	filterTLD := ""

	if len(filterAddress) > 0 {
		url, err := url.Parse(filterAddress)
		if err != nil {
			return nil
		}

		parts := strings.Split(url.Hostname(), ".")
		partCount := len(parts)
		if partCount >= 2 {
			filterDomain = parts[partCount-2]
			filterTLD = parts[partCount-1]
		}
	}

	for _, provider := range providers {
		providerAccount, found := k.GetProviders(ctx, provider.Address)
		if !found {
			continue
		}
		url, err := url.Parse(providerAccount.Ip)
		if err != nil {
			continue
		}

		parts := strings.Split(url.Hostname(), ".")
		partCount := len(parts)
		if partCount < 2 {
			continue
		}

		domain := parts[partCount-2]
		tld := parts[partCount-1]

		if domain == filterDomain && tld == filterTLD {
			continue
		}

		allowedProviders = append(allowedProviders, provider)

	}
	providers = allowedProviders

	size := len(providers)

	rounds := Rounds * size

	i64Size := int64(size)

	r := rand.NewRand() // creating a new random generator to ensure no interference

	r.Seed(ctx.BlockHeight())

	for i := 0; i < rounds; i++ {
		x := r.Int63n(i64Size)
		y := r.Int63n(i64Size)

		providers[x], providers[y] = providers[y], providers[x]
	}

	return providers
}

// GetRandomizedProviders returns a list of providers in a random order
func (k Keeper) GetRandomizedProviders(ctx sdk.Context) []types.Providers {
	providers := k.GetAllProviders(ctx)
	size := len(providers)

	// Use Fisher-Yates algorithm - O(n) complexity
	r := rand.NewRand()
	r.Seed(ctx.BlockHeight())

	for i := size - 1; i > 0; i-- {
		j := r.Int63n(int64(i + 1))
		providers[i], providers[j] = providers[j], providers[i]
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
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProvidersKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Providers
		k.cdc.MustUnmarshal(iterator.Value(), &val)

		_, err := k.GetOneProofForProver(ctx, val.Address)
		if err == nil {
			list = append(list, types.ActiveProviders{Address: val.Address})
		}
	}

	return list
}
