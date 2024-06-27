//nolint:all
package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
)

// SetActiveDeals set a specific activeDeals in the store from its index
func (k Keeper) SetLegacyActiveDeals(ctx sdk.Context, activeDeals types.LegacyActiveDeals) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LegacyActiveDealsKeyPrefix))
	b := k.cdc.MustMarshal(&activeDeals)
	store.Set(types.LegacyActiveDealsKey(
		activeDeals.Cid,
	), b)
}

// GetActiveDeals returns a activeDeals from its index
func (k Keeper) GetLegacyActiveDeals(
	ctx sdk.Context,
	cid string,
) (val types.LegacyActiveDeals, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LegacyActiveDealsKeyPrefix))

	b := store.Get(types.LegacyActiveDealsKey(
		cid,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveActiveDeals removes a activeDeals from the store
func (k Keeper) RemoveLegacyActiveDeals(
	ctx sdk.Context,
	cid string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LegacyActiveDealsKeyPrefix))
	store.Delete(types.LegacyActiveDealsKey(
		cid,
	))
}

// GetAllActiveDeals returns all activeDeals
func (k Keeper) GetAllLegacyActiveDeals(ctx sdk.Context) (list []types.LegacyActiveDeals) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LegacyActiveDealsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LegacyActiveDeals
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
