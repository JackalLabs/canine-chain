package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

// SetActiveDeals set a specific activeDeals in the store from its index
func (k Keeper) SetActiveDeals(ctx sdk.Context, activeDeals types.ActiveDealsV2) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveDealsV2KeyPrefix))
	b := k.cdc.MustMarshal(&activeDeals)
	store.Set(types.ActiveDealsKey(
		activeDeals.Cid,
	), b)
}

// GetActiveDeals returns a activeDeals from its index
func (k Keeper) GetActiveDeals(
	ctx sdk.Context,
	cid string,
) (val types.ActiveDealsV2, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveDealsV2KeyPrefix))

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
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveDealsV2KeyPrefix))
	store.Delete(types.ActiveDealsKey(
		cid,
	))
}

// GetAllActiveDeals returns all activeDeals
func (k Keeper) GetAllActiveDeals(ctx sdk.Context) (list []types.ActiveDealsV2) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveDealsV2KeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ActiveDealsV2
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
