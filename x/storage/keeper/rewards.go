package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerror "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func getTotalSize(allDeals []types.ActiveDealsV2) sdk.Dec {
	networkSize := sdk.NewDecFromInt(sdk.NewInt(0))
	for i := 0; i < len(allDeals); i++ {
		deal := allDeals[i]
		ss := sdk.NewDecFromInt(sdk.NewInt(deal.FileSize))

		networkSize = networkSize.Add(ss)
	}
	return networkSize
}

func (k Keeper) manageDealReward(ctx sdk.Context, deal types.ActiveDealsV2, networkSize sdk.Dec, balance sdk.Coin) error {
	toprove := sdk.NewInt(deal.BlockToProve)

	iprove := toprove.Int64()

	totalSize := sdk.NewDecFromInt(sdk.NewInt(deal.FileSize))

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

	deal.BlockToProve = iprove

	if !deal.ProofVerified {
		ctx.Logger().Debug("%s\n", "Not verified!")

		DayBlocks := k.GetParams(ctx).ProofWindow

		if deal.StartBlock >= ctx.BlockHeight()-DayBlocks {
			return sdkerror.Wrapf(sdkerror.ErrUnauthorized, "ignore young deals")
		}

		misses := deal.ProofsMissed + 1
		missesToBurn := k.GetParams(ctx).MissesToBurn

		if misses > missesToBurn {
			provider, ok := k.GetProviders(ctx, deal.Provider)
			if !ok {
				return sdkerror.Wrapf(sdkerror.ErrKeyNotFound, "provider not found")
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
				Signee:   deal.Signer,
				Filesize: fmt.Sprintf("%d", deal.FileSize),
				Merkle:   deal.Merkle,
				End:      deal.EndBlock,
			}
			k.SetStrays(ctx, strayDeal)
			k.RemoveActiveDeals(ctx, deal.Cid)
			return nil
		}

		deal.ProofsMissed = misses
		k.SetActiveDeals(ctx, deal)
		return nil
	}

	ctx.Logger().Debug(fmt.Sprintf("File size: %d\n", deal.FileSize))
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

	updatedMisses := deal.ProofsMissed - 1

	if updatedMisses < 0 {
		updatedMisses = 0
	}

	deal.ProofsMissed = updatedMisses
	deal.ProofVerified = false
	k.SetActiveDeals(ctx, deal)

	return nil
}

func (k Keeper) loopDeals(ctx sdk.Context, allDeals []types.ActiveDealsV2, networkSize sdk.Dec, balance sdk.Coin) {
	for _, deal := range allDeals {
		info, found := k.GetStoragePaymentInfo(ctx, deal.Signer)
		if !found {
			ctx.Logger().Debug(fmt.Sprintf("Removing %s due to no payment info", deal.Cid))
			cerr := CanContract(ctx, deal.Cid, deal.Signer, k)
			if cerr != nil {
				ctx.Logger().Error(cerr.Error())
			}
			continue
		}
		grace := info.End.Add(time.Hour * 24 * 30)
		if grace.Before(ctx.BlockTime()) {
			ctx.Logger().Debug(fmt.Sprintf("Removing %s after grace period", deal.Cid))
			cerr := CanContract(ctx, deal.Cid, deal.Signer, k)
			if cerr != nil {
				ctx.Logger().Error(cerr.Error())
			}
			continue
		}

		if info.SpaceUsed > info.SpaceAvailable { // remove file if the user doesn't have enough space
			ctx.Logger().Debug(fmt.Sprintf("Removing %s for space used", deal.Cid))
			err := CanContract(ctx, deal.Cid, deal.Signer, k)
			if err != nil {
				ctx.Logger().Error(err.Error())
			}
			continue
		}

		err := k.manageDealReward(ctx, deal, networkSize, balance)
		if err != nil {
			ctx.Logger().Error(err.Error())
			continue
		}

	}
}

func (k Keeper) InternalRewards(ctx sdk.Context, allDeals []types.ActiveDealsV2, address sdk.AccAddress) error {
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

	DayBlocks := k.GetParams(ctx).ProofWindow

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
