package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/jklmining/types"
)

// SetMinerClaims set a specific minerClaims in the store from its index
func (k Keeper) SetMinerClaims(ctx sdk.Context, minerClaims types.MinerClaims) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MinerClaimsKeyPrefix))
	b := k.cdc.MustMarshal(&minerClaims)
	store.Set(types.MinerClaimsKey(
		minerClaims.Hash,
	), b)
}

// GetMinerClaims returns a minerClaims from its index
func (k Keeper) GetMinerClaims(
	ctx sdk.Context,
	hash string,

) (val types.MinerClaims, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MinerClaimsKeyPrefix))

	b := store.Get(types.MinerClaimsKey(
		hash,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMinerClaims removes a minerClaims from the store
func (k Keeper) RemoveMinerClaims(
	ctx sdk.Context,
	hash string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MinerClaimsKeyPrefix))
	store.Delete(types.MinerClaimsKey(
		hash,
	))
}

// GetAllMinerClaims returns all minerClaims
func (k Keeper) GetAllMinerClaims(ctx sdk.Context) (list []types.MinerClaims) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MinerClaimsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.MinerClaims
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
