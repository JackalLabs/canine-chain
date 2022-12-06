package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

// SetStoragePaymentInfo set a specific payBlocks in the store from its index
func (k Keeper) SetStoragePaymentInfo(ctx sdk.Context, payInfo types.StoragePaymentInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoragePaymentInfoKeyPrefix))
	b := k.cdc.MustMarshal(&payInfo)
	store.Set(types.StoragePaymentInfoKey(
		payInfo.Address,
	), b)
}

// GetStoragePaymentInfo returns a payBlocks from its index
func (k Keeper) GetStoragePaymentInfo(
	ctx sdk.Context,
	blockid string,
) (val types.StoragePaymentInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoragePaymentInfoKeyPrefix))

	b := store.Get(types.StoragePaymentInfoKey(
		blockid,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStoragePaymentInfo removes a payBlocks from the store
func (k Keeper) RemoveStoragePaymentInfo(
	ctx sdk.Context,
	blockid string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoragePaymentInfoKeyPrefix))
	store.Delete(types.StoragePaymentInfoKey(
		blockid,
	))
}

// GetAllStoragePaymentInfo returns all payBlocks
func (k Keeper) GetAllStoragePaymentInfo(ctx sdk.Context) (list []*types.StoragePaymentInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoragePaymentInfoKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.StoragePaymentInfo
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, &val)
	}

	return
}
