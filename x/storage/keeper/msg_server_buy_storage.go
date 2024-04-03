package keeper

import (
	"context"
	"fmt"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerr "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

func (k msgServer) BuyStorage(goCtx context.Context, msg *types.MsgBuyStorage) (*types.MsgBuyStorageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.ForAddress)
	if err != nil {
		return nil, err
	}

	duration, err := time.ParseDuration(msg.Duration)
	if err != nil {
		return nil, fmt.Errorf("duration can't be parsed: %s", err.Error())
	}

	timeMonth := time.Hour * 24 * 30
	if duration.Truncate(timeMonth) <= 0 {
		return nil, fmt.Errorf("duration can't be less than 1 month")
	}

	bytes, err := strconv.ParseInt(msg.Bytes, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("bytes can't be parsed: %s", err.Error())
	}

	denom := msg.PaymentDenom
	if denom != "ujkl" {
		return nil, sdkerr.Wrap(sdkerr.ErrInvalidCoins, "cannot pay with anything other than ujkl")
	}

	const gb int64 = 1000000000

	gbs := bytes / gb
	if gbs <= 0 {
		return nil, fmt.Errorf("cannot buy less than a gb")
	}

	hours := sdk.NewDec(duration.Milliseconds()).Quo(sdk.NewDec(60 * 60 * 1000))
	priceTokens := sdk.NewCoin(denom, k.GetStorageCost(ctx, gbs, hours.TruncateInt().Int64()))

	add, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	err = k.bankkeeper.SendCoinsFromAccountToModule(ctx, add, types.ModuleName, sdk.NewCoins(priceTokens))
	if err != nil {
		return nil, err
	}

	deposit, err := sdk.AccAddressFromBech32(k.GetParams(ctx).DepositAccount)
	if err != nil {
		return nil, err
	}

	var spi types.StoragePaymentInfo

	payInfo, found := k.GetStoragePaymentInfo(ctx, msg.ForAddress)
	if found {

		if payInfo.SpaceUsed > bytes {
			return nil, fmt.Errorf("cannot buy less than your current gb usage")
		}

		if payInfo.End.After(ctx.BlockTime()) {
			return nil, fmt.Errorf("please use MsgUpgradeStorage if you want to upgrade/downgrade")
		}

		spi = types.StoragePaymentInfo{
			Start:          ctx.BlockTime(),
			End:            ctx.BlockTime().Add(duration),
			SpaceAvailable: bytes,
			SpaceUsed:      payInfo.SpaceUsed,
			Address:        msg.ForAddress,
		}
	} else {
		spi = types.StoragePaymentInfo{
			Start:          ctx.BlockTime(),
			End:            ctx.BlockTime().Add(duration),
			SpaceAvailable: bytes,
			SpaceUsed:      0,
			Address:        msg.ForAddress,
		}
	}

	k.SetStoragePaymentInfo(ctx, spi)

	err = k.bankkeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, deposit, sdk.NewCoins(priceTokens))
	if err != nil {
		return nil, err
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
			sdk.NewAttribute(types.AttributeKeyBytesBought, msg.Bytes),
			sdk.NewAttribute(types.AttributeKeyTimeBought, hours.String()),
		),
	)

	return &types.MsgBuyStorageResponse{}, nil
}
