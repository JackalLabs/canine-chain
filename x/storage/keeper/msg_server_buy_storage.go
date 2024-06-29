package keeper

import (
	"context"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
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
	params := k.GetParams(ctx)

	forAddress, err := k.rnsKeeper.Resolve(ctx, msg.ForAddress) // converting for address into an actual bech32 using RNS
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "cannot parse RNS or address %s", msg.ForAddress)
	}

	duration, bytes, gbs, denom, err := validateBuy(msg.DurationDays, msg.Bytes, msg.PaymentDenom)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to validate buy request")
	}

	hours := sdk.NewDec(duration.Milliseconds()).Quo(sdk.NewDec(60 * 60 * 1000))

	// We can calculate all the unit adjustments here. We will only consider the price adjustment when actually
	// buying storage, if you are upgrading, you get a discount based on the price you're upgrading not the entire storage
	storageCost := k.GetStorageCost(ctx, gbs, hours.TruncateInt().Int64())
	toPay := sdk.NewCoin(msg.PaymentDenom, storageCost)

	forAddr, err := sdk.AccAddressFromBech32(msg.ForAddress)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "for address is not a proper bech32")
	}

	accExists := k.accountKeeper.HasAccount(ctx, forAddr)
	if !accExists {
		defer telemetry.IncrCounter(1, "new", "account")
		k.accountKeeper.SetAccount(ctx, k.accountKeeper.NewAccountWithAddress(ctx, forAddr))
	}

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

	referred := false
	refAcc, err := k.rnsKeeper.Resolve(ctx, msg.Referral)
	if err == nil {
		if !(refAcc.String() == msg.Creator) {
			referred = true
		}
	}

	pol := sdk.NewDec(params.PolRatio).QuoInt64(100)
	discount := sdk.NewDec(0)
	fmt.Printf("POL: %d / %f\n", params.PolRatio, pol.MustFloat64())
	if referred {

		p := toPay.Amount.ToDec()

		var hour int64 = 1000 * 60 * 60
		if duration.Milliseconds() > 365*24*hour {
			p = p.Mul(sdk.MustNewDecFromStr("0.95"))
			discount = sdk.NewDec(5).QuoInt64(100) // 5% discount
			pol = pol.Sub(sdk.MustNewDecFromStr("0.05"))
		} else {
			p = p.Mul(sdk.MustNewDecFromStr("0.90"))
			discount = sdk.NewDec(10).QuoInt64(100) // 10% discount
			pol = pol.Sub(sdk.MustNewDecFromStr("0.1"))
		}

		toPay = sdk.NewCoin(toPay.Denom, p.TruncateInt())
	}

	spi = types.StoragePaymentInfo{
		Start:          ctx.BlockTime(),
		End:            ctx.BlockTime().Add(duration),
		SpaceAvailable: bytes,
		SpaceUsed:      spaceUsed,
		Address:        forAddress.String(),
	}

	add, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "cannot parse creator address %s", msg.Creator)
	}
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, add, types.ModuleName, sdk.NewCoins(toPay)) // taking money from user
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "cannot send tokens from %s", msg.Creator)
	}

	k.SetStoragePaymentInfo(ctx, spi)

	acc, err := types.GetTokenHolderAccount()
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "cannot get token holder account")
	}

	refDec := sdk.NewDec(params.ReferralCommission).QuoInt64(100)
	fmt.Printf("RATIOS!\nref: %d\npol: %d\ndiscount: %d\n", refDec.MulInt64(100).TruncateInt64(), pol.MulInt64(100).TruncateInt64(), discount.MulInt64(100).TruncateInt64())
	spr := sdk.NewDec(1).Sub(refDec).Sub(pol).Sub(discount) // whatever is left from pol and referrals

	// 1 - 0.25 = 0.75
	//

	fmt.Printf("storageprovider ratio: %d\n", spr.MulInt64(100).TruncateInt().Int64())

	storageProviderCut := toPay.Amount.ToDec().Mul(spr)
	spcToken := sdk.NewCoin(toPay.Denom, storageProviderCut.TruncateInt())
	spcTokens := sdk.NewCoins(spcToken)

	k.NewGauge(ctx, spcTokens, spi.End) // creating new payment gauge
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, acc, spcTokens)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "cannot send tokens to token holder account")
	}

	polAcc, err := types.GetPOLAccount()
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "cannot get pol account")
	}
	polCut := toPay.Amount.ToDec().Mul(pol) // 40,35,30% to pol
	polToken := sdk.NewCoin(toPay.Denom, polCut.TruncateInt())
	polTokens := sdk.NewCoins(polToken)

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, polAcc, polTokens)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "cannot send tokens to pol account")
	}

	refCut := toPay.Amount.ToDec().Mul(refDec) // 25% to referrals
	refToken := sdk.NewCoin(toPay.Denom, refCut.TruncateInt())
	refTokens := sdk.NewCoins(refToken)

	if referred {
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, refAcc, polTokens)
		if err != nil {
			return nil, sdkerrors.Wrapf(err, "cannot send tokens to referral account")
		}
	} else { // if we have no referral then we send the tokens to stakers
		err := k.AddCollectedFees(ctx, refTokens)
		if err != nil {
			return nil, sdkerrors.Wrapf(err, "cannot send tokens to stakers")
		}
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeBuyStorage,
			sdk.NewAttribute(types.AttributeKeyBuyer, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyReceiver, msg.ForAddress),
			sdk.NewAttribute(types.AttributeKeyBytesBought, fmt.Sprintf("%d", msg.Bytes)),
			sdk.NewAttribute(types.AttributeKeyTimeBought, hours.String()),
		),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeJackalMessage,
			sdk.NewAttribute(types.AttributeKeySigner, msg.Creator),
		),
	)

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
