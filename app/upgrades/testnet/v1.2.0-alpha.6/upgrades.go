package v120alpha6

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/jackalLabs/canine-chain/v5/types"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		logger := ctx.Logger().With("upgrade", UpgradeName)

		if types.IsTestnet(ctx.ChainID()) {
			logger.Debug("Updating to 1.2.0-alpha.6")
		}

		if types.IsMainnet(ctx.ChainID()) {
			logger.Debug("Ignoring Infra & Storage Deals for mainnet")
		}

		return mm.RunMigrations(ctx, configurator, vm)
	}
}
