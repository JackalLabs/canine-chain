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

func (k msgServer) BuyStorage(goCtx context.Context, msg *types.MsgBuyStorage) (*types.MsgBuyStorageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.ForAddress)
	if err != nil {
		return nil, err
	}

	duration, err := strconv.ParseInt(msg.Duration, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("duration can't be parsed: %s", err.Error())
	}
	if duration <= 0 {
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

	pricePerGB := sdk.MustNewDecFromStr("0.008")

	pricePerMonth := pricePerGB.Mul(sdk.NewDec(gbs))

	price := pricePerMonth.Mul(sdk.NewDec(duration))

	jklPrice, err := sdk.NewDecFromStr("0.20")
	if err != nil {
		return nil, err
	}
	feed, found := k.oraclekeeper.GetFeed(ctx, "jklprice")
	if found {
		type data struct {
			Price  float64 `json:"price"`
			Change float64 `json:"24h_change"`
		}
		var d data
		err = json.Unmarshal([]byte(feed.Data), &d)
		if err != nil {
			return nil, err
		}

		jklPrice, err = sdk.NewDecFromStr(fmt.Sprintf("%f", d.Price))
		if err != nil {
			return nil, err
		}

	}

	jklTokens := price.Quo(jklPrice)

	ujklTokens := jklTokens.Mul(sdk.NewDec(1000000)) // converting jkl to ujkl

	priceTokens := sdk.NewCoin(denom, ujklTokens.TruncateInt())

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

	timeMonth := time.Hour * 24 * 30

	timeTotal := timeMonth * time.Duration(duration)

	var spi types.StoragePaymentInfo

	payInfo, found := k.GetStoragePaymentInfo(ctx, msg.ForAddress)
	if found {

		if payInfo.SpaceUsed > bytes {
			return nil, fmt.Errorf("cannot buy less than your current gb usage")
		}

		spi = types.StoragePaymentInfo{
			Start:          ctx.BlockTime(),
			End:            ctx.BlockTime().Add(timeTotal),
			SpaceAvailable: bytes,
			SpaceUsed:      payInfo.SpaceUsed,
			Address:        msg.ForAddress,
		}
	} else {
		spi = types.StoragePaymentInfo{
			Start:          ctx.BlockTime(),
			End:            ctx.BlockTime().Add(timeTotal),
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

	return &types.MsgBuyStorageResponse{}, nil
}
