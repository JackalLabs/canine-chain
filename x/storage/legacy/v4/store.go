package v4

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

type ParamGetterSetter interface {
	GetParams(ctx sdk.Context) (p types.Params)
	SetParams(ctx sdk.Context, params types.Params)
}

// MigrateStore performs in-place store migrations from v1 to v2
// The things done here are the following:
// 1. setting up the next reason id and report id keys for existing subspaces
// 2. setting up the module params
func MigrateStore(ctx sdk.Context, pgs ParamGetterSetter) error {
	ctx.Logger().Error("Migrating Storage Store to V4!")
	params := pgs.GetParams(ctx)

	ctx.Logger().Info(params.String())

	params.AttestFormSize = 5
	params.AttestMinToPass = 3

	pgs.SetParams(ctx, params)

	ctx.Logger().Info(params.String())

	ctx.Logger().Info("DONE MIGRATING!")

	return nil
}
