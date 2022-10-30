package keeper

import (
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/x/lp/types"
)

// Set unlock time using block time and lock duration.
// Saves updated record to KVStore.
// Returns error when the record is not found.
func (k Keeper) EngageLock(
	ctx sdk.Context,
	recordKey []byte,
) error {
	record, found := k.GetLProviderRecord(ctx, recordKey)

	if !found {
		return sdkerrors.Wrapf(
			types.ErrLProviderRecordNotFound,
			"Cannot engage lock on record %s",
			recordKey,
		)
	}

	lockDuration, _ := time.ParseDuration(record.LockDuration)

	timeNow := ctx.BlockTime()

	record.UnlockTime = TimeToString(timeNow.Add(lockDuration))

	k.SetLProviderRecord(ctx, record)

	return nil
}

// Create then store LProviderRecord and reference to KVStore.
// Lock is not engaged.
// It returns error when pool doesn't exist.
func (k Keeper) InitLProviderRecord(
	ctx sdk.Context,
	provider sdk.AccAddress,
	poolName string,
	lockDuration time.Duration,
) error {
	// Find pool
	_, found := k.GetLPool(ctx, poolName)

	if !found {
		return sdkerrors.Wrapf(
			types.ErrLiquidityPoolNotFound,
			"Cannot initialize LProviderRecord, pool(%s) not found",
			poolName,
		)
	}

	// Create record
	record := types.LProviderRecord{
		Provider:     provider.String(),
		PoolName:     poolName,
		LockDuration: lockDuration.String(),
	}

	k.SetLProviderRecord(ctx, record)

	// Add reference
	if err := k.AddProviderRef(ctx, record); err != nil {
		return sdkerrors.Wrapf(
			err,
			"Cannot initialize LProviderRecord",
		)
	}

	return nil
}

// SetLProviderRecord set a specific lProviderRecord in the store from its index
func (k Keeper) SetLProviderRecord(ctx sdk.Context, lProviderRecord types.LProviderRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LProviderRecordKeyPrefix))
	b := k.cdc.MustMarshal(&lProviderRecord)
	store.Set(types.LProviderRecordKey(
		lProviderRecord.PoolName,
		lProviderRecord.Provider,
	), b)
}

// GetLProviderRecord returns a lProviderRecord from its index
func (k Keeper) GetLProviderRecord(
	ctx sdk.Context,
	recordKey []byte,
) (val types.LProviderRecord, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LProviderRecordKeyPrefix))

	b := store.Get(recordKey)

	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// Removes LProviderRecord and reference from store.
func (k Keeper) EraseLProviderRecord(
	ctx sdk.Context,
	provider sdk.AccAddress,
	poolName string,
) error {
	recordKey := types.LProviderRecordKey(poolName, provider.String())

	record, found := k.GetLProviderRecord(ctx, recordKey)

	if !found {
		return types.ErrLProviderRecordNotFound
	}

	k.RemoveProviderRef(ctx, record)
	k.RemoveLProviderRecord(ctx, recordKey)

	return nil
}

// RemoveLProviderRecord removes a lProviderRecord from the store
func (k Keeper) RemoveLProviderRecord(
	ctx sdk.Context,
	recordKey []byte,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LProviderRecordKeyPrefix))
	store.Delete(recordKey)
}

// GetAllLProviderRecord returns all lProviderRecord
func (k Keeper) GetAllLProviderRecord(ctx sdk.Context) (list []types.LProviderRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LProviderRecordKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix("stake-token"))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LProviderRecord
		if err := k.cdc.Unmarshal(iterator.Value(), &val); err != nil {
			ctx.Logger().Error("\nFailed to unmarshal at GetAllLProviderRecord()\n", err)
		}
		list = append(list, val)
	}

	return
}

// Collect all LProviderRecord of provider.
// Parse through all keys in KVStore that has {provider} as its prefix.
func (k Keeper) GetAllRecordOfProvider(
	ctx sdk.Context,
	provider sdk.AccAddress,
) (list []types.LProviderRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey),
		types.KeyPrefix(types.LProviderRecordKeyPrefix))

	iterator := sdk.KVStorePrefixIterator(store, []byte(provider.String()))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var record types.LProviderRecord

		rawRecord := store.Get(iterator.Value())

		k.cdc.MustUnmarshal(rawRecord, &record)

		list = append(list, record)
	}

	return
}

func (k Keeper) GetAllRecordOfPool(ctx sdk.Context, poolName string,
) (list []types.LProviderRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey),
		types.KeyPrefix(types.LProviderRecordKeyPrefix))

	strKey := poolName

	iterator := sdk.KVStorePrefixIterator(store, []byte(strKey))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var record types.LProviderRecord

		rawRecord := store.Get(iterator.Value())

		k.cdc.MustUnmarshal(rawRecord, &record)

		list = append(list, record)
	}

	return
}

// Add LProviderRecord reference to KVStore.
// Reference to LProviderRecord is LProviderRecord's key.
// Key to the reference key is {provider}{poolName}
// It returns error when reference key already exists.
func (k Keeper) AddProviderRef(ctx sdk.Context, record types.LProviderRecord) error {
	// Generate keys
	refKey := types.GetProviderRefKey(record)
	recordKey := types.GetProviderKey(record)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LProviderRecordKeyPrefix))

	if store.Has(refKey) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"Reference key already exists: %s", refKey)
	}

	store.Set(refKey, recordKey)

	return nil
}

// Remove LProviderRecord reference from KVStore.
func (k Keeper) RemoveProviderRef(ctx sdk.Context, record types.LProviderRecord) {
	// Generate reference key
	refKey := types.GetProviderRefKey(record)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LProviderRecordKeyPrefix))

	store.Delete(refKey)
}
