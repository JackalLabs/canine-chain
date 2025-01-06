package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
)

func (k Keeper) FixStoragePaymentInfo(ctx sdk.Context, info types.StoragePaymentInfo) types.StoragePaymentInfo {
	address := info.Address

	filesByOwner := k.GetAllFilesWithOwner(ctx, address)

	var size int64
	for _, file := range filesByOwner {
		size += file.FileSize
	}

	info.SpaceUsed = size

	return info
}

// SetStoragePaymentInfo set a specific payBlocks in the store from its x
func (k Keeper) SetStoragePaymentInfo(ctx sdk.Context, payInfo types.StoragePaymentInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoragePaymentInfoKeyPrefix))
	b := k.cdc.MustMarshal(&payInfo)
	store.Set(types.StoragePaymentInfoKey(
		payInfo.Address,
	), b)
}

// GetStoragePaymentInfo returns StoragePaymentInfo from its address
func (k Keeper) GetStoragePaymentInfo(
	ctx sdk.Context,
	address string,
) (val types.StoragePaymentInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoragePaymentInfoKeyPrefix))

	b := store.Get(types.StoragePaymentInfoKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return k.FixStoragePaymentInfo(ctx, val), true
}

// RemoveStoragePaymentInfo removes a StoragePaymentInfo  from the store
func (k Keeper) RemoveStoragePaymentInfo(
	ctx sdk.Context,
	address string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoragePaymentInfoKeyPrefix))
	store.Delete(types.StoragePaymentInfoKey(
		address,
	))
}

// GetAllStoragePaymentInfo returns all payBlocks
func (k Keeper) GetAllStoragePaymentInfo(ctx sdk.Context) (list []types.StoragePaymentInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoragePaymentInfoKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.StoragePaymentInfo
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, k.FixStoragePaymentInfo(ctx, val))
	}

	return
}
