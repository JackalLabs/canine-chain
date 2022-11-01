package v2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		logger := ctx.Logger().With("upgrade", UpgradeName)

		for k := range mm.Modules {
			// removes the following modules
			if k == "storage" || k == "dsig" || k == "notifications" || k == "filetreekeeper" {
				logger.Debug("removing module", k, "from version map")
				delete(vm, k)
			}
		}

		logger.Debug("running module migrations")
		return mm.RunMigrations(ctx, configurator, vm)
	}
}
