package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/jklaccounts/types"
)

// SetAccounts set a specific accounts in the store from its index
func (k Keeper) SetAccounts(ctx sdk.Context, accounts types.Accounts) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountsKeyPrefix))
	b := k.cdc.MustMarshal(&accounts)
	store.Set(types.AccountsKey(
		accounts.Address,
	), b)
}

// GetAccounts returns a accounts from its index
func (k Keeper) GetAccounts(
	ctx sdk.Context,
	address string,

) (val types.Accounts, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountsKeyPrefix))

	b := store.Get(types.AccountsKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAccounts removes a accounts from the store
func (k Keeper) RemoveAccounts(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountsKeyPrefix))
	store.Delete(types.AccountsKey(
		address,
	))
}

// GetAllAccounts returns all accounts
func (k Keeper) GetAllAccounts(ctx sdk.Context) (list []types.Accounts) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Accounts
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
