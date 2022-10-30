package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/filetree/types"
)

// SetPubkey set a specific pubkey in the store from its index
func (k Keeper) SetPubkey(ctx sdk.Context, pubkey types.Pubkey) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PubkeyKeyPrefix))
	b := k.cdc.MustMarshal(&pubkey)
	store.Set(types.PubkeyKey(
		pubkey.Address,
	), b)
}

// GetPubkey returns a pubkey from its index
func (k Keeper) GetPubkey(
	ctx sdk.Context,
	address string,
) (val types.Pubkey, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PubkeyKeyPrefix))

	b := store.Get(types.PubkeyKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePubkey removes a pubkey from the store
func (k Keeper) RemovePubkey(
	ctx sdk.Context,
	address string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PubkeyKeyPrefix))
	store.Delete(types.PubkeyKey(
		address,
	))
}

// GetAllPubkey returns all pubkey
func (k Keeper) GetAllPubkey(ctx sdk.Context) (list []types.Pubkey) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PubkeyKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Pubkey
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
