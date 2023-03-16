package bouncybulldog

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/jackalLabs/canine-chain/app/upgrades"
	"github.com/jackalLabs/canine-chain/types"
	filetreemoduletypes "github.com/jackalLabs/canine-chain/x/filetree/types"
	notificationsmoduletypes "github.com/jackalLabs/canine-chain/x/notifications/types"
	oraclekeeper "github.com/jackalLabs/canine-chain/x/oracle/keeper"
	oraclemoduletypes "github.com/jackalLabs/canine-chain/x/oracle/types"
	rnsmoduletypes "github.com/jackalLabs/canine-chain/x/rns/types"
	storagemoduletypes "github.com/jackalLabs/canine-chain/x/storage/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

var _ upgrades.Upgrade = &Upgrade{}

// Upgrade represents the v4 upgrade
type Upgrade struct {
	mm           *module.Manager
	configurator module.Configurator
	ok           oraclekeeper.Keeper
}

// NewUpgrade returns a new Upgrade instance
func NewUpgrade(mm *module.Manager, configurator module.Configurator, ok oraclekeeper.Keeper) *Upgrade {
	return &Upgrade{
		mm:           mm,
		configurator: configurator,
		ok:           ok,
	}
}

// Name implements upgrades.Upgrade
func (u *Upgrade) Name() string {
	return "bouncybulldog"
}

// Handler implements upgrades.Upgrade
func (u *Upgrade) Handler() upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		if types.IsTestnet(ctx.ChainID()) {
			ctx.Logger().Error("Upgrade shouldn't run on testnet!")
			return fromVM, nil
		}

		fromVM[storagemoduletypes.ModuleName] = 3
		fromVM[filetreemoduletypes.ModuleName] = 1
		fromVM[oraclemoduletypes.ModuleName] = 1
		fromVM[notificationsmoduletypes.ModuleName] = 1
		fromVM[rnsmoduletypes.ModuleName] = 2

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
			storagemoduletypes.StoreKey,
			filetreemoduletypes.StoreKey,
			oraclemoduletypes.StoreKey,
			notificationsmoduletypes.StoreKey,
		},
	}
}
