package keeper

import (
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/notifications/types"
)

// SetNotification set a specific notifications in the store from its index
func (k Keeper) SetNotification(ctx sdk.Context, notification types.Notification) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NotificationsKeyPrefix))
	b := k.cdc.MustMarshal(&notification)
	store.Set(types.NotificationsKey(
		notification.To,
		notification.From,
		notification.Time,
	), b)
}

// GetNotification returns a notification from its index
func (k Keeper) GetNotification(
	ctx sdk.Context,
	to string,
	from string,
	timeStamp time.Time,
) (val types.Notification, found bool) {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NotificationsKeyPrefix))

	key := types.NotificationsKey(
		to,
		from,
		timeStamp,
	)

	if !store.Has(key) {
		return types.Notification{}, false
	}

	k.cdc.MustUnmarshal(store.Get(key), &val)
	return val, true
}

// RemoveNotification removes a notification from the store
func (k Keeper) RemoveNotification(
	ctx sdk.Context,
	to string,
	from string,
	timeStamp time.Time,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NotificationsKeyPrefix))
	store.Delete(types.NotificationsKey(
		to,
		from,
		timeStamp,
	))
}

// GetAllNotifications returns all notifications
func (k Keeper) GetAllNotifications(ctx sdk.Context) (list []types.Notification) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NotificationsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Notification
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
