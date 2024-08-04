package v3

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v4/modules/apps/29-fee/types"
	intertxtypes "github.com/cosmos/interchain-accounts/x/inter-tx/types"
	"github.com/jackalLabs/canine-chain/v4/app/upgrades"
	storagekeeper "github.com/jackalLabs/canine-chain/v4/x/storage/keeper"

	storagemoduletypes "github.com/jackalLabs/canine-chain/v4/x/storage/types"
)

var _ upgrades.Upgrade = &Upgrade{}

// Upgrade represents the v3 upgrade
type Upgrade struct {
	mm           *module.Manager
	configurator module.Configurator
	sk           storagekeeper.Keeper
}

// NewUpgrade returns a new Upgrade instance
func NewUpgrade(mm *module.Manager, configurator module.Configurator, sk storagekeeper.Keeper) *Upgrade {
	return &Upgrade{
		mm:           mm,
		configurator: configurator,
		sk:           sk,
	}
}

// Name implements upgrades.Upgrade
func (u *Upgrade) Name() string {
	return "v3"
}

// Handler implements upgrades.Upgrade
func (u *Upgrade) Handler() upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		fromVM[storagemoduletypes.ModuleName] = 4

		newVM, err := u.mm.RunMigrations(ctx, u.configurator, fromVM)
		if err != nil {
			return newVM, err
		}

		return newVM, err
	}
}

// StoreUpgrades implements upgrades.Upgrade
func (u *Upgrade) StoreUpgrades() *storetypes.StoreUpgrades {
	return &storetypes.StoreUpgrades{
		Added: []string{
			ibcfeetypes.StoreKey,
		},
		Deleted: []string{
			intertxtypes.StoreKey,
		},
	}
}
