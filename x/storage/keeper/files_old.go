package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
)

// Deprecated: This is simply a utility function to facilitate a migration
func (k Keeper) setFilePrimaryOld(ctx sdk.Context, file types.UnifiedFile) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FilePrimaryKeyPrefix))
	b := k.cdc.MustMarshal(&file)
	store.Set(types.FilesPrimaryKey(
		file.Merkle,
		file.Owner,
		file.Start,
	), b)
}

// SetFileOld set a specific File in the store from its index
//
// Deprecated: This is simply a utility function to facilitate a migration
func (k Keeper) SetFileOld(ctx sdk.Context, file types.UnifiedFile) {
	k.setFilePrimaryOld(ctx, file)
}

// GetFileOld returns a File from its index
//
// Deprecated: This is simply a utility function to facilitate a migration
func (k Keeper) GetFileOld(
	ctx sdk.Context,
	merkle []byte,
	owner string,
	start int64,
) (val types.UnifiedFile, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FilePrimaryKeyPrefix))

	b := store.Get(types.FilesPrimaryKey(
		merkle, owner, start,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// Deprecated: This is simply a utility function to facilitate a migration
func (k Keeper) removeFilePrimaryOld(
	ctx sdk.Context,
	merkle []byte,
	owner string,
	start int64,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FilePrimaryKeyPrefix))
	store.Delete(types.FilesPrimaryKey(
		merkle,
		owner,
		start,
	))
}

// RemoveFileOld removes a File from the store
//
// Deprecated: This is simply a utility function to facilitate a migration
func (k Keeper) RemoveFileOld(
	ctx sdk.Context,
	merkle []byte,
	owner string,
	start int64,
) {
	file, found := k.GetFileOld(ctx, merkle, owner, start)
	if !found {
		return
	}

	for _, proof := range file.Proofs { // deleting all the associated proofs too
		k.RemoveProofWithBuiltKey(ctx, []byte(proof))
	}

	k.removeFilePrimaryOld(ctx, merkle, owner, start)
}

// GetAllFileByMerkleOld returns all File
//
// Deprecated: This is simply a utility function to facilitate a migration
func (k Keeper) GetAllFileByMerkleOld(ctx sdk.Context) (list []types.UnifiedFile) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FilePrimaryKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.UnifiedFile
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// IterateFilesByMerkleOld iterates through every file
//
// Deprecated: This is simply a utility function to facilitate a migration
func (k Keeper) IterateFilesByMerkleOld(ctx sdk.Context, reverse bool, fn func(key []byte, val []byte) bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FilePrimaryKeyPrefix))

	var iterator storetypes.Iterator
	if reverse {
		iterator = sdk.KVStoreReversePrefixIterator(store, []byte{})
	} else {
		iterator = sdk.KVStorePrefixIterator(store, []byte{})
	}

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		b := fn(iterator.Key(), iterator.Value())
		if b {
			return
		}
	}
}

// IterateAndParseFilesByMerkleOld iterates through every file and parses them for you
//
// Deprecated: This is simply a utility function to facilitate a migration
func (k Keeper) IterateAndParseFilesByMerkleOld(ctx sdk.Context, reverse bool, fn func(key []byte, val types.UnifiedFile) bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FilePrimaryKeyPrefix))

	var iterator storetypes.Iterator
	if reverse {
		iterator = sdk.KVStoreReversePrefixIterator(store, []byte{})
	} else {
		iterator = sdk.KVStorePrefixIterator(store, []byte{})
	}

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		val := iterator.Value()
		var file types.UnifiedFile
		if err := k.cdc.Unmarshal(val, &file); err != nil {
			return
		}

		b := fn(iterator.Key(), file)
		if b {
			return
		}
	}
}

// GetAllFilesWithMerkleOld returns all Files that start with a specific merkle
//
// Deprecated: This is simply a utility function to facilitate a migration
func (k Keeper) GetAllFilesWithMerkleOld(ctx sdk.Context, merkle []byte) (list []types.UnifiedFile) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.FilesMerklePrefix(merkle))
	iterator := sdk.KVStorePrefixIterator(store, nil)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.UnifiedFile
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
