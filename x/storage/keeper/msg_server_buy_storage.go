package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/storage/types"
)

func (k msgServer) BuyStorage(goCtx context.Context, msg *types.MsgBuyStorage) (*types.MsgBuyStorageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sblock, ok := sdk.NewIntFromString(msg.StartBlock)
	if !ok {
		return nil, fmt.Errorf("start block can't be parsed")
	}

	duration, ok := sdk.NewIntFromString(msg.Duration)
	if !ok {
		return nil, fmt.Errorf("duration can't be parsed")
	}

	bytes, ok := sdk.NewIntFromString(msg.Bytes)
	if !ok {
		return nil, fmt.Errorf("bytes can't be parsed")
	}

	err := k.CreatePayBlock(ctx, msg.ForAddress, sblock.Int64(), duration.Int64(), bytes.Int64())

	if err != nil {
		return nil, err
	}

	return &types.MsgBuyStorageResponse{}, nil
}
