package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/filetree/types"
)

// SetFiles set a specific files in the store from its index
func (k Keeper) SetFiles(ctx sdk.Context, files types.Files) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FilesKeyPrefix))

	b := k.cdc.MustMarshal(&files)

	store.Set(types.FilesKey(
		files.Address,
		files.Owner,
	), b)

}

// GetFiles returns a files from its index
func (k Keeper) GetFiles(
	ctx sdk.Context,
	address string,
	ownerAddress string,

) (val types.Files, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FilesKeyPrefix))

	b := store.Get(types.FilesKey(
		address,
		ownerAddress,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveFiles removes a files from the store
func (k Keeper) RemoveFiles(
	ctx sdk.Context,
	address string,
	ownerAddress string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FilesKeyPrefix))
	store.Delete(types.FilesKey(
		address,
		ownerAddress,
	))
}

// GetAllFiles returns all files
func (k Keeper) GetAllFiles(ctx sdk.Context) (list []types.Files) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FilesKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Files
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
