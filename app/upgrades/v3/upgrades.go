package v3

import (
	"strconv"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/jackalLabs/canine-chain/app/upgrades"
	storageKeeper "github.com/jackalLabs/canine-chain/x/storage/keeper"
	storagemoduletypes "github.com/jackalLabs/canine-chain/x/storage/types"
)

var _ upgrades.Upgrade = &Upgrade{}

// Upgrade represents the v4 upgrade
type Upgrade struct {
	mm           *module.Manager
	configurator module.Configurator
	keeper       storageKeeper.Keeper
}

// NewUpgrade returns a new Upgrade instance
func NewUpgrade(mm *module.Manager, configurator module.Configurator, ok storageKeeper.Keeper) *Upgrade {
	return &Upgrade{
		mm:           mm,
		configurator: configurator,
		keeper:       ok,
	}
}

// Name implements upgrades.Upgrade
func (u *Upgrade) Name() string {
	return "v3"
}

// Handler implements upgrades.Upgrade
func (u *Upgrade) Handler() upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		fromVM[storagemoduletypes.ModuleName] = 4

		deals := u.keeper.GetAllActiveDealsLegacy(ctx) // converting every active deal into a new active deal data type
		for _, deal := range deals {

			startBlock, _ := strconv.ParseInt(deal.Startblock, 10, 64)
			endBlock, _ := strconv.ParseInt(deal.Endblock, 10, 64)
			fileSize, _ := strconv.ParseInt(deal.Filesize, 10, 64)
			proofVerified, _ := strconv.ParseBool(deal.Proofverified)
			proofMissed, _ := strconv.ParseInt(deal.Proofsmissed, 10, 64)
			blockToProve, _ := strconv.ParseInt(deal.Blocktoprove, 10, 64)

			newDeal := storagemoduletypes.ActiveDealsV2{
				Cid:           deal.Cid,
				Signer:        deal.Signee,
				Provider:      deal.Provider,
				StartBlock:    startBlock,
				EndBlock:      endBlock,
				FileSize:      fileSize,
				ProofVerified: proofVerified,
				ProofsMissed:  proofMissed,
				BlockToProve:  blockToProve,
				Creator:       deal.Creator,
				Merkle:        deal.Merkle,
				Fid:           deal.Fid,
				DealVersion:   0,
			}
			u.keeper.SetActiveDeals(ctx, newDeal)
			u.keeper.RemoveActiveDealsLegacy(ctx, deal.Cid)
		}

		strays := u.keeper.GetAllStraysLegacy(ctx) // converting every active deal into a new active deal data type
		for _, stray := range strays {
			fileSize, _ := strconv.ParseInt(stray.Filesize, 10, 64)
			newStray := storagemoduletypes.StrayV2{
				Cid:      stray.Cid,
				Signer:   stray.Signee,
				End:      stray.End,
				FileSize: fileSize,
				DealType: 0,
				Merkle:   stray.Merkle,
				Fid:      stray.Fid,
			}
			u.keeper.SetStrays(ctx, newStray)
			u.keeper.RemoveStraysLegacy(ctx, stray.Cid)
		}

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
