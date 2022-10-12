package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/storage/types"
)

const (
	START_BLOCK_TYPE = "start"
	END_BLOCK_TYPE   = "end"
	TWO_GIGS         = 2000000000
)

func (k Keeper) GetPaidAmount(ctx sdk.Context, address string, blockh int64) int64 {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PayBlocksKeyPrefix))

	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	var highestBlock int64 = 0

	eblock, found := k.GetPayBlocks(ctx, fmt.Sprintf(".%s", address))
	if !found {
		return TWO_GIGS
	}

	endblock, ok := sdk.NewIntFromString(eblock.Blocknum)
	if !ok {
		return TWO_GIGS
	}

	if endblock.Int64() <= blockh {
		return TWO_GIGS
	}

	for ; iterator.Valid(); iterator.Next() {
		var val types.PayBlocks
		k.cdc.MustUnmarshal(iterator.Value(), &val)

		fmt.Printf("BLOCK %s: %s", val.Blocktype, val.Blocknum)

		if val.Blocktype == END_BLOCK_TYPE {
			continue
		}

		blocknum, ok := sdk.NewIntFromString(val.Blocknum)
		if !ok {
			continue
		}

		if blocknum.Int64() > blockh {
			continue
		}

		if blocknum.Int64() > highestBlock {
			highestBlock = blocknum.Int64()
			fmt.Printf("NEW HIGHEST BLOCK: %s", val.Blocknum)

		}

	}

	if highestBlock == 0 {
		return TWO_GIGS
	}

	hblock, found := k.GetPayBlocks(ctx, fmt.Sprintf("%s%d", address, highestBlock))
	if !found {
		return TWO_GIGS
	}

	bytes, ok := sdk.NewIntFromString(hblock.Bytes)
	if !ok {
		return TWO_GIGS
	}

	return bytes.Int64()
}

func (k Keeper) CreatePayBlock(ctx sdk.Context, address string, length int64, bytes int64) error {

	startBlock := ctx.BlockHeight()

	endBlock := startBlock + length

	sBlock := types.PayBlocks{
		Blockid:   fmt.Sprintf("%s%d", address, startBlock),
		Bytes:     fmt.Sprintf("%d", bytes),
		Blocktype: START_BLOCK_TYPE,
		Blocknum:  fmt.Sprintf("%d", startBlock),
	}

	eBlock := types.PayBlocks{
		Blockid:   fmt.Sprintf(".%s", address),
		Bytes:     fmt.Sprintf("%d", bytes),
		Blocktype: END_BLOCK_TYPE,
		Blocknum:  fmt.Sprintf("%d", endBlock),
	}

	paidamt := k.GetPaidAmount(ctx, address, endBlock)

	if paidamt > 0 {
		return fmt.Errorf("can't buy storage within another storage window")
	}

	k.SetPayBlocks(ctx, sBlock)
	k.SetPayBlocks(ctx, eBlock)

	return nil
}

func (k Keeper) GetProviderUsing(ctx sdk.Context, provider string) int64 {
	allDeals := k.GetAllActiveDeals(ctx)

	var space int64 = 0
	for i := 0; i < len(allDeals); i++ {
		deal := allDeals[i]
		if deal.Provider != provider {
			continue
		}
		size, ok := sdk.NewIntFromString(deal.Filesize)
		if !ok {
			continue
		}

		space += size.Int64()

	}

	return space
}
