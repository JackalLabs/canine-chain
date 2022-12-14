package keeper

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerr "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

const (
	timeMonth       = time.Hour * 24 * 30
	gb        int64 = 1_000_000_000
)

func (k Keeper) UpgradeStorage(goCtx context.Context, msg *types.MsgUpgradeStorage) (*types.MsgUpgradeStorageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.ForAddress)
	if err != nil {
		return nil, err
	}

	payInfo, found := k.GetStoragePaymentInfo(ctx, msg.ForAddress)
	if !found {
		return nil, fmt.Errorf("can't upgrade non-existing storage, please use MsgBuyStorage")
	}

	duration, err := strconv.ParseInt(msg.Duration, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("duration can't be parsed: %s", err.Error())
	}

	if duration <= 0 {
		return nil, fmt.Errorf("duration can't be less than 1 month ")
	}

	timeTotal := timeMonth * time.Duration(duration)

	proratedDurationInHour := payInfo.End.Sub(ctx.BlockTime())
	proratedDuration := sdk.NewDec(int64(proratedDurationInHour)).Quo(sdk.NewDec(int64(timeMonth)))

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

	pricePerGB := sdk.MustNewDecFromStr("0.008")

	pricePerMonth := pricePerGB.Mul(sdk.NewDec(gbs))
	currentPricePerMonth := pricePerGB.Mul(sdk.NewDec(currentGbs))

	priceBefore := pricePerMonth.Mul(sdk.NewDec(duration))
	proratedRefund := currentPricePerMonth.Mul(proratedDuration)
	price := priceBefore.Sub(proratedRefund)

	if price.IsNegative() {
		return nil, fmt.Errorf("cannot downgrade at the moment, please wait till your subscription is over")
	}

	jklPrice, err := sdk.NewDecFromStr("0.20")
	if err != nil {
		return nil, err
	}
	feed, found := k.oraclekeeper.GetFeed(ctx, "jklPrice")
	if found {
		type data struct {
			Price  string `json:"price"`
			Change string `json:"24h_change"`
		}
		var d data
		err = json.Unmarshal([]byte(feed.Data), &d)
		if err != nil {
			return nil, err
		}

		jklPrice, err = sdk.NewDecFromStr(d.Price)
		if err != nil {
			return nil, err
		}
	}

	jklTokens := price.Quo(jklPrice)

	ujklTokens := jklTokens.Mul(sdk.NewDec(1_000_000))

	priceTokens := sdk.NewCoin(denom, ujklTokens.TruncateInt())

	add, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	err = k.bankkeeper.SendCoinsFromAccountToModule(ctx, add, types.ModuleName, sdk.NewCoins(priceTokens))
	if err != nil {
		return nil, err
	}

	depositAccount, err := sdk.AccAddressFromBech32(k.GetParams(ctx).DepositAccount)
	if err != nil {
		return nil, err
	}

	spi := types.StoragePaymentInfo{
		Start:          ctx.BlockTime(),
		End:            ctx.BlockTime().Add(timeTotal),
		SpaceAvailable: bytes,
		SpaceUsed:      payInfo.SpaceUsed,
		Address:        msg.ForAddress,
	}

	k.SetStoragePaymentInfo(ctx, spi)

	err = k.bankkeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, depositAccount, sdk.NewCoins(priceTokens))
	if err != nil {
		return nil, err
	}

	return &types.MsgUpgradeStorageResponse{}, nil
}
