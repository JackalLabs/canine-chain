package v2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/jackalLabs/canine-chain/x/rns/types"
)

// MigrateStore performs in-place store migrations from v1 to v2
// The things done here are the following:
// 1. setting up the next reason id and report id keys for existing subspaces
// 2. setting up the module params
func MigrateStore(ctx sdk.Context, paramsSubspace paramstypes.Subspace) error {
	ctx.Logger().Error("MIGRATING RNS STORE V2!")
	// Set the module params
	params := types.NewParams()

	params.DepositAccount = "jkl1t35eusvx97953uk47r3z4ckwd2prkn3fay76r8"

	paramsSubspace.SetParamSet(ctx, &params)

	return nil
}
