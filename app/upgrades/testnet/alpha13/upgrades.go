package alpha13

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/jackalLabs/canine-chain/v3/app/upgrades"
	"github.com/jackalLabs/canine-chain/v3/types"
)

var _ upgrades.Upgrade = &Upgrade{}

// Upgrade represents the v4 upgrade
type Upgrade struct {
	mm           *module.Manager
	configurator module.Configurator
}

// NewUpgrade returns a new Upgrade instance
func NewUpgrade(mm *module.Manager, configurator module.Configurator) *Upgrade {
	return &Upgrade{
		mm:           mm,
		configurator: configurator,
	}
}

// Name implements upgrades.Upgrade
func (u *Upgrade) Name() string {
	return "alpha13"
}

// Handler implements upgrades.Upgrade
func (u *Upgrade) Handler() upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		if types.IsTestnet(ctx.ChainID()) {
			newVM, err := u.mm.RunMigrations(ctx, u.configurator, fromVM)
			if err != nil {
				return newVM, err
			}

			return newVM, err
		}

		return fromVM, nil
	}
}

// StoreUpgrades implements upgrades.Upgrade
func (u *Upgrade) StoreUpgrades() *storetypes.StoreUpgrades {
	return &storetypes.StoreUpgrades{}
}
