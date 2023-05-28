package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

// SetAttestationForm sets a specific attestation in the store from its index
func (k Keeper) SetAttestationForm(ctx sdk.Context, attestation types.AttestationForm) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AttestationKeyPrefix))
	b := k.cdc.MustMarshal(&attestation)
	store.Set(types.AttestationKey(
		attestation.Cid,
	), b)
}

// GetAttestationForm returns an attestation from its index
func (k Keeper) GetAttestationForm(
	ctx sdk.Context,
	cid string,
) (val types.AttestationForm, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AttestationKeyPrefix))

	b := store.Get(types.AttestationKey(
		cid,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAttestation removes an attestation from the store
func (k Keeper) RemoveAttestation(
	ctx sdk.Context,
	cid string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AttestationKeyPrefix))
	store.Delete(types.AttestationKey(
		cid,
	))
}

// GetAllAttestation returns all attestations
func (k Keeper) GetAllAttestation(ctx sdk.Context) (list []types.AttestationForm) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AttestationKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AttestationForm
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// RemoveAllAttestation removes all attestations
func (k Keeper) RemoveAllAttestation(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AttestationKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AttestationForm
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		store.Delete(types.AttestationKey(
			val.Cid,
		))
	}
}
