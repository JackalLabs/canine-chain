package keeper

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/jackalLabs/canine-chain/v5/x/storage/types"
)

// func MakeFid(data []byte) (string, error) {
//	return bech32.ConvertAndEncode(types.FidPrefix, data)
// }

const (
	False = "false"
	True  = "true"
)

func MakeCid(data []byte) (string, error) {
	return bech32.ConvertAndEncode(types.CidPrefix, data)
}

func (k Keeper) GetPaidAmount(ctx sdk.Context, address string) int64 {
	payInfo, found := k.GetStoragePaymentInfo(
		ctx,
		address,
	)
	if !found {
		return 0
	}

	return payInfo.SpaceAvailable
}

func (k Keeper) GetProviderDeals(ctx sdk.Context, provider string) int64 {
	allDeals, _ := k.GetAllProofsForProver(ctx, provider)

	return int64(len(allDeals))
}

func (k Keeper) GetProviderUsing(ctx sdk.Context, provider string) int64 {
	allDeals, _ := k.GetAllProofsForProver(ctx, provider)

	var space int64
	for _, proof := range allDeals {
		deal, found := k.GetFile(ctx, proof.Merkle, proof.Owner, proof.Start)
		if !found {
			continue
		}

		space += deal.FileSize

	}

	return space
}

// GetStorageCostKbs calculates storage cost in ujkl
// Uses kilobytes and months to calculate how much user has to pay
func (k Keeper) GetStorageCostKbs(ctx sdk.Context, kbs int64, hours int64) sdk.Int {
	return k.GetStorageCostKbsWithPrice(ctx, kbs, hours, k.GetParams(ctx).PricePerTbPerMonth)
}

// GetStorageCostKbs calculates storage cost in ujkl
// Uses kilobytes and months to calculate how much user has to pay
func (k Keeper) GetStorageCostKbsWithPrice(ctx sdk.Context, kbs int64, hours int64, pricePerTBMonth int64) sdk.Int {
	pricePerTBPerMonth := sdk.NewDec(pricePerTBMonth)
	quantifiedPricePerTBPerMonth := pricePerTBPerMonth.QuoInt64(3)
	pricePerGbPerMonth := quantifiedPricePerTBPerMonth.QuoInt64(1000)
	pricePerMbPerMonth := pricePerGbPerMonth.QuoInt64(1000)
	pricePerKbPerMonth := pricePerMbPerMonth.QuoInt64(1000)
	pricePerKbPerHour := pricePerKbPerMonth.QuoInt64(720)

	pricePerHour := pricePerKbPerHour.MulInt64(kbs)

	totalCost := pricePerHour.MulInt64(hours)

	jklPrice := k.GetJklPrice(ctx)

	// TODO: fetch denom unit from bank module
	var ujklUnit int64 = 1000000
	jklCost := totalCost.Quo(jklPrice)

	ujklCost := jklCost.MulInt64(ujklUnit)

	return ujklCost.TruncateInt()
}

// GetStorageCost calculates storage cost in ujkl
// Uses gigabytes and months to calculate how much user has to pay
func (k Keeper) GetStorageCost(ctx sdk.Context, gbs int64, hours int64) sdk.Int {
	basePricePerTBPerMonth := sdk.NewDec(k.GetParams(ctx).PricePerTbPerMonth)
	basePricePerTBPerMonthYearly := basePricePerTBPerMonth.Mul(sdk.MustNewDecFromStr("12.5").QuoInt64(15)) // we only really care about the ratio in case the price changes

	var finalPricePerTbPerMonth sdk.Dec

	if hours < 365*24 { // calculating monthly
		switch {
		case gbs >= 20_000:
			finalPricePerTbPerMonth = basePricePerTBPerMonth.Mul(sdk.MustNewDecFromStr("12.5").QuoInt64(15)) // we only really care about the ratio in case the price changes
		case gbs >= 5_000:
			finalPricePerTbPerMonth = basePricePerTBPerMonth.Mul(sdk.NewDec(14).QuoInt64(15)) // we only really care about the ratio in case the price changes
		default:
			finalPricePerTbPerMonth = basePricePerTBPerMonth
		}
	} else { // calculating yearly
		switch {
		case gbs >= 20_000:
			finalPricePerTbPerMonth = basePricePerTBPerMonthYearly.Mul(sdk.MustNewDecFromStr("10.42").Quo(sdk.MustNewDecFromStr("12.5"))) // we only really care about the ratio in case the price changes
		case gbs >= 5_000:
			finalPricePerTbPerMonth = basePricePerTBPerMonthYearly.Mul(sdk.MustNewDecFromStr("11.67").Quo(sdk.MustNewDecFromStr("12.5"))) // we only really care about the ratio in case the price changes
		default:
			finalPricePerTbPerMonth = basePricePerTBPerMonthYearly
		}
	}

	quantifiedPricePerTBPerMonth := finalPricePerTbPerMonth.QuoInt64(3)
	pricePerGbPerMonth := quantifiedPricePerTBPerMonth.QuoInt64(1000)
	pricePerGbPerHour := pricePerGbPerMonth.QuoInt64(720)

	pricePerHour := pricePerGbPerHour.MulInt64(gbs)

	totalCost := pricePerHour.MulInt64(hours)

	jklPrice := k.GetJklPrice(ctx)

	// TODO: fetch denom unit from bank module
	var ujklUnit int64 = 1000000
	jklCost := totalCost.Quo(jklPrice)

	ujklCost := jklCost.MulInt64(ujklUnit)

	return ujklCost.TruncateInt()
}

// GetJklPrice uses oracle module to get jkl price
// Returns 0.20 if feed doesn't exist or failed to unmarshal data
// Unmarshal failure is logged
func (k Keeper) GetJklPrice(ctx sdk.Context) (price sdk.Dec) {
	price = sdk.MustNewDecFromStr("0.20")

	priceFeed := k.GetParams(ctx).PriceFeed
	feed, found := k.oracleKeeper.GetFeed(ctx, priceFeed)
	if found {
		type data struct {
			Price  string `json:"price"`
			Change string `json:"24h_change"`
		}

		var d data
		err := json.Unmarshal([]byte(feed.Data), &d)
		if err != nil {
			ctx.Logger().Debug("Failed to unmarshal Feed.Data (%s)", feed.Data)
		}

		p, err := sdk.NewDecFromStr(d.Price)
		if err != nil {
			ctx.Logger().Debug("Failed to convert Feed.Data.Price to sdk.Dec (%s)", d.Price)
		} else {
			price = p
		}
	}

	return price
}
