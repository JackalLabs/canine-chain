package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/notifications/types"
)

// SetNotifications set a specific notifications in the store from its index
func (k Keeper) SetNotifications(ctx sdk.Context, notifications types.Notifications, address string) {
	keyPrefix := fmt.Sprintf("%s%s/", types.NotificationsKeyPrefix, address)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(keyPrefix))
	b := k.cdc.MustMarshal(&notifications)
	store.Set(types.NotificationsKey(
		notifications.Count,
	), b)
}

// GetNotifications returns a notifications from its index
func (k Keeper) GetNotifications(
	ctx sdk.Context,
	count uint64,
	address string,
) (val types.Notifications, found bool) {
	keyPrefix := fmt.Sprintf("%s%s/", types.NotificationsKeyPrefix, address)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(keyPrefix))

	b := store.Get(types.NotificationsKey(
		count,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNotifications removes a notifications from the store
func (k Keeper) RemoveNotifications(
	ctx sdk.Context,
	count uint64,
	address string,
) {
	keyPrefix := fmt.Sprintf("%s%s/", types.NotificationsKeyPrefix, address)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(keyPrefix))
	store.Delete(types.NotificationsKey(
		count,
	))
}

// GetAllNotificationsForUser returns all notifications for a user
func (k Keeper) GetAllNotificationsForUser(ctx sdk.Context, address string) (list []types.Notifications) {
	keyPrefix := fmt.Sprintf("%s%s/", types.NotificationsKeyPrefix, address)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(keyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{}) // replace []byte{} with keyPrefix?

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Notifications
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllNotifications returns all notifications
func (k Keeper) GetAllNotifications(ctx sdk.Context) (list []types.Notifications) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NotificationsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{}) // replace []byte{} with keyPrefix?

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Notifications
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
