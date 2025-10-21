package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v5/x/rns/types"
)

// SetBids set a specific bids in the store from its index
func (k Keeper) SetBids(ctx sdk.Context, bids types.Bids) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BidsKeyPrefix))
	b := k.cdc.MustMarshal(&bids)
	store.Set(types.BidsKey(
		bids.Index,
	), b)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventSetBid,
			sdk.NewAttribute(types.AttributeName, bids.Name),
			sdk.NewAttribute(types.AttributeBidder, bids.Bidder),
		),
	)
}

// GetBids returns a bids from its index
func (k Keeper) GetBids(
	ctx sdk.Context,
	index string,
) (val types.Bids, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BidsKeyPrefix))

	b := store.Get(types.BidsKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveBids removes a bids from the store
func (k Keeper) RemoveBids(
	ctx sdk.Context,
	index string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BidsKeyPrefix))
	store.Delete(types.BidsKey(
		index,
	))
}

// GetAllBids returns all bids
func (k Keeper) GetAllBids(ctx sdk.Context) (list []types.Bids) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BidsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Bids
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return list
}
