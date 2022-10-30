package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/x/lp/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) MakeValidPair(
	goCtx context.Context,
	req *types.QueryMakeValidPairRequest,
) (*types.QueryMakeValidPairResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	pool, found := k.GetLPool(ctx, req.PoolName)

	if !found {
		return nil, types.ErrLiquidityPoolNotFound
	}

	deposit, err := sdk.ParseCoinNormalized(req.Coin)
	if err != nil {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrInvalidCoins,
			"Failed to parse coin",
		)
	}

	result, err := MakeValidPair(pool, deposit)
	if err != nil {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrInvalidRequest,
			"Error occured while estimating other coins to deposit",
		)
	}

	return &types.QueryMakeValidPairResponse{Coin: result[0]}, nil
}
