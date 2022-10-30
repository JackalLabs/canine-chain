package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackal-dao/canine/x/lp/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EstimatePoolRemove(
	goCtx context.Context,
	req *types.QueryEstimatePoolRemoveRequest,
) (*types.QueryEstimatePoolRemoveResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	pool, found := k.GetLPool(ctx, req.PoolName)

	if !found {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrInvalidRequest,
			"Pool not found",
		)
	}

	burnAmt, ok := sdk.NewIntFromString(req.Amount)

	if !ok {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrInvalidRequest,
			"Failed to parse burn amount into sdk.Int",
		)
	}

	returns, err := CalculatePoolShareBurnReturn(pool, burnAmt)
	if err != nil {
		return nil, sdkerrors.Wrapf(
			err,
			"Failed to calculate pool coin return",
		)
	}

	return &types.QueryEstimatePoolRemoveResponse{Coins: returns}, nil
}
