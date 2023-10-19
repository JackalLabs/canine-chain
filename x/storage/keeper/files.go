package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

func (k Keeper) setFilePrimary(ctx sdk.Context, File types.UnifiedFile) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FilePrimaryKeyPrefix))
	b := k.cdc.MustMarshal(&File)
	store.Set(types.FilesPrimaryKey(
		File.Merkle,
		File.Owner,
		File.Start,
	), b)
}

func (k Keeper) setFileSecondary(ctx sdk.Context, File types.UnifiedFile) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FileSecondaryKeyPrefix))
	b := k.cdc.MustMarshal(&File)
	store.Set(types.FilesSecondaryKey(
		File.Merkle,
		File.Owner,
		File.Start,
	), b)
}

// SetFile set a specific File in the store from its index
func (k Keeper) SetFile(ctx sdk.Context, File types.UnifiedFile) {
	k.setFilePrimary(ctx, File)
	k.setFileSecondary(ctx, File)
}

// GetFile returns a File from its index
func (k Keeper) GetFile(
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

func (k Keeper) removeFilePrimary(
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

func (k Keeper) removeFileSecondary(
	ctx sdk.Context,
	merkle []byte,
	owner string,
	start int64,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FileSecondaryKeyPrefix))
	store.Delete(types.FilesSecondaryKey(
		merkle,
		owner,
		start,
	))
}

// RemoveFile removes a File from the store
func (k Keeper) RemoveFile(
	ctx sdk.Context,
	merkle []byte,
	owner string,
	start int64,
) {
	file, found := k.GetFile(ctx, merkle, owner, start)
	if !found {
		return
	}

	for _, proof := range file.Proofs { // deleting all the associated proofs too
		k.RemoveProofWithBuiltKey(ctx, []byte(proof))
	}

	k.removeFilePrimary(ctx, merkle, owner, start)
	k.removeFileSecondary(ctx, merkle, owner, start)
}

// GetAllFileByMerkle returns all File
func (k Keeper) GetAllFileByMerkle(ctx sdk.Context) (list []types.UnifiedFile) {
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

// GetAllFileByOwner returns all File
func (k Keeper) GetAllFileByOwner(ctx sdk.Context) (list []types.UnifiedFile) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FileSecondaryKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.UnifiedFile
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
