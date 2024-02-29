package v320

import (
	"fmt"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/jackalLabs/canine-chain/v3/app/upgrades"
	storagekeeper "github.com/jackalLabs/canine-chain/v3/x/storage/keeper"
	storagemoduletypes "github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

var _ upgrades.Upgrade = &Upgrade{}

// Upgrade represents the v3 upgrade
type Upgrade struct {
	mm           *module.Manager
	configurator module.Configurator
	keeper       storagekeeper.Keeper
}

// NewUpgrade returns a new Upgrade instance
func NewUpgrade(mm *module.Manager, configurator module.Configurator, keeper storagekeeper.Keeper) *Upgrade {
	return &Upgrade{
		mm:           mm,
		configurator: configurator,
		keeper:       keeper,
	}
}

// Name implements upgrades.Upgrade
func (u *Upgrade) Name() string {
	return "v320"
}

// Handler implements upgrades.Upgrade
func (u *Upgrade) Handler() upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		fromVM[storagemoduletypes.ModuleName] = 5

		ctx.Logger().Info("Upgrading the Jackal Network to:\n██╗   ██╗██████╗    ██████╗     ██████╗ \n██║   ██║╚════██╗   ╚════██╗   ██╔═████╗\n██║   ██║ █████╔╝    █████╔╝   ██║██╔██║\n╚██╗ ██╔╝ ╚═══██╗   ██╔═══╝    ████╔╝██║\n ╚████╔╝ ██████╔╝██╗███████╗██╗╚██████╔╝\n  ╚═══╝  ╚═════╝ ╚═╝╚══════╝╚═╝ ╚═════╝ ")

		u.keeper.IterateLegacyActiveDeals(ctx, func(deal storagemoduletypes.LegacyActiveDeals) bool {
			ctx.Logger().Info(fmt.Sprintf("%s being migrated", deal.Cid))

			u.keeper.SetActiveDeals(ctx, storagemoduletypes.ActiveDeals{
				Cid:          deal.Cid,
				Signee:       deal.Signee,
				Provider:     deal.Provider,
				Startblock:   deal.Startblock,
				Endblock:     deal.Endblock,
				Filesize:     deal.Filesize,
				LastProof:    ctx.BlockHeight(),
				Proofsmissed: deal.Proofsmissed,
				Blocktoprove: deal.Blocktoprove,
				Creator:      deal.Creator,
				Merkle:       deal.Merkle,
				Fid:          deal.Fid,
			})

			return false
		})

		newVM, err := u.mm.RunMigrations(ctx, u.configurator, fromVM)
		if err != nil {
			return newVM, err
		}

		return newVM, err
	}
}

// StoreUpgrades implements upgrades.Upgrade
func (u *Upgrade) StoreUpgrades() *storetypes.StoreUpgrades {
	return &storetypes.StoreUpgrades{}
}
