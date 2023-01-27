package killdeals

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/jackalLabs/canine-chain/app/upgrades"
	storagemodulekeeper "github.com/jackalLabs/canine-chain/x/storage/keeper"
)

var _ upgrades.Upgrade = &Upgrade{}

// Upgrade represents the v4 upgrade
type Upgrade struct {
	mm             *module.Manager
	configurator   module.Configurator
	storeageKeeper storagemodulekeeper.Keeper
}

// NewUpgrade returns a new Upgrade instance
func NewUpgrade(mm *module.Manager, configurator module.Configurator, storeageKeeper storagemodulekeeper.Keeper) *Upgrade {
	return &Upgrade{
		mm:             mm,
		configurator:   configurator,
		storeageKeeper: storeageKeeper,
	}
}

// Name implements upgrades.Upgrade
func (u *Upgrade) Name() string {
	return "killdeals"
}

// Handler implements upgrades.Upgrade
func (u *Upgrade) Handler() upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		newVM, err := u.mm.RunMigrations(ctx, u.configurator, fromVM)
		if err != nil {
			return newVM, err
		}

		deals := u.storeageKeeper.GetAllActiveDeals(ctx)
		for _, deal := range deals {
			u.storeageKeeper.RemoveActiveDeals(ctx, deal.Cid)
		}

		strays := u.storeageKeeper.GetAllStrays(ctx)
		for _, stray := range strays {
			u.storeageKeeper.RemoveStrays(ctx, stray.Cid)
		}

		return newVM, err
	}
}

// StoreUpgrades implements upgrades.Upgrade
func (u *Upgrade) StoreUpgrades() *storetypes.StoreUpgrades {
	return &storetypes.StoreUpgrades{}
}
