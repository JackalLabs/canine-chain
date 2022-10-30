package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/notifications/types"
)

// SetNotifications set a specific notifications in the store from its index
func (k Keeper) SetNotifications(ctx sdk.Context, notifications types.Notifications) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NotificationsKeyPrefix))
	b := k.cdc.MustMarshal(&notifications)
	store.Set(types.NotificationsKey(
		notifications.Count,
		notifications.Address,
	), b)
}

// GetNotifications returns a notifications from its index
func (k Keeper) GetNotifications(
	ctx sdk.Context,
	count uint64,
	address string,
) (val types.Notifications, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NotificationsKeyPrefix))

	b := store.Get(types.NotificationsKey(
		count,
		address,
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
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NotificationsKeyPrefix))
	store.Delete(types.NotificationsKey(
		count,
		address,
	))
}

// GetAllNotifications returns all notifications
func (k Keeper) GetAllNotifications(ctx sdk.Context) (list []types.Notifications) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NotificationsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Notifications
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
