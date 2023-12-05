//nolint:all
package keeper

import (
	"bytes"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

// SetOracleRequest set a specific oracle request in the store from its index
func (k Keeper) SetOracleRequest(ctx sdk.Context, req types.OracleRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OracleRequestKey))
	b := k.cdc.MustMarshal(&req)
	store.Set(types.RequestKey(
		req.Requester,
		req.Merkle,
		req.Chunk,
	), b)
}

// GetOracleRequest returns an oracle request from its index
func (k Keeper) GetOracleRequest(
	ctx sdk.Context,
	owner string,
	merkle []byte,
	chunk int64,
) (val types.OracleRequest, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OracleRequestKey))

	b := store.Get(types.RequestKey(
		owner,
		merkle,
		chunk,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOracleRequest removes a activeDeals from the store
func (k Keeper) RemoveOracleRequest(
	ctx sdk.Context,
	owner string,
	merkle []byte,
	chunk int64,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OracleRequestKey))
	store.Delete(types.RequestKey(
		owner,
		merkle,
		chunk,
	))
}

// GetAllOracleRequests returns all oracle requests
func (k Keeper) GetAllOracleRequests(ctx sdk.Context) (list []types.OracleRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OracleRequestKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.OracleRequest
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// SetOracleEntry set a specific oracle entry in the store from its index
func (k Keeper) SetOracleEntry(ctx sdk.Context, req types.OracleEntry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OracleEntryKey))
	b := k.cdc.MustMarshal(&req)
	store.Set(types.EntryKey(
		req.Owner,
		req.Merkle,
		req.Chunk,
	), b)
}

// GetOracleEntry returns an oracle entry from its index
func (k Keeper) GetOracleEntry(
	ctx sdk.Context,
	owner string,
	merkle []byte,
	chunk int64,
) (val types.OracleEntry, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OracleEntryKey))

	b := store.Get(types.EntryKey(
		owner,
		merkle,
		chunk,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOracleEntry removes a activeDeals from the store
func (k Keeper) RemoveOracleEntry(
	ctx sdk.Context,
	owner string,
	merkle []byte,
	chunk int64,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OracleEntryKey))
	store.Delete(types.EntryKey(
		owner,
		merkle,
		chunk,
	))
}

// GetAllOracleEntries returns all oracle entries
func (k Keeper) GetAllOracleEntries(ctx sdk.Context) (list []types.OracleEntry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OracleEntryKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.OracleEntry
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetGeneralOracleEntry returns the first oracle entry matching the merkle and chunk
func (k Keeper) GetGeneralOracleEntry(ctx sdk.Context, merkle []byte, chunk int64) (*types.OracleEntry, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OracleEntryKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val *types.OracleEntry
		k.cdc.MustUnmarshal(iterator.Value(), val)
		if bytes.Equal(val.Merkle, merkle) && val.Chunk == chunk {
			return val, nil
		}
	}

	return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "cannot find the oracle entry with merkle: %x and chunk %d", merkle, chunk)
}
