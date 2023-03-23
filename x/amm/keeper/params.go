package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/lp/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramstore.GetParamSet(ctx, &params)
	return params
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// Returns account address that protocol fees are collected to
// Panics if the address couldn't be converted to sdk.AccAddress
func (k Keeper) ProtocolFeeToAcc(ctx sdk.Context) sdk.AccAddress {
	params := k.GetParams(ctx)
	return sdk.MustAccAddressFromBech32(params.ProtocolFeeToAddr)
}

// Returns protocol fee rate
// Panics if the rate couldn't be converted to sdk.Dec
func (k Keeper) ProtocolFeeRate(ctx sdk.Context) sdk.Dec {
	params := k.GetParams(ctx)
	return sdk.MustNewDecFromStr(params.ProtocolFeeRate)
}
