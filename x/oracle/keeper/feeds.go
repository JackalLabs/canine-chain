package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v5/x/oracle/types"
)

// SetFeed set a specific Feed in the store from its index
func (k Keeper) SetFeed(ctx sdk.Context, feed types.Feed) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeedKeyPrefix))
	f := k.cdc.MustMarshal(&feed)
	store.Set(types.FeedKey(feed.Name), f)
}

// GetFeed returns a Feed from its index
func (k Keeper) GetFeed(
	ctx sdk.Context,
	index string,
) (val types.Feed, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeedKeyPrefix))

	key := types.FeedKey(index)
	if !store.Has(key) {
		return types.Feed{}, false
	}

	k.cdc.MustUnmarshal(store.Get(key), &val)
	return val, true
}

// RemoveFeed removes a Feed from the store
func (k Keeper) RemoveFeed(
	ctx sdk.Context,
	index string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeedKeyPrefix))
	store.Delete(types.FeedKey(
		index,
	))
}

// GetAllFeeds returns all Feed
func (k Keeper) GetAllFeeds(ctx sdk.Context) (list []types.Feed) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeedKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Feed
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return list
}
