package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

// SetActiveDeals set a specific activeDeals in the store from its index
func (k Keeper) SetActiveDeals(ctx sdk.Context, activeDeals types.ActiveDeals) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveDealsKeyPrefix))
	b := k.cdc.MustMarshal(&activeDeals)
	store.Set(types.ActiveDealsKey(
		activeDeals.Cid,
	), b)
}

// GetActiveDeals returns a activeDeals from its index
func (k Keeper) GetActiveDeals(
	ctx sdk.Context,
	cid string,
) (val types.ActiveDeals, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveDealsKeyPrefix))

	b := store.Get(types.ActiveDealsKey(
		cid,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveActiveDeals removes a activeDeals from the store
func (k Keeper) RemoveActiveDeals(
	ctx sdk.Context,
	cid string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveDealsKeyPrefix))
	store.Delete(types.ActiveDealsKey(
		cid,
	))
}

// GetAllActiveDeals returns all activeDeals
func (k Keeper) GetAllActiveDeals(ctx sdk.Context) (list []types.ActiveDeals) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveDealsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ActiveDeals
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// IterateActiveDeals runs `fn` for each active deal in the store
func (k Keeper) IterateActiveDeals(ctx sdk.Context, fn func(deal types.ActiveDeals) bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveDealsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ActiveDeals
		k.cdc.MustUnmarshal(iterator.Value(), &val)

		if fn(val) {
			return
		}

	}
}
