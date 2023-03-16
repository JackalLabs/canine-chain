package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackal-dao/canine/x/lp/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EstimateSwapIn(
	goCtx context.Context,
	req *types.QueryEstimateSwapInRequest,
) (
	*types.QueryEstimateSwapInResponse,
	error,
) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	pool, found := k.GetLPool(ctx, req.PoolName)

	if !found {
		return nil, types.ErrLiquidityPoolNotFound
	}

	poolCoins := sdk.NewCoins(pool.Coins...)

	desiredCoins, err := sdk.ParseCoinNormalized(req.OutputCoins)

	if err != nil {
		return nil, sdkerrors.Wrapf(err, sdkerrors.ErrInvalidRequest.Error())
	}

	if desiredCoins.IsNegative() {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"swap out amount cannot be negative: %s",
			desiredCoins.Amount.String())
	}

	AMM, _ := types.GetAMM(pool.AMM_Id)

	estDeposit, err := AMM.EstimateDeposit(poolCoins, sdk.NewCoins(desiredCoins))

	if err != nil {
		return nil, err
	}

	return &types.QueryEstimateSwapInResponse{InputCoins: estDeposit[0]}, nil
}
