package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

// SetProof set a specific proof in the store from its index
func (k Keeper) SetProof(ctx sdk.Context, proof types.FileProof) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProofKeyPrefix))
	b := k.cdc.MustMarshal(&proof)
	store.Set(types.ProofKey(
		proof.Prover,
		proof.Merkle,
		proof.Owner,
		proof.Start,
	), b)
}

// GetProof returns a Proof from its index
func (k Keeper) GetProof(
	ctx sdk.Context,
	prover string,
	merkle []byte,
	owner string,
	start int64,
) (val types.FileProof, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProofKeyPrefix))

	b := store.Get(types.ProofKey(
		prover,
		merkle, owner, start,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetProofWithBuiltKey returns a Proof from its index using an already built key (for example from the internal file proof list)
func (k Keeper) GetProofWithBuiltKey(
	ctx sdk.Context,
	key []byte,
) (val types.FileProof, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProofKeyPrefix))

	b := store.Get(key)
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) RemoveProof(
	ctx sdk.Context,
	prover string,
	merkle []byte,
	owner string,
	start int64,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProofKeyPrefix))
	store.Delete(types.ProofKey(
		prover,
		merkle,
		owner,
		start,
	))
}

func (k Keeper) RemoveProofWithBuiltKey(
	ctx sdk.Context,
	key []byte,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProofKeyPrefix))
	store.Delete(key)
}

// GetAllProofs returns all File
func (k Keeper) GetAllProofs(ctx sdk.Context) (list []types.FileProof) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProofKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.FileProof
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllProofsForProver returns all Proofs for the given prover
func (k Keeper) GetAllProofsForProver(ctx sdk.Context, prover string) (list []types.FileProof) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProofKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte(prover))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.FileProof
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
