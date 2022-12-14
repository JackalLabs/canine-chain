package keeper

import (
	"context"
	"fmt"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerr "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

const (
	// replace costPerMonth when oracle/price is ready
	// cost below is for a gb
	costPerMonth = 4000
	costPerDay   = costPerMonth / 30

	gb           int64 = 1_000_000_000
	hoursInMonth       = time.Hour * 720
)

func (k Keeper) UpgradeStorage(goCtx context.Context, msg *types.MsgUpgradeStorage) (*types.MsgUpgradeStorageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	payInfo, found := k.GetStoragePaymentInfo(ctx, msg.ForAddress)
	if !found {
		return nil, fmt.Errorf("can't upgrade non-existing storage, please use MsgBuyStorage")
	}

	duration, err := time.ParseDuration(msg.Duration)
	if err != nil {
		return nil, fmt.Errorf("duration can't be parsed: %s", err.Error())
	}
	duration = duration.Truncate(time.Hour)
	if duration < hoursInMonth {
		return nil, fmt.Errorf("cannot buy less than a month (720h)")
	}
	dm := duration.Truncate(hoursInMonth)
	durationInMonth := int64(dm / hoursInMonth)

	proratedDuration := payInfo.End.Sub(ctx.BlockTime())
	pdInMonth := int64(proratedDuration / hoursInMonth)

	bytes, err := strconv.ParseInt(msg.Bytes, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("bytes can't be parsed: %s", err.Error())
	}
	currentBytes := payInfo.SpaceAvailable
	currentGbs := currentBytes / gb

	if bytes == currentBytes {
		return nil, fmt.Errorf("cannot upgrade to the same SpaceAvailable of %d GB", bytes/gb)
	}

	gbs := bytes / gb
	if gbs <= 0 {
		return nil, fmt.Errorf("cannot buy less than a gb")
	}

	denom := msg.PaymentDenom
	if denom != "ujkl" {
		return nil, sdkerr.Wrap(sdkerr.ErrInvalidCoins, "cannot pay with anything other than ujkl")
	}

	cost := gbs * costPerMonth * durationInMonth
	proratedRefund := currentGbs * costPerMonth * pdInMonth
	finalCost := cost - proratedRefund

	if finalCost < 0 {
		return nil, fmt.Errorf("cannot downgrade at the moment, please wait till your subscription is over")
	}

	payment := sdk.NewCoin(denom, sdk.NewInt(finalCost))
	add, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	err = k.bankkeeper.SendCoinsFromAccountToModule(ctx, add, types.ModuleName, sdk.NewCoins(payment))
	if err != nil {
		return nil, err
	}

	depositAccount, err := sdk.AccAddressFromBech32(k.GetParams(ctx).DepositAccount)
	if err != nil {
		return nil, err
	}

	err = k.bankkeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, depositAccount, sdk.NewCoins(payment))
	if err != nil {
		return nil, err
	}

	_, err = sdk.AccAddressFromBech32(msg.ForAddress)
	if err != nil {
		return nil, err
	}

	spi := types.StoragePaymentInfo{
		Start:          ctx.BlockTime(),
		End:            ctx.BlockTime().Add(dm),
		SpaceAvailable: bytes,
		SpaceUsed:      payInfo.SpaceUsed,
		Address:        msg.ForAddress,
	}

	k.SetStoragePaymentInfo(ctx, spi)

	return &types.MsgUpgradeStorageResponse{}, nil
}
