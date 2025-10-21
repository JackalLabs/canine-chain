package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v5/x/jklmint/types"
)

func (k Keeper) Inflation(c context.Context, _ *types.QueryInflation) (*types.QueryInflationResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	inflation, err := k.GetInflation(ctx)

	return &types.QueryInflationResponse{Inflation: inflation}, err
}
