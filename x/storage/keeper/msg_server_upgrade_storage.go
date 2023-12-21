package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerr "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

const (
	timeMonth       = time.Hour * 24 * 30
	gb        int64 = 1_000_000_000
)

func (k Keeper) UpgradeStorage(ctx sdk.Context, creator string, bytes int64, payInfo types.StoragePaymentInfo, duration time.Duration, storageCost sdk.Int, denom string) error {
	proratedDuration := payInfo.End.Sub(ctx.BlockTime())
	proratedDurationInHour := sdk.NewDec(proratedDuration.Milliseconds()).Quo(sdk.NewDec(60 * 60 * 1000))

	currentBytes := payInfo.SpaceAvailable
	currentGbs := currentBytes / gb

	oldCost := k.GetStorageCost(ctx, currentGbs, proratedDurationInHour.TruncateInt64())

	if duration.Truncate(timeMonth) <= 0 {
		return sdkerr.Wrap(sdkerr.ErrInvalidRequest, "duration can't be less than 1 month")
	}

	newGbs := bytes / gb
	if newGbs <= 0 {
		return sdkerr.Wrap(sdkerr.ErrInvalidRequest, "cannot buy less than a gb")
	}

	if bytes < payInfo.SpaceUsed {
		return sdkerr.Wrap(sdkerr.ErrInvalidRequest, "cannot downgrade below current usage")
	}

	newCost := storageCost

	price := newCost.Sub(oldCost)

	if price.LTE(sdk.ZeroInt()) {
		return sdkerr.Wrap(sdkerr.ErrInvalidRequest, "cannot downgrade until current plan expires")
	}

	priceTokens := sdk.NewCoin(denom, price)

	add, err := sdk.AccAddressFromBech32(creator)
	if err != nil {
		return err
	}
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, add, types.ModuleName, sdk.NewCoins(priceTokens))
	if err != nil {
		return err
	}

	depositAccount, err := sdk.AccAddressFromBech32(k.GetParams(ctx).DepositAccount)
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, depositAccount, sdk.NewCoins(priceTokens))
	if err != nil {
		return err
	}

	spi := types.StoragePaymentInfo{
		Start:          ctx.BlockTime(),
		End:            ctx.BlockTime().Add(duration),
		SpaceAvailable: bytes,
		SpaceUsed:      payInfo.SpaceUsed,
		Address:        payInfo.Address,
	}

	k.SetStoragePaymentInfo(ctx, spi)

	return nil
}
