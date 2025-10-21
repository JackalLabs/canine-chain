package v4

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/jackalLabs/canine-chain/v5/x/jklmint/exported"
	"github.com/jackalLabs/canine-chain/v5/x/jklmint/types"
)

// MigrateStore performs in-place store migrations from v1 to v2
// The things done here are the following:
// 1. setting up the next reason id and report id keys for existing subspaces
// 2. setting up the module params
func MigrateStore(ctx sdk.Context, legacy exported.Subspace, paramsSubspace *paramstypes.Subspace) error {
	ctx.Logger().Info("MIGRATING MINT STORE!")
	// Set the module params
	var currParams Params
	legacy.GetParamSet(ctx, &currParams)

	t := types.Params{
		MintDenom:             currParams.MintDenom,
		DevGrantsRatio:        8,
		StorageProviderRatio:  12,
		StakerRatio:           80,
		TokensPerBlock:        3442500,
		MintDecrease:          6,
		StorageStipendAddress: "jkl1scpjepy0vgdpku7dwev20yxmvcezakv24k9arn",
	}

	paramsSubspace.SetParamSet(ctx, &t)

	return nil
}
