package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/lp/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EstimateSwapOut(
	goCtx context.Context,
	req *types.QueryEstimateSwapOutRequest,
) (*types.QueryEstimateSwapOutResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	pool, found := k.GetLPool(ctx, req.PoolName)

	if !found {
		return nil, types.ErrLiquidityPoolNotFound
	}

	poolCoins := sdk.NewCoins(pool.Coins...)

	depositCoins, err := sdk.ParseCoinsNormalized(req.InputCoin)

	if err != nil || depositCoins.IsAnyNegative() {
		return nil, status.Error(codes.InvalidArgument, "invalid coinIn")
	}

	AMM, _ := types.GetAMM(pool.AMM_Id)

	returnAmt, err := AMM.EstimateReturn(poolCoins, depositCoins)
	if err != nil {
		// TODO: return better msg
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	return &types.QueryEstimateSwapOutResponse{OutputCoin: returnAmt[0]}, nil
}
