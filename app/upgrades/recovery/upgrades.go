package recovery

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/jackalLabs/canine-chain/app/upgrades"
	storagekeeper "github.com/jackalLabs/canine-chain/x/storage/keeper"
)

var _ upgrades.Upgrade = &Upgrade{}

// Upgrade represents the v4 upgrade
type Upgrade struct {
	mm            *module.Manager
	configurator  module.Configurator
	storageKeeper storagekeeper.Keeper
}

// NewUpgrade returns a new Upgrade instance
func NewUpgrade(mm *module.Manager, configurator module.Configurator, storageKeeper storagekeeper.Keeper) *Upgrade {
	return &Upgrade{
		mm:            mm,
		configurator:  configurator,
		storageKeeper: storageKeeper,
	}
}

// Name implements upgrades.Upgrade
func (u *Upgrade) Name() string {
	return "recovery"
}

// Handler implements upgrades.Upgrade
func (u *Upgrade) Handler() upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		strays := u.storageKeeper.GetAllStrays(ctx)
		deals := u.storageKeeper.GetAllActiveDeals(ctx)

		for _, stray := range strays {
			found := false
			for _, deal := range deals {
				if stray.Fid == deal.Fid {
					found = true
					break
				}
			}
			if found {
				continue
			}

			u.storageKeeper.RemoveStrays(ctx, stray.Cid)
		}

		return fromVM, nil
	}
}

// StoreUpgrades implements upgrades.Upgrade
func (u *Upgrade) StoreUpgrades() *storetypes.StoreUpgrades {
	return &storetypes.StoreUpgrades{}
}
