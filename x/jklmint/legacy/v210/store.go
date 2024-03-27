package v210

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	v3 "github.com/jackalLabs/canine-chain/v3/x/jklmint/legacy/v3"
)

// MigrateStore performs in-place store migrations from v1 to v2
// The things done here are the following:
// 1. setting up the next reason id and report id keys for existing subspaces
// 2. setting up the module params
func MigrateStore(ctx sdk.Context, paramsSubspace *paramstypes.Subspace) error {
	ctx.Logger().Error("MIGRATING MINT STORE!")
	// Set the module params
	params := v3.DefaultParams()

	params.ProviderRatio = 0

	paramsSubspace.SetParamSet(ctx, &params)

	return nil
}
