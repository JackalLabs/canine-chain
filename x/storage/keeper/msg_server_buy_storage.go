package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerr "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

func (k msgServer) BuyStorage(goCtx context.Context, msg *types.MsgBuyStorage) (*types.MsgBuyStorageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.ForAddress)
	if err != nil {
		return nil, sdkerr.Wrap(err, "for address is poorly formatted")
	}

	duration := time.Duration(msg.DurationDays) * time.Hour * 24

	timeMonth := time.Hour * 24 * 30
	if duration < timeMonth {
		return nil, fmt.Errorf("duration can't be less than 1 month")
	}

	bytes := msg.Bytes

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
	storageCost := k.GetStorageCost(ctx, gbs, hours.TruncateInt().Int64())
	priceTokens := sdk.NewCoin(denom, storageCost)

	priceTokenList := sdk.NewCoins(priceTokens)

	add, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerr.Wrap(err, "creator address is poorly formatted")
	}
	err = k.bankkeeper.SendCoinsFromAccountToModule(ctx, add, types.ModuleName, priceTokenList)
	if err != nil {
		return nil, sdkerr.Wrap(err, "cannot send tokens from account")
	}

	deposit, err := sdk.AccAddressFromBech32(k.GetParams(ctx).DepositAccount)
	if err != nil {
		return nil, sdkerr.Wrap(err, "global deposit address is poorly formatted (most likely didn't get set properly in the genesis params)")
	}

	var spi types.StoragePaymentInfo

	payInfo, found := k.GetStoragePaymentInfo(ctx, msg.ForAddress)
	if found {

		if payInfo.SpaceUsed > bytes {
			return nil, fmt.Errorf("cannot buy less than your current gb usage")
		}

		if payInfo.End.After(ctx.BlockTime()) {

			err := k.UpgradeStorage(ctx, msg.Creator, bytes, payInfo, duration, storageCost, denom)
			if err != nil {
				return nil, sdkerr.Wrap(err, "failed to upgrade storage")
			}
			return &types.MsgBuyStorageResponse{}, nil
		}

		c := payInfo.Coins.Add(priceTokens)

		spi = types.StoragePaymentInfo{
			Start:          ctx.BlockTime(),
			End:            ctx.BlockTime().Add(duration),
			SpaceAvailable: bytes,
			SpaceUsed:      payInfo.SpaceUsed,
			Address:        msg.ForAddress,
			Coins:          c,
		}
	} else {
		spi = types.StoragePaymentInfo{
			Start:          ctx.BlockTime(),
			End:            ctx.BlockTime().Add(duration),
			SpaceAvailable: bytes,
			SpaceUsed:      0,
			Address:        msg.ForAddress,
			Coins:          priceTokenList,
		}
	}

	k.SetStoragePaymentInfo(ctx, spi)

	err = k.bankkeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, deposit, sdk.NewCoins(priceTokens))
	if err != nil {
		return nil, err
	}

	return &types.MsgBuyStorageResponse{}, nil
}
