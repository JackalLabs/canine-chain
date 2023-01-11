package keeper

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerror "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/jackalLabs/canine-chain/x/storage/types"
)

const (
	fchunks   int64 = 1024
	DayBlocks int64 = 10 * 5 // 10 blocks is about 1 minute
)

func getTotalSize(allDeals []types.ActiveDeals) sdk.Dec {
	networkSize := sdk.NewDecFromInt(sdk.NewInt(0))
	for i := 0; i < len(allDeals); i++ {
		deal := allDeals[i]
		ss, err := sdk.NewDecFromStr(deal.Filesize)
		if err != nil {
			continue
		}
		networkSize = networkSize.Add(ss)
	}
	return networkSize
}

func (k Keeper) manageDealReward(ctx sdk.Context, deal types.ActiveDeals, networkSize sdk.Dec, balance sdk.Coin) error {
	toprove, ok := sdk.NewIntFromString(deal.Blocktoprove)
	if !ok {
		return sdkerror.Wrapf(sdkerror.ErrInvalidType, "int parse failed")
	}

	iprove := toprove.Int64()

	totalSize, err := sdk.NewDecFromStr(deal.Filesize)
	if err != nil {
		return err
	}

	var byteHash byte
	if len(ctx.HeaderHash().Bytes()) > 2 {
		byteHash = ctx.HeaderHash().Bytes()[0] + ctx.HeaderHash().Bytes()[1] + ctx.HeaderHash().Bytes()[2]
	} else {
		byteHash = byte(ctx.BlockHeight()) // support for running simulations
	}

	d := totalSize.TruncateInt().Int64() / fchunks

	if d > 0 {
		iprove = (int64(byteHash) + int64(ctx.BlockGasMeter().GasConsumed())) % d
	}

	deal.Blocktoprove = fmt.Sprintf("%d", iprove)

	verified, errb := strconv.ParseBool(deal.Proofverified)

	if errb != nil {
		return errb
	}

	if !verified {
		ctx.Logger().Debug("%s\n", "Not verified!")
		intt, ok := sdk.NewIntFromString(deal.Proofsmissed)
		if !ok {
			return sdkerror.Wrapf(sdkerror.ErrInvalidType, "int parse failed")
		}

		sb, ok := sdk.NewIntFromString(deal.Startblock)
		if !ok {
			return sdkerror.Wrapf(sdkerror.ErrInvalidType, "int parse failed")
		}

		if sb.Int64() >= ctx.BlockHeight()-DayBlocks {
			return sdkerror.Wrapf(sdkerror.ErrUnauthorized, "ignore young deals")
		}

		misses := intt.Int64() + 1
		const missesToBurn int64 = 3

		if misses > missesToBurn {
			provider, ok := k.GetProviders(ctx, deal.Provider)
			if !ok {
				return sdkerror.Wrapf(sdkerror.ErrInvalidType, "int parse failed")
			}

			curburn, ok := sdk.NewIntFromString(provider.BurnedContracts)
			if !ok {
				return sdkerror.Wrapf(sdkerror.ErrInvalidType, "int parse failed")
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
			return sdkerror.Wrapf(sdkerror.ErrInvalidType, "int parse failed")
		}

		deal.Proofsmissed = fmt.Sprintf("%d", misses)
		k.SetActiveDeals(ctx, deal)
		return sdkerror.Wrapf(sdkerror.ErrInvalidType, "int parse failed")
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
		return err
	}
	ctx.Logger().Debug("Sending coins to %s\n", provider.String())
	errorr := k.bankkeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, provider, coins)
	if errorr != nil {
		ctx.Logger().Debug("ERR: %v\n", errorr)
		ctx.Logger().Error(errorr.Error())
		return errorr
	}

	ctx.Logger().Debug("%s\n", deal.Cid)

	deal.Proofverified = "false"
	k.SetActiveDeals(ctx, deal)

	return nil
}

func (k Keeper) loopDeals(ctx sdk.Context, allDeals []types.ActiveDeals, networkSize sdk.Dec, balance sdk.Coin) {
	for _, deal := range allDeals {

		info, _ := k.GetStoragePaymentInfo(ctx, deal.Signee)
		if info.End.After(ctx.BlockTime()) {
			if info.SpaceUsed > TwoGigs {

				cerr := CanContract(ctx, deal.Cid, "admin", k)
				if cerr != nil {
					ctx.Logger().Error(cerr.Error())
				}
			}
			info.SpaceAvailable = TwoGigs
			k.SetStoragePaymentInfo(ctx, info)
		}

		err := k.manageDealReward(ctx, deal, networkSize, balance)
		if err != nil {
			ctx.Logger().Error(err.Error())
			continue
		}

	}
}

func (k Keeper) InternalRewards(ctx sdk.Context, allDeals []types.ActiveDeals, address sdk.AccAddress) error {
	ctx.Logger().Debug("%s\n", "checking blocks")

	networkSize := getTotalSize(allDeals)

	balance := k.bankkeeper.GetBalance(ctx, address, "ujkl")

	k.loopDeals(ctx, allDeals, networkSize, balance)

	balance = k.bankkeeper.GetBalance(ctx, address, "ujkl")

	err := k.bankkeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(balance))
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) HandleRewardBlock(ctx sdk.Context) error {
	allDeals := k.GetAllActiveDeals(ctx)

	ctx.Logger().Debug("blockdiff : %d\n", ctx.BlockHeight()%DayBlocks)

	if ctx.BlockHeight()%DayBlocks > 0 {
		return sdkerror.Wrapf(sdkerror.ErrUnauthorized, "cannot check rewards before timer has been met")
	}

	address := k.accountkeeper.GetModuleAddress(types.ModuleName)

	err := k.InternalRewards(ctx, allDeals, address)
	if err != nil {
		return err
	}

	return nil
}
