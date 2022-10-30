package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/dsig/types"
)

// SetForm set a specific form in the store from its index
func (k Keeper) SetForm(ctx sdk.Context, form types.Form) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FormKeyPrefix))
	b := k.cdc.MustMarshal(&form)
	store.Set(types.FormKey(
		form.Ffid,
	), b)
}

// GetForm returns a form from its index
func (k Keeper) GetForm(
	ctx sdk.Context,
	ffid string,
) (val types.Form, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FormKeyPrefix))

	b := store.Get(types.FormKey(
		ffid,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveForm removes a form from the store
func (k Keeper) RemoveForm(
	ctx sdk.Context,
	ffid string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FormKeyPrefix))
	store.Delete(types.FormKey(
		ffid,
	))
}

// GetAllForm returns all form
func (k Keeper) GetAllForm(ctx sdk.Context) (list []types.Form) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FormKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Form
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
