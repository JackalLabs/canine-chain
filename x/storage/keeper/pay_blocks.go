package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

// SetPayBlocks set a specific payBlocks in the store from its index
func (k Keeper) SetPayBlocks(ctx sdk.Context, payBlocks types.PayBlocks) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PayBlocksKeyPrefix))
	b := k.cdc.MustMarshal(&payBlocks)
	store.Set(types.PayBlocksKey(
		payBlocks.Blockid,
	), b)
}

// GetPayBlocks returns a payBlocks from its index
func (k Keeper) GetPayBlocks(
	ctx sdk.Context,
	blockid string,
) (val types.PayBlocks, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PayBlocksKeyPrefix))

	b := store.Get(types.PayBlocksKey(
		blockid,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePayBlocks removes a payBlocks from the store
func (k Keeper) RemovePayBlocks(
	ctx sdk.Context,
	blockid string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PayBlocksKeyPrefix))
	store.Delete(types.PayBlocksKey(
		blockid,
	))
}

// GetAllPayBlocks returns all payBlocks
func (k Keeper) GetAllPayBlocks(ctx sdk.Context) (list []types.PayBlocks) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PayBlocksKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PayBlocks
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
