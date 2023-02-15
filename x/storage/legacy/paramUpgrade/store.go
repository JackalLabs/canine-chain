package paramUpgrade

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

// MigrateStore performs in-place store migrations from v3 to v4
// The things done here are the following:
// 1. setting up the next reason id and report id keys for existing subspaces
// 2. setting up the module params
func MigrateStore(ctx sdk.Context, paramsSubspace *paramstypes.Subspace) error {
	ctx.Logger().Error("MIGRATING STORAGE STORE!")
	// Set the module params

	var params types.Params

	paramsSubspace.GetParamSet(ctx, &params)

	params.MissesToBurn = 3

	params.MaxContractAgeInBlocks = 100

	params.ChunkSize = 10240

	params.PriceFeed = "jklprice"

	params.PricePerTbPerMonth = 8

	paramsSubspace.SetParamSet(ctx, &params)

	return nil
}
