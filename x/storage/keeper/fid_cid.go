package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/storage/types"
)

// SetFidCid set a specific fidCid in the store from its index
func (k Keeper) SetFidCid(ctx sdk.Context, fidCid types.FidCid) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FidCidKeyPrefix))
	b := k.cdc.MustMarshal(&fidCid)
	store.Set(types.FidCidKey(
		fidCid.Fid,
	), b)
}

// GetFidCid returns a fidCid from its index
func (k Keeper) GetFidCid(
	ctx sdk.Context,
	fid string,
) (val types.FidCid, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FidCidKeyPrefix))

	b := store.Get(types.FidCidKey(
		fid,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveFidCid removes a fidCid from the store
func (k Keeper) RemoveFidCid(
	ctx sdk.Context,
	fid string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FidCidKeyPrefix))
	store.Delete(types.FidCidKey(
		fid,
	))
}

// GetAllFidCid returns all fidCid
func (k Keeper) GetAllFidCid(ctx sdk.Context) (list []types.FidCid) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FidCidKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.FidCid
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
