package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackal-dao/canine/x/lp/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EstimateContribution(
	goCtx context.Context,
	req *types.QueryEstimateContributionRequest,
) (
	*types.QueryEstimateContributionResponse,
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

	desiredAmt, ok := sdk.NewIntFromString(req.DesiredAmount)

	if !ok {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrInvalidRequest,
			"cannot convert desired amount into type sdk.Int: %s",
			req.DesiredAmount,
		)
	}

	result, err := CoinsToDepositForLPToken(pool, desiredAmt)

	if err != nil {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrInvalidRequest,
			"Failed to calculate deposit coins for lp token",
		)
	}

	return &types.QueryEstimateContributionResponse{Coins: result}, nil
}
