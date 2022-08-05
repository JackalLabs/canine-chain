package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/dsig/types"
)

// SetUserUploads set a specific userUploads in the store from its index
func (k Keeper) SetUserUploads(ctx sdk.Context, userUploads types.UserUploads) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserUploadsKeyPrefix))
	b := k.cdc.MustMarshal(&userUploads)
	store.Set(types.UserUploadsKey(
		userUploads.Fid,
	), b)
}

// GetUserUploads returns a userUploads from its index
func (k Keeper) GetUserUploads(
	ctx sdk.Context,
	fid string,

) (val types.UserUploads, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserUploadsKeyPrefix))

	b := store.Get(types.UserUploadsKey(
		fid,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveUserUploads removes a userUploads from the store
func (k Keeper) RemoveUserUploads(
	ctx sdk.Context,
	fid string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserUploadsKeyPrefix))
	store.Delete(types.UserUploadsKey(
		fid,
	))
}

// GetAllUserUploads returns all userUploads
func (k Keeper) GetAllUserUploads(ctx sdk.Context) (list []types.UserUploads) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserUploadsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.UserUploads
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
