package fixstrays

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

// MigrateStore performs in-place store migrations from v2 to v3
// The things done here are the following:
// 1. setting up the next reason id and report id keys for existing subspaces
// 2. setting up the module params
func MigrateStore(ctx sdk.Context, paramsSubspace *paramstypes.Subspace) error {
	ctx.Logger().Error("MIGRATING STORAGE STORE!")
	// Set the module params

	var params types.Params

	paramsSubspace.GetParamSet(ctx, &params)

	params.ProofWindow = 50

	paramsSubspace.SetParamSet(ctx, &params)

	return nil
}
