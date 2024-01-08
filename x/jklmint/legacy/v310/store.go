package v310

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/jackalLabs/canine-chain/v3/x/jklmint/types"
)

func MigrateStore(ctx sdk.Context, paramsSubspace *paramstypes.Subspace) error {
	var params types.Params
	paramsSubspace.GetParamSet(ctx, &params)

	params.TokensPerBlock = params.TokensPerBlock * 1_000_000

	if err := params.Validate(); err != nil {
		return err
	}

	paramsSubspace.SetParamSet(ctx, &params)

	return nil
}
