package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

const (
	StartBlockType = "start"
	EndBlockType   = "end"
	TwoGigs        = 2000000000
)

func (k Keeper) GetPaidAmount(ctx sdk.Context, address string, blockh int64) (int64, bool, *types.PayBlocks) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PayBlocksKeyPrefix))

	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	var highestBlock int64

	eblock, found := k.GetPayBlocks(ctx, fmt.Sprintf(".%s", address))
	if !found {
		return TwoGigs, true, nil
	}

	endblock, ok := sdk.NewIntFromString(eblock.Blocknum)
	if !ok {
		return TwoGigs, true, nil
	}

	if endblock.Int64() <= blockh {
		// one month grace period
		if blockh - endblock.Int64() <= 432000 {
			bytes, ok := sdk.NewIntFromString(eblock.Bytes)
			if ok {
				return bytes.Int64(), true, nil
			}
		}
		return TwoGigs, true, &eblock
	}

	highestBlock = 0

	// Look for highest start block
	for ; iterator.Valid(); iterator.Next() {
		var val types.PayBlocks
		k.cdc.MustUnmarshal(iterator.Value(), &val)

		ctx.Logger().Debug("BLOCK %s: %s", val.Blocktype, val.Blocknum)

		if val.Blocktype == EndBlockType {
			continue
		}

		adr := val.Blockid[:len(address)]
		if adr != address {
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
			ctx.Logger().Debug(fmt.Sprintf("NEW HIGHEST BLOCK: %s", val.Blocknum))
		}

	}

	if highestBlock == 0 {
		return TwoGigs, true, &eblock
	}

	hblock, found := k.GetPayBlocks(ctx, fmt.Sprintf("%s%d", address, highestBlock))
	if !found {
		return TwoGigs, true, &eblock
	}

	bytes, ok := sdk.NewIntFromString(hblock.Bytes)
	if !ok {
		return TwoGigs, true, &eblock
	}

	return bytes.Int64(), false, &eblock
}

func (k Keeper) CreatePayBlock(ctx sdk.Context, address string, length int64, bytes int64) error {
	startBlock := ctx.BlockHeight()

	endBlock := startBlock + length

	sBlock := types.PayBlocks{
		Blockid:   fmt.Sprintf("%s%d", address, startBlock),
		Bytes:     fmt.Sprintf("%d", bytes),
		Blocktype: StartBlockType,
		Blocknum:  fmt.Sprintf("%d", startBlock),
	}

	eBlock := types.PayBlocks{
		Blockid:   fmt.Sprintf(".%s", address),
		Bytes:     fmt.Sprintf("%d", bytes),
		Blocktype: EndBlockType,
		Blocknum:  fmt.Sprintf("%d", endBlock),
	}

	amount, trial, _ := k.GetPaidAmount(ctx, address, startBlock)

	if !trial && bytes <= amount { // Not in trial and new storage space is
		// smaller than already paid amount
		return fmt.Errorf("can't buy storage within another storage window")
	}

	k.SetPayBlocks(ctx, sBlock)
	k.SetPayBlocks(ctx, eBlock)

	return nil
}

func (k Keeper) GetProviderUsing(ctx sdk.Context, provider string) int64 {
	allDeals := k.GetAllActiveDeals(ctx)

	var space int64
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
