package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/amm/types"
)

func (k Keeper) SetPoolCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	key := types.KeyPrefix(types.PoolCountKey)
	
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(key, bz)
}

func (k Keeper) GetPoolCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	key := types.KeyPrefix(types.PoolCountKey)

	bz := store.Get(key)
	if bz == nil{
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) GetPoolCountAndIncrement(ctx sdk.Context) (count uint64) {
	count = k.GetPoolCount(ctx)
	k.SetPoolCount(ctx, count+1)

	return
}
