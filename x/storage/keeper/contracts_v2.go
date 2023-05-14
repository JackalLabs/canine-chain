package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

// SetContracts set a specific contracts in the store from its index
func (k Keeper) SetContracts(ctx sdk.Context, contracts types.ContractV2) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractsV2KeyPrefix))
	b := k.cdc.MustMarshal(&contracts)
	store.Set(types.ContractsKey(
		contracts.Cid,
	), b)
}

// GetContracts returns a contracts from its index
func (k Keeper) GetContracts(
	ctx sdk.Context,
	cid string,
) (val types.ContractV2, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractsV2KeyPrefix))

	b := store.Get(types.ContractsKey(
		cid,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveContracts removes a contracts from the store
func (k Keeper) RemoveContracts(
	ctx sdk.Context,
	cid string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractsV2KeyPrefix))
	store.Delete(types.ContractsKey(
		cid,
	))
}

// GetAllContracts returns all contracts
func (k Keeper) GetAllContracts(ctx sdk.Context) (list []types.ContractV2) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractsV2KeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ContractV2
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
