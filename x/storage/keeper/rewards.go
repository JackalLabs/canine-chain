package keeper

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func (k Keeper) HandleRewardBlock(ctx sdk.Context) {
	allDeals := k.GetAllActiveDeals(ctx)

	height := ctx.BlockHeight()

	const fchunks int64 = 1024

	var dayBlocks int64 = 10 * 5 // 10 blocks is about 1 minute

	ctx.Logger().Debug("blockdiff : %d\n", height%dayBlocks)
	if height%dayBlocks == 0 {
		ctx.Logger().Debug("%s\n", "checking blocks")

		networkSize := sdk.NewDecFromInt(sdk.NewInt(0))
		for i := 0; i < len(allDeals); i++ {
			deal := allDeals[i]
			ss, err := sdk.NewDecFromStr(deal.Filesize)
			if err != nil {
				continue
			}
			networkSize = networkSize.Add(ss)
		}

		address := k.accountkeeper.GetModuleAddress(types.ModuleName)
		balance := k.bankkeeper.GetBalance(ctx, address, "ujkl")

		for i := 0; i < len(allDeals); i++ {
			deal := allDeals[i]

			toprove, ok := sdk.NewIntFromString(deal.Blocktoprove)
			if !ok {
				ctx.Logger().Debug("Int Parse Failed!\n")
				continue
			}

			iprove := toprove.Int64()

			totalSize, err := sdk.NewDecFromStr(deal.Filesize)
			if err != nil {
				ctx.Logger().Error(err.Error())
				continue
			}

			byteHash := ctx.HeaderHash().Bytes()[0] + ctx.HeaderHash().Bytes()[1] + ctx.HeaderHash().Bytes()[2]

			d := totalSize.TruncateInt().Int64() / fchunks

			if d > 0 {
				iprove = (iprove + ctx.BlockHeight()*int64(byteHash)) % d
			}

			deal.Blocktoprove = fmt.Sprintf("%d", iprove)

			verified, errb := strconv.ParseBool(deal.Proofverified)

			if errb != nil {
				ctx.Logger().Debug("ERR %v\n", errb)
				continue
			}

			if !verified {
				ctx.Logger().Debug("%s\n", "Not verified!")
				intt, ok := sdk.NewIntFromString(deal.Proofsmissed)
				if !ok {
					ctx.Logger().Debug("Int Parse Failed!\n")
					continue
				}

				sb, ok := sdk.NewIntFromString(deal.Startblock)
				if !ok {
					ctx.Logger().Debug("Int Parse Failed!\n")
					continue
				}

				if sb.Int64() >= height-dayBlocks {
					continue
				}

				misses := intt.Int64() + 1
				const missesToBurn int64 = 3

				if misses > missesToBurn {
					provider, ok := k.GetProviders(ctx, deal.Provider)
					if !ok {
						continue
					}

					curburn, ok := sdk.NewIntFromString(provider.BurnedContracts)
					if !ok {
						continue
					}
					provider.BurnedContracts = fmt.Sprintf("%d", curburn.Int64()+1)
					k.SetProviders(ctx, provider)

					// Creating new stray file from the burned active deal
					strayDeal := types.Strays{
						Cid:      deal.Cid,
						Fid:      deal.Fid,
						Signee:   deal.Signee,
						Filesize: deal.Filesize,
						Merkle:   deal.Merkle,
					}
					k.SetStrays(ctx, strayDeal)
					k.RemoveActiveDeals(ctx, deal.Cid)
					continue
				}

				deal.Proofsmissed = fmt.Sprintf("%d", misses)
				k.SetActiveDeals(ctx, deal)
				continue
			}

			ctx.Logger().Debug(fmt.Sprintf("File size: %s\n", deal.Filesize))
			ctx.Logger().Debug(fmt.Sprintf("Total size: %d\n", networkSize))

			nom := totalSize

			den := networkSize

			res := nom.Quo(den)

			ctx.Logger().Debug("Percentage of network space * 1000: %f\n", res)

			coinfloat := res.Mul(balance.Amount.ToDec())

			ctx.Logger().Debug("%f\n", coinfloat)
			coin := sdk.NewCoin("ujkl", coinfloat.TruncateInt())
			coins := sdk.NewCoins(coin)

			provider, err := sdk.AccAddressFromBech32(deal.Provider)
			if err != nil {
				ctx.Logger().Error(err.Error())
				continue
			}
			ctx.Logger().Debug("Sending coins to %s\n", provider.String())
			errorr := k.bankkeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, provider, coins)
			if errorr != nil {
				ctx.Logger().Debug("ERR: %v\n", errorr)
				ctx.Logger().Error(errorr.Error())
				continue
			}

			ctx.Logger().Debug("%s\n", deal.Cid)

			deal.Proofverified = "false"
			k.SetActiveDeals(ctx, deal)

		}
		balance = k.bankkeeper.GetBalance(ctx, k.accountkeeper.GetModuleAddress(types.ModuleName), "ujkl")

		err := k.bankkeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(balance))
		if err != nil {
			ctx.Logger().Error("ERR: %v\n", err)
			return
		}
	}
}
