package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v5/x/jklmint/types"
)

func (k Keeper) SetMintedBlock(ctx sdk.Context, block types.MintedBlock) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LastBlockMinted))
	b := k.cdc.MustMarshal(&block)
	store.Set(types.MintedBlockKey(
		block.Height,
	), b)
}

func (k Keeper) GetMintedBlock(
	ctx sdk.Context,
	height int64,
) (val types.MintedBlock, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LastBlockMinted))

	b := store.Get(types.MintedBlockKey(
		height,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}
