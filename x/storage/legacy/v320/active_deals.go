package v320

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

// RemoveLegacyActiveDeals removes a legacy activeDeals from the store
func RemoveLegacyActiveDeals(
	ctx sdk.Context,
	cid string,
	storeKey sdk.StoreKey,
) {
	store := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.ActiveDealsKeyPrefix))
	store.Delete(types.ActiveDealsKey(
		cid,
	))
}

// IterateLegacyActiveDeals runs `fn` for each legacy active deal in the store
func IterateLegacyActiveDeals(ctx sdk.Context, storeKey sdk.StoreKey, cdc codec.BinaryCodec, fn func(deal types.LegacyActiveDeals) bool) {
	store := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.ActiveDealsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LegacyActiveDeals
		cdc.MustUnmarshal(iterator.Value(), &val)

		if fn(val) {
			return
		}

	}
}
