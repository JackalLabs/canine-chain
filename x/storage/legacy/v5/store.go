package v5

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/jackalLabs/canine-chain/v3/x/storage/exported"
)

func MigrateStore(ctx sdk.Context, legacySubspace exported.Subspace, paramsSubspace *paramstypes.Subspace) error {
	_ = ctx
	_ = legacySubspace
	_ = paramsSubspace
	return nil
}
