package keeper

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerror "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
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

	d := totalSize.TruncateInt().Int64() / k.GetParams(ctx).ChunkSize

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

		DayBlocks := k.GetParams(ctx).ProofWindow

		if sb.Int64() >= ctx.BlockHeight()-DayBlocks {
			ctx.Logger().Info("ignore young deals")
			return nil
		}

		misses := intt.Int64() + 1
		missesToBurn := k.GetParams(ctx).MissesToBurn

		if misses > missesToBurn {
			return k.DropDeal(ctx, deal, true)
		}

		deal.Proofsmissed = fmt.Sprintf("%d", misses)
		k.SetActiveDeals(ctx, deal)
		return nil
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

	misses, ok := sdk.NewIntFromString(deal.Proofsmissed)
	if !ok {
		e := errors.New("cannot parse string")
		ctx.Logger().Error(e.Error())
		return e
	}
	updatedMisses := misses.SubRaw(1)

	if updatedMisses.LT(sdk.NewInt(0)) {
		updatedMisses = sdk.NewInt(0)
	}

	deal.Proofsmissed = updatedMisses.String()
	deal.Proofverified = "false"
	k.SetActiveDeals(ctx, deal)

	ap := types.ActiveProviders{
		Address: deal.Provider,
	}

	k.SetActiveProviders(ctx, ap)

	return nil
}

func (k Keeper) loopDeals(ctx sdk.Context, allDeals []types.ActiveDeals, networkSize sdk.Dec, balance sdk.Coin) {
	currentBlock := ctx.BlockHeight()
	for _, deal := range allDeals {
		end, err := strconv.ParseInt(deal.Endblock, 10, 64)
		if err != nil {
			ctx.Logger().Error(err.Error())
			continue
		}

		info, found := k.GetStoragePaymentInfo(ctx, deal.Signee)
		if !found { // user has no storage plan, we'll check if the deal was made with no plan before removing it

			if end == 0 { // the deal was made with a plan yet the user has no plan, remove it
				ctx.Logger().Debug(fmt.Sprintf("Removing %s due to no payment info", deal.Cid))
				cerr := k.CanContract(ctx, deal.Cid, deal.Signee)
				if cerr != nil {
					ctx.Logger().Error(cerr.Error())
				}
				continue
			}
		}

		if end == 0 { // for deals that were made with a subscription, we remove them if there is not enough space in the plan
			if info.SpaceUsed > info.SpaceAvailable { // remove file if the user doesn't have enough space
				ctx.Logger().Debug(fmt.Sprintf("Removing %s for space used", deal.Cid))
				err := k.CanContract(ctx, deal.Cid, deal.Signee)
				if err != nil {
					ctx.Logger().Error(err.Error())
				}
				continue
			}
		}

		if currentBlock > end && end > 0 { // check if end block has passed and was made with a timed storage deal
			ctx.Logger().Info(fmt.Sprintf("deal has expired at %d", ctx.BlockHeight()))

			grace := info.End.Add(time.Hour * 24 * 30)
			if grace.Before(ctx.BlockTime()) {
				ctx.Logger().Debug(fmt.Sprintf("Removing %s after grace period", deal.Cid))
				cerr := k.CanContract(ctx, deal.Cid, deal.Signee)
				if cerr != nil {
					ctx.Logger().Error(cerr.Error())
				}
				continue
			}
		}

		err = k.manageDealReward(ctx, deal, networkSize, balance)
		if err != nil {
			ctx.Logger().Error(err.Error())
		}

	}
}

func (k Keeper) InternalRewards(ctx sdk.Context, allDeals []types.ActiveDeals, address sdk.AccAddress) error {
	ctx.Logger().Debug("%s\n", "checking blocks")

	k.RemoveAllActiveProviders(ctx) // clearing recent provider list

	networkSize := getTotalSize(allDeals)

	balance := k.bankkeeper.GetBalance(ctx, address, "ujkl")

	k.loopDeals(ctx, allDeals, networkSize, balance)

	balance = k.bankkeeper.GetBalance(ctx, address, "ujkl")

	err := k.bankkeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(balance))
	if err != nil {
		return err
	}

	k.RemoveAllAttestation(ctx)

	return nil
}

func (k Keeper) HandleRewardBlock(ctx sdk.Context) error {
	allDeals := k.GetAllActiveDeals(ctx)

	DayBlocks := k.GetParams(ctx).ProofWindow

	if ctx.BlockHeight()%DayBlocks > 0 {
		ctx.Logger().Debug("skipping reward handling for this block")
		return nil
	}

	address := k.accountkeeper.GetModuleAddress(types.ModuleName)

	err := k.InternalRewards(ctx, allDeals, address)
	if err != nil {
		return err
	}

	return nil
}
