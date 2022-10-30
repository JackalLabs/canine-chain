package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

// SetProofs set a specific proofs in the store from its index
func (k Keeper) SetProofs(ctx sdk.Context, proofs types.Proofs) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProofsKeyPrefix))
	b := k.cdc.MustMarshal(&proofs)
	store.Set(types.ProofsKey(
		proofs.Cid,
	), b)
}

// GetProofs returns a proofs from its index
func (k Keeper) GetProofs(
	ctx sdk.Context,
	cid string,
) (val types.Proofs, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProofsKeyPrefix))

	b := store.Get(types.ProofsKey(
		cid,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveProofs removes a proofs from the store
func (k Keeper) RemoveProofs(
	ctx sdk.Context,
	cid string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProofsKeyPrefix))
	store.Delete(types.ProofsKey(
		cid,
	))
}

// GetAllProofs returns all proofs
func (k Keeper) GetAllProofs(ctx sdk.Context) (list []types.Proofs) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProofsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Proofs
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
