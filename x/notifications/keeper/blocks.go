package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v5/x/notifications/types"
)

// SetBlock sets a specific block value in the store
func (k Keeper) SetBlock(ctx sdk.Context, block types.Block) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NotificationsKeyPrefix))
	b := k.cdc.MustMarshal(&block)
	store.Set(types.BlockKey(
		block.Address,
		block.BlockedAddress,
	), b)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventBlock,
			sdk.NewAttribute(types.AttributeBlocker, block.Address),
			sdk.NewAttribute(types.AttributeBlockee, block.BlockedAddress),
		),
	)
}

// IsBlocked returns if a user is blocked
func (k Keeper) IsBlocked(
	ctx sdk.Context,
	owner string,
	from string,
) (found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NotificationsKeyPrefix))

	key := types.BlockKey(
		owner,
		from,
	)

	return store.Has(key)
}

// RemoveBlock removes a block from the store
func (k Keeper) RemoveBlock(
	ctx sdk.Context,
	owner string,
	from string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NotificationsKeyPrefix))
	store.Delete(types.BlockKey(
		owner,
		from,
	))
}
