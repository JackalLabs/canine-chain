package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

const (
	timeMonth       = time.Hour * 24 * 30
	gb        int64 = 1_000_000_000
)

func validateBuy(days int64, bytesIn int64, denomIn string) (duration time.Duration, bytes int64, gbs int64, denom string, err error) {
	duration = time.Duration(days) * time.Hour * 24
	if duration < timeMonth {
		err = fmt.Errorf("duration can't be less than 1 month")
		return
	}

	bytes = bytesIn
	gbs = bytes / gb
	if gbs <= 0 {
		err = fmt.Errorf("cannot buy less than a gb")
		return
	}

	denom = denomIn
	if denomIn != "ujkl" {
		err = sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "cannot pay with anything other than ujkl")
		return
	}

	return
}

func (k msgServer) BuyStorage(goCtx context.Context, msg *types.MsgBuyStorage) (*types.MsgBuyStorageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	forAddress, err := k.rnsKeeper.Resolve(ctx, msg.ForAddress) // converting for address into an actual bech32 using RNS
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "cannot parse RNS or address %s", msg.ForAddress)
	}

	duration, bytes, gbs, denom, err := validateBuy(msg.DurationDays, msg.Bytes, msg.PaymentDenom)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to validate buy request")
	}

	hours := sdk.NewDec(duration.Milliseconds()).Quo(sdk.NewDec(60 * 60 * 1000))
	storageCost := k.GetStorageCost(ctx, gbs, hours.TruncateInt().Int64())
	toPay := sdk.NewCoin(msg.PaymentDenom, storageCost)

	var spi types.StoragePaymentInfo
	var spaceUsed int64 // default 0
	payInfo, found := k.GetStoragePaymentInfo(ctx, forAddress.String())
	if found {
		if payInfo.SpaceUsed > bytes {
			return nil, fmt.Errorf("cannot buy less than your current gb usage")
		}
		spaceUsed = payInfo.SpaceUsed

		if payInfo.End.After(ctx.BlockTime()) { // should we upgrade storage instead of buy fresh?
			toPay, err = k.UpgradeStorage(ctx, bytes, payInfo, duration, storageCost, denom)
			if err != nil {
				return nil, sdkerrors.Wrapf(err, "cannot upgrade storage")
			}
		}
	}

	spi = types.StoragePaymentInfo{
		Start:          ctx.BlockTime(),
		End:            ctx.BlockTime().Add(duration),
		SpaceAvailable: bytes,
		SpaceUsed:      spaceUsed,
		Address:        forAddress.String(),
	}

	priceTokenList := sdk.NewCoins(toPay)

	add, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "cannot parse creator address %s", msg.Creator)
	}
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, add, types.ModuleName, priceTokenList) // taking money from user
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "cannot send tokens from %s", msg.Creator)
	}

	k.NewGauge(ctx, priceTokenList, spi.End) // creating new payment gauge

	k.SetStoragePaymentInfo(ctx, spi)

	acc, err := types.GetTokenHolderAccount()
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "cannot get token holder account")
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, acc, priceTokenList)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "cannot send tokens to module from holder account")
	}

	return &types.MsgBuyStorageResponse{}, nil
}

func (k Keeper) UpgradeStorage(
	ctx sdk.Context,
	bytes int64,
	payInfo types.StoragePaymentInfo,
	duration time.Duration,
	storageCost sdk.Int,
	denom string,
) (sdk.Coin, error) {
	proratedDuration := payInfo.End.Sub(ctx.BlockTime())
	proratedDurationInHour := sdk.NewDec(proratedDuration.Milliseconds()).Quo(sdk.NewDec(60 * 60 * 1000))

	currentBytes := payInfo.SpaceAvailable
	currentGbs := currentBytes / gb

	oldCost := k.GetStorageCost(ctx, currentGbs, proratedDurationInHour.TruncateInt64())

	if duration.Truncate(timeMonth) <= 0 {
		return sdk.Coin{}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "duration can't be less than 1 month")
	}

	if bytes < payInfo.SpaceUsed {
		return sdk.Coin{}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "cannot downgrade below current usage")
	}

	newCost := storageCost

	price := newCost.Sub(oldCost)

	if price.LTE(sdk.ZeroInt()) {
		return sdk.Coin{}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "cannot downgrade until current plan expires")
	}

	priceTokens := sdk.NewCoin(denom, price)

	return priceTokens, nil
}
