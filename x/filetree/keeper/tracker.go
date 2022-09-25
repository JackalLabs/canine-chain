package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/filetree/types"
)

// SetTracker set tracker in the store
func (k Keeper) SetTracker(ctx sdk.Context, tracker types.Tracker) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TrackerKey))
	b := k.cdc.MustMarshal(&tracker)
	store.Set([]byte{0}, b)
}

// GetTracker returns tracker
func (k Keeper) GetTracker(ctx sdk.Context) (val types.Tracker, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TrackerKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTracker removes tracker from the store
func (k Keeper) RemoveTracker(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TrackerKey))
	store.Delete([]byte{0})
}
