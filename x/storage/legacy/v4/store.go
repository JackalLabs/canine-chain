package v4

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/jackalLabs/canine-chain/x/storage/exported"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func MigrateStore(ctx sdk.Context, legacySubspace exported.Subspace, paramsSubspace *paramstypes.Subspace) error {
	var currParams types.Params
	legacySubspace.GetParamSet(ctx, &currParams)

	if err := currParams.Validate(); err != nil {
		return err
	}

	paramsSubspace.SetParamSet(ctx, &currParams)

	return nil
}
