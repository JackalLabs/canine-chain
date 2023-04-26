package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/x/amm/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EstimatePoolExit(
	goCtx context.Context,
	req *types.QueryEstimatePoolExitRequest,
) (
	*types.QueryEstimatePoolExitResponse,
	error,
) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	pool, found := k.GetPool(ctx, req.PoolId)

	if !found {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrInvalidRequest,
			"Pool not found",
		)
	}

	burnAmt := sdk.NewInt(int64(req.Amount))

	retuns, err := CalcShareExit(pool, burnAmt)

	if err != nil {
		return nil, sdkerrors.Wrapf(
			err,
			"Failed to calculate pool coin return",
		)
	}

	return &types.QueryEstimatePoolExitResponse{Coins: retuns}, nil
}

func (k Keeper) EstimatePoolJoin(
	goCtx context.Context,
	req *types.QueryEstimatePoolJoinRequest,
) (
	*types.QueryEstimatePoolJoinResponse,
	error,
) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	pool, found := k.GetPool(ctx, req.PoolId)

	if !found {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrInvalidRequest,
			"Pool not found",
		)
	}

	share, excess, err := CalcShareJoin(pool.PoolToken, pool.Coins, req.Liquidity)

	if err != nil {
		return nil, sdkerrors.Wrapf(
			err,
			"Failed to calculate pool coin return",
		)
	}
	
	return &types.QueryEstimatePoolJoinResponse{Share: sdk.NewCoin(pool.PoolToken.Denom, share), Excess: excess}, nil
}
