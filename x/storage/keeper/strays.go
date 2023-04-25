package keeper

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

// SetStrays set a specific strays in the store from its index
func (k Keeper) SetStrays(ctx sdk.Context, strays types.Strays) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StraysKeyPrefix))
	b := k.cdc.MustMarshal(&strays)
	store.Set(types.StraysKey(
		strays.Cid,
	), b)
}

// GetStrays returns a strays from its index
func (k Keeper) GetStrays(
	ctx sdk.Context,
	cid string,
) (val types.Strays, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StraysKeyPrefix))

	b := store.Get(types.StraysKey(
		cid,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStrays removes a strays from the store
func (k Keeper) RemoveStrays(
	ctx sdk.Context,
	cid string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StraysKeyPrefix))
	store.Delete(types.StraysKey(
		cid,
	))
}

// GetAllStrays returns all strays
func (k Keeper) GetAllStrays(ctx sdk.Context) (list []types.Strays) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StraysKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Strays
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) ClearDeadFiles(ctx sdk.Context) {
	strays := k.GetAllStrays(ctx)
	deals := k.GetAllActiveDeals(ctx)

	for _, stray := range strays {
		found := false
		for _, deal := range deals {
			if stray.Fid == deal.Fid {
				found = true
				break
			}
		}
		if found {
			continue
		}

		paymentInfo, found := k.GetStoragePaymentInfo(ctx, stray.Signee)
		if found {
			fSize, err := strconv.ParseInt(stray.Filesize, 10, 64)
			if err == nil {
				paymentInfo.SpaceUsed -= fSize // remove the deal from the users paid amount.
			}
		}
		k.RemoveStrays(ctx, stray.Cid)
	}
}
