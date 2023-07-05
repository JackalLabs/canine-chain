package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/jklmint/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) (p types.Params) {
	k.paramSpace.GetParamSet(ctx, &p)
	return
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// MintDenom returns the MintDenom param
func (k Keeper) MintDenom(ctx sdk.Context) (res string) {
	k.paramSpace.Get(ctx, types.KeyMintDenom, &res)
	return
}
