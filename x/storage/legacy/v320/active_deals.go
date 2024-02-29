package v320

import (
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
