package v3

import (
	"fmt"

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

		for _, moduleName := range []string{"storage"} {
			logger.Debug("adding module", moduleName, "to version map")
		}

		logger.Debug("running module migrations")

		logger.Debug(fmt.Sprintf("mm is nil = %v", mm == nil))

		return mm.RunMigrations(ctx, configurator, vm)
	}
}
