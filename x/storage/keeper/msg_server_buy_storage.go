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

func (k msgServer) BuyStorage(goCtx context.Context, msg *types.MsgBuyStorage) (*types.MsgBuyStorageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	duration, err := time.ParseDuration(msg.Duration)
	if err != nil {
		return nil, fmt.Errorf("duration can't be parsed: %s", err.Error())
	}

	// Truncate duration into hours
	dh := time.Hour
	duration = duration.Truncate(dh)

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

	const hoursInMonth = time.Hour * 720
	if duration <= hoursInMonth {
		return nil, fmt.Errorf("cannot buy less than a month(720h)")
	}

	// Truncate month
	dm := duration.Truncate(hoursInMonth)

	cost := gbs * 4000 * int64(dm/hoursInMonth)

	price := sdk.NewCoin(denom, sdk.NewInt(cost))
	add, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	err = k.bankkeeper.SendCoinsFromAccountToModule(ctx, add, types.ModuleName, sdk.NewCoins(price))
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
		SpaceUsed:      0,
		Address:        msg.ForAddress,
	}

	k.SetStoragePaymentInfo(ctx, spi)

	return &types.MsgBuyStorageResponse{}, nil
}
