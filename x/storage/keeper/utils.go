package keeper

import (
	"fmt"
	"net/url"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/storage/types"
)

func ParseIP(ip string) string {
	u, err := url.ParseRequestURI(ip)
	if err != nil {
		panic(err)
	}
	fmt.Println(u)

	return fmt.Sprintf("%s%s%s", u.Scheme, u.Host, u.Path)
}

const (
	START_BLOCK_TYPE = "start"
	END_BLOCK_TYPE   = "end"
	TWO_GIGS         = 2000000000
)

func (k Keeper) GetPaidAmount(ctx sdk.Context, address string, blockh int64) (int64, bool, *types.PayBlocks) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PayBlocksKeyPrefix))

	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	var highestBlock int64 = 0

	eblock, found := k.GetPayBlocks(ctx, fmt.Sprintf(".%s", address))
	if !found {
		return TWO_GIGS, true, nil
	}

	endblock, ok := sdk.NewIntFromString(eblock.Blocknum)
	if !ok {
		return TWO_GIGS, true, nil
	}

	if endblock.Int64() <= blockh {
		if endblock.Int64() <= blockh+432000 {
			bytes, ok := sdk.NewIntFromString(eblock.Bytes)
			if !ok {
				return bytes.Int64(), true, nil
			}
		}
		return TWO_GIGS, true, &eblock
	}

	for ; iterator.Valid(); iterator.Next() {
		var val types.PayBlocks
		k.cdc.MustUnmarshal(iterator.Value(), &val)

		fmt.Printf("BLOCK %s: %s", val.Blocktype, val.Blocknum)

		if val.Blocktype == END_BLOCK_TYPE {
			continue
		}

		adr := val.Blockid[:42]
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
		return TWO_GIGS, true, &eblock
	}

	hblock, found := k.GetPayBlocks(ctx, fmt.Sprintf("%s%d", address, highestBlock))
	if !found {
		return TWO_GIGS, true, &eblock
	}

	bytes, ok := sdk.NewIntFromString(hblock.Bytes)
	if !ok {
		return TWO_GIGS, true, &eblock
	}

	return bytes.Int64(), false, &eblock
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

	amount, trial, _ := k.GetPaidAmount(ctx, address, startBlock)

	if !trial && bytes <= amount {
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
