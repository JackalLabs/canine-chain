package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/jklmining/types"
)

// GetMinedCount get the total number of mined
func (k Keeper) GetMinedStarting(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MinedStartKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// GetMinedCount get the total number of mined
func (k Keeper) PushMinedStarting(ctx sdk.Context, amount uint64) {

	current := k.GetMinedStarting(ctx)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MinedStartKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, current+amount)
	store.Set(byteKey, bz)

}

// GetMinedCount get the total number of mined
func (k Keeper) GetMinedCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MinedCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetMinedCount set the total number of mined
func (k Keeper) SetMinedCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MinedCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendMined appends a mined in the store with a new id and update the count
func (k Keeper) AppendMined(
	ctx sdk.Context,
	mined types.Mined,
) uint64 {
	// Create the mined
	count := k.GetMinedCount(ctx)

	// Set the ID of the appended value
	mined.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MinedKey))
	appendedValue := k.cdc.MustMarshal(&mined)
	store.Set(GetMinedIDBytes(mined.Id), appendedValue)

	// Update mined count
	k.SetMinedCount(ctx, count+1)

	return count
}

// SetMined set a specific mined in the store
func (k Keeper) SetMined(ctx sdk.Context, mined types.Mined) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MinedKey))
	b := k.cdc.MustMarshal(&mined)
	store.Set(GetMinedIDBytes(mined.Id), b)
}

// GetMined returns a mined from its id
func (k Keeper) GetMined(ctx sdk.Context, id uint64) (val types.Mined, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MinedKey))
	b := store.Get(GetMinedIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMined removes a mined from the store
func (k Keeper) RemoveMined(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MinedKey))
	store.Delete(GetMinedIDBytes(id))
}

// GetAllMined returns all mined
func (k Keeper) GetAllMined(ctx sdk.Context) (list []types.Mined) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MinedKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Mined
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetMinedIDBytes returns the byte representation of the ID
func GetMinedIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetMinedIDFromBytes returns ID in uint64 format from a byte array
func GetMinedIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
