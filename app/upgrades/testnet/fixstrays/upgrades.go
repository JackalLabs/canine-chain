package fixstrays

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/jackalLabs/canine-chain/app/upgrades"
	"github.com/jackalLabs/canine-chain/types"
	storagemodulekeeper "github.com/jackalLabs/canine-chain/x/storage/keeper"
	storeagemoduletypes "github.com/jackalLabs/canine-chain/x/storage/types"
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
	return "fixstrays"
}

// Handler implements upgrades.Upgrade
func (u *Upgrade) Handler() upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		if types.IsTestnet(ctx.ChainID()) || ctx.ChainID() == "test" {

			fromVM[storeagemoduletypes.ModuleName] = 2

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

			payinfo := u.storeageKeeper.GetAllStoragePaymentInfo(ctx)
			for _, info := range payinfo {
				u.storeageKeeper.RemoveStoragePaymentInfo(ctx, info.Address)
			}

			fidcid := u.storeageKeeper.GetAllFidCid(ctx)
			for _, fc := range fidcid {
				u.storeageKeeper.RemoveFidCid(ctx, fc.Fid)
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
