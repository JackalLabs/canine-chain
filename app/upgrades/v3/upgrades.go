package v3

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/jackalLabs/canine-chain/app/upgrades"
	filetreemoduletypes "github.com/jackalLabs/canine-chain/x/filetree/types"
	oraclemoduletypes "github.com/jackalLabs/canine-chain/x/oracle/types"
	storagemoduletypes "github.com/jackalLabs/canine-chain/x/storage/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

var (
	_ upgrades.Upgrade = &Upgrade{}
)

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
	return "v3"
}

// Handler implements upgrades.Upgrade
func (u *Upgrade) Handler() upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {

		// Do nothing here as we don't have anything particular in this update
		return u.mm.RunMigrations(ctx, u.configurator, fromVM)
	}
}

// StoreUpgrades implements upgrades.Upgrade
func (u *Upgrade) StoreUpgrades() *storetypes.StoreUpgrades {
	return &storetypes.StoreUpgrades{
		Added: []string{
			storagemoduletypes.StoreKey,
			filetreemoduletypes.StoreKey,
			oraclemoduletypes.StoreKey,
		},
	}
}
