package beta7

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/jackalLabs/canine-chain/app/upgrades"
	"github.com/jackalLabs/canine-chain/types"
	notificationkeeper "github.com/jackalLabs/canine-chain/x/notifications/keeper"
	notificationtypes "github.com/jackalLabs/canine-chain/x/notifications/types"
)

var _ upgrades.Upgrade = &Upgrade{}

// Upgrade represents the v4 upgrade
type Upgrade struct {
	mm                 *module.Manager
	configurator       module.Configurator
	notificationKeeper notificationkeeper.Keeper
}

// NewUpgrade returns a new Upgrade instance
func NewUpgrade(mm *module.Manager, configurator module.Configurator, notificationKeeper notificationkeeper.Keeper) *Upgrade {
	return &Upgrade{
		mm:                 mm,
		configurator:       configurator,
		notificationKeeper: notificationKeeper,
	}
}

// Name implements upgrades.Upgrade
func (u *Upgrade) Name() string {
	return "beta7"
}

// Handler implements upgrades.Upgrade
func (u *Upgrade) Handler() upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		if types.IsTestnet(ctx.ChainID()) || ctx.ChainID() == "test" {

			fromVM[notificationtypes.ModuleName] = 1

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
	return &storetypes.StoreUpgrades{
		Added: []string{notificationtypes.StoreKey},
	}
}
