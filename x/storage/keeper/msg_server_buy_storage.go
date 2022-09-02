package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/storage/types"
)

func (k msgServer) BuyStorage(goCtx context.Context, msg *types.MsgBuyStorage) (*types.MsgBuyStorageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	duration, ok := sdk.NewIntFromString(msg.Duration)
	if !ok {
		return nil, fmt.Errorf("duration can't be parsed")
	}

	bytes, ok := sdk.NewIntFromString(msg.Bytes)
	if !ok {
		return nil, fmt.Errorf("bytes can't be parsed")
	}

	denom := msg.PaymentDenom
	var gb int64 = 1000000000

	gbs := bytes.Int64() / gb
	if gbs == 0 {
		return nil, fmt.Errorf("cannot buy less than a gb")
	}
	price := sdk.NewCoin(denom, sdk.NewInt(gbs*8000))
	add, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	err = k.bankkeeper.SendCoinsFromAccountToModule(ctx, add, types.ModuleName, sdk.NewCoins(price))
	if err != nil {
		return nil, err
	}

	err = k.CreatePayBlock(ctx, msg.ForAddress, duration.Int64(), bytes.Int64())

	if err != nil {
		return nil, err
	}

	return &types.MsgBuyStorageResponse{}, nil
}
