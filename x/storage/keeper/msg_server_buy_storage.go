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

func (k Keeper) buyStorageInternal(ctx sdk.Context, forAddress string, duration time.Duration, bytes int64, paymentDenom string, creator string) error {
	_, err := sdk.AccAddressFromBech32(forAddress)
	if err != nil {
		return err
	}

	timeMonth := time.Hour * 24 * 30
	if duration.Truncate(timeMonth) <= 0 {
		return fmt.Errorf("duration can't be less than 1 month")
	}

	const gb int64 = 1000000000

	gbs := bytes / gb
	if gbs <= 0 {
		return fmt.Errorf("cannot buy less than a gb")
	}
	hours := sdk.NewDec(duration.Milliseconds()).Quo(sdk.NewDec(60 * 60 * 1000))

	cost, err := k.GetStorageCost(ctx, gbs, hours.TruncateInt().Int64(), paymentDenom)
	if err != nil {
		return sdkerr.Wrap(err, "cannot pay with anything other than ujkl or ujwl")
	}

	priceTokens := sdk.NewCoin(paymentDenom, cost)

	add, err := sdk.AccAddressFromBech32(creator)
	if err != nil {
		return err
	}

	deposit, err := sdk.AccAddressFromBech32(k.GetParams(ctx).DepositAccount)
	if err != nil {
		return err
	}

	err = k.bankkeeper.SendCoinsFromAccountToModule(ctx, add, types.ModuleName, sdk.NewCoins(priceTokens))
	if err != nil {
		return err
	}

	err = k.bankkeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, deposit, sdk.NewCoins(priceTokens))
	if err != nil {
		return err
	}

	return nil
}

func (k msgServer) BuyStorage(goCtx context.Context, msg *types.MsgBuyStorage) (*types.MsgBuyStorageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	bytes, err := strconv.ParseInt(msg.Bytes, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("bytes can't be parsed: %s", err.Error())
	}

	duration, err := time.ParseDuration(msg.Duration)
	if err != nil {
		return nil, fmt.Errorf("duration can't be parsed: %s", err.Error())
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

	err = k.buyStorageInternal(ctx, msg.ForAddress, duration, bytes, msg.PaymentDenom, msg.Creator)
	if err != nil {
		return nil, err
	}

	k.SetStoragePaymentInfo(ctx, spi)

	return &types.MsgBuyStorageResponse{}, nil
}

func (k msgServer) BuyStorageToken(goCtx context.Context, msg *types.MsgBuyStorageToken) (*types.MsgBuyStorageTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	account, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	duration, err := time.ParseDuration("720h")
	if err != nil {
		return nil, err
	}

	err = k.buyStorageInternal(ctx, msg.Creator, duration, 8_000_000_000_000*msg.Amount, msg.PaymentDenom, msg.Creator)
	if err != nil {
		return nil, err
	}

	jwlCoin := sdk.NewInt64Coin("ujwl", 1_000_000)
	coins := sdk.NewCoins(jwlCoin)

	err = k.bankkeeper.MintCoins(ctx, types.ModuleName, coins)
	if err != nil {
		return nil, err
	}

	err = k.bankkeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, account, coins)
	if err != nil {
		return nil, err
	}

	return &types.MsgBuyStorageTokenResponse{}, err
}
