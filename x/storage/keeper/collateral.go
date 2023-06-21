package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

// SetCollateral sets a specific collateral in the store from its index
func (k Keeper) SetCollateral(ctx sdk.Context, collateral types.Collateral) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CollateralKeyPrefix))
	b := k.cdc.MustMarshal(&collateral)
	store.Set(types.CollateralKey(
		collateral.Address,
	), b)
}

// GetCollateral returns a collateral from its index
func (k Keeper) GetCollateral(
	ctx sdk.Context,
	address string,
) (val types.Collateral, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CollateralKeyPrefix))

	b := store.Get(types.CollateralKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCollateral removes a collateral from the store
func (k Keeper) RemoveCollateral(
	ctx sdk.Context,
	address string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CollateralKeyPrefix))
	store.Delete(types.CollateralKey(
		address,
	))
}

// GetAllCollateral returns all collaterals
func (k Keeper) GetAllCollateral(ctx sdk.Context) (list []types.Collateral) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CollateralKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Collateral
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
