package keeper

import (
	"context"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/jklmint/types"
)

func FloatToBigInt(val float64) *big.Int {
	bigval := new(big.Float)
	bigval.SetFloat64(val)

	coin := new(big.Float)
	coin.SetInt(big.NewInt(1000000000000000000))

	bigval.Mul(bigval, coin)

	result := new(big.Int)
	bigval.Int(result)

	return result
}

func (k Keeper) Inflation(c context.Context, _ *types.QueryInflationRequest) (*types.QueryInflationResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	coins := k.bankKeeper.GetSupply(ctx, k.GetParams(ctx).MintDenom)

	amt := coins.Amount.ToDec()
	famt, err := amt.Float64()
	if err != nil {
		return nil, types.ErrCannotParseFloat
	}

	var tokens float64 = 10

	if famt <= 0 {
		return nil, types.ErrZeroDivision
	}

	ratio := tokens / famt

	ratioDec := FloatToBigInt(ratio)

	ratioSDK := sdk.NewDecFromBigInt(ratioDec)

	return &types.QueryInflationResponse{Inflation: ratioSDK}, nil
}
