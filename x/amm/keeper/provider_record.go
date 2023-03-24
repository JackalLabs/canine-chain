package keeper

import (
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/x/amm/types"
)

// Set unlock time using block time and lock duration.
// Saves updated record to KVStore.
// Returns error when the record is not found.
func (k Keeper) EngageLock(
	ctx sdk.Context,
	recordKey []byte,
) error {

	record, found := k.GetProviderRecord(ctx, recordKey)

	if !found {
		return sdkerrors.Wrapf(
			types.ErrProviderRecordNotFound,
			"Cannot engage lock on record %s",
			recordKey,
		)
	}

	lockDuration, _ := time.ParseDuration(record.LockDuration)

	timeNow := ctx.BlockTime()

	record.UnlockTime = TimeToString(timeNow.Add(lockDuration))

	k.SetProviderRecord(ctx, record)

	return nil
}

// Create then store ProviderRecord and reference to KVStore.
// Lock is not engaged.
// It returns error when pool doesn't exist.
func (k Keeper) InitProviderRecord(
	ctx sdk.Context,
	provider sdk.AccAddress,
	poolName string,
	lockDuration time.Duration,
) error {

	// Find pool
	_, found := k.GetPool(ctx, poolName)

	if !found {
		return sdkerrors.Wrapf(
			types.ErrLiquidityPoolNotFound,
			"Cannot initialize ProviderRecord, pool(%s) not found",
			poolName,
		)
	}

	// Create record
	record := types.ProviderRecord{
		Provider:     provider.String(),
		PoolName:     poolName,
		LockDuration: lockDuration.String(),
	}

	k.SetProviderRecord(ctx, record)

	// Add reference
	if err := k.AddProviderRef(ctx, record); err != nil {
		return sdkerrors.Wrapf(
			err,
			"Cannot initialize ProviderRecord",
		)
	}

	return nil
}

// SetProviderRecord set a specific lProviderRecord in the store from its index
func (k Keeper) SetProviderRecord(ctx sdk.Context, lProviderRecord types.ProviderRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProviderRecordKeyPrefix))
	b := k.cdc.MustMarshal(&lProviderRecord)
	store.Set(types.ProviderRecordKey(
		lProviderRecord.PoolName,
		lProviderRecord.Provider,
	), b)
}

// GetProviderRecord returns a lProviderRecord from its index
func (k Keeper) GetProviderRecord(
	ctx sdk.Context,
	recordKey []byte,
) (val types.ProviderRecord, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProviderRecordKeyPrefix))

	b := store.Get(recordKey)

	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// Removes ProviderRecord and reference from store.
func (k Keeper) EraseProviderRecord(
	ctx sdk.Context,
	provider sdk.AccAddress,
	poolName string,
) error {
	recordKey := types.ProviderRecordKey(poolName, provider.String())

	record, found := k.GetProviderRecord(ctx, recordKey)

	if !found {
		return types.ErrProviderRecordNotFound
	}

	k.RemoveProviderRef(ctx, record)
	k.RemoveProviderRecord(ctx, recordKey)

	return nil
}

// RemoveProviderRecord removes a lProviderRecord from the store
func (k Keeper) RemoveProviderRecord(
	ctx sdk.Context,
	recordKey []byte,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProviderRecordKeyPrefix))
	store.Delete(recordKey)
}

// GetAllProviderRecord returns all lProviderRecord
func (k Keeper) GetAllProviderRecord(ctx sdk.Context) (list []types.ProviderRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProviderRecordKeyPrefix))
	iterator := store.Iterator(nil, nil)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ProviderRecord
		if err := k.cdc.Unmarshal(iterator.Value(), &val); err != nil {
			ctx.Logger().Error("\nFailed to unmarshal at GetAllProviderRecord()\n", err)
		}
		list = append(list, val)
	}

	return
}

// Collect all ProviderRecord of provider.
// Parse through all keys in KVStore that has {provider} as its prefix.
func (k Keeper) GetAllRecordOfProvider(
	ctx sdk.Context,
	provider sdk.AccAddress,
) (list []types.ProviderRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey),
		types.KeyPrefix(types.ProviderRecordKeyPrefix))

	iterator := sdk.KVStorePrefixIterator(store, []byte(provider.String()))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var record types.ProviderRecord

		rawRecord := store.Get(iterator.Value())

		if err := k.cdc.Unmarshal(rawRecord, &record); err != nil {
			ctx.Logger().Error("\nFailed to unmarshal at GetAllProviderRecord()\n")
		}

		list = append(list, record)
	}

	return
}

func (k Keeper) GetAllRecordOfPool(ctx sdk.Context, poolName string,
) (list []types.ProviderRecord) {

	store := prefix.NewStore(ctx.KVStore(k.storeKey),
		types.KeyPrefix(types.ProviderRecordKeyPrefix))


	iterator := sdk.KVStorePrefixIterator(store, []byte("ujkl-ujwl"))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ProviderRecord

		if err := k.cdc.Unmarshal(iterator.Value(), &val); err != nil {
			ctx.Logger().Error("\nFailed to unmarshal at GetAllProviderRecord()\n", err)
		}
		list = append(list, val)
	}

	return
}

// Add ProviderRecord reference to KVStore.
// Reference to ProviderRecord is ProviderRecord's key.
// Key to the reference key is {provider}{poolName}
// It returns error when reference key already exists.
func (k Keeper) AddProviderRef(ctx sdk.Context, record types.ProviderRecord) error {
	// Generate keys
	refKey := types.GetProviderRefKey(record)
	recordKey := types.GetProviderKey(record)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RefKeyPrefix))

	if store.Has(refKey) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"Reference key already exists: %s", refKey)
	}

	store.Set(refKey, recordKey)

	return nil
}

// Remove ProviderRecord reference from KVStore.
func (k Keeper) RemoveProviderRef(ctx sdk.Context, record types.ProviderRecord) {
	// Generate reference key
	refKey := types.GetProviderRefKey(record)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RefKeyPrefix))

	store.Delete(refKey)
}
