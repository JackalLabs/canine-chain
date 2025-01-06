package v440

import (
	_ "embed"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/jackalLabs/canine-chain/v4/app/upgrades"
	mintkeeper "github.com/jackalLabs/canine-chain/v4/x/jklmint/keeper"
	storageKeeper "github.com/jackalLabs/canine-chain/v4/x/storage/keeper"
)

var _ upgrades.Upgrade = &Upgrade{}

//go:embed upgrade_data
var UpgradeData string

type Upgrade struct {
	mm           *module.Manager
	configurator module.Configurator
	sk           *storageKeeper.Keeper
	mk           *mintkeeper.Keeper
}

// NewUpgrade returns a new Upgrade instance
func NewUpgrade(mm *module.Manager, configurator module.Configurator, sk *storageKeeper.Keeper, mk *mintkeeper.Keeper) *Upgrade {
	return &Upgrade{
		mm:           mm,
		configurator: configurator,
		sk:           sk,
		mk:           mk,
	}
}

// Name implements upgrades.Upgrade
func (u *Upgrade) Name() string {
	return "v440"
}

// Handler implements upgrades.Upgrade
func (u *Upgrade) Handler() upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		var newWindow int64 = 7200
		params := u.mk.GetParams(ctx)
		params.TokensPerBlock = 3_830_000
		u.mk.SetParams(ctx, params)

		storageParams := u.sk.GetParams(ctx)
		oldProofWindow := storageParams.ProofWindow

		storageParams.CheckWindow = 300
		storageParams.ProofWindow = newWindow
		u.sk.SetParams(ctx, storageParams)

		files := u.sk.GetAllFileByMerkle(ctx)
		for _, file := range files {
			if file.ProofInterval == oldProofWindow { // updating default files to the new window
				file.ProofInterval = newWindow
				u.sk.SetFile(ctx, file)
			}
		}
		err := upgrades.RecoverFiles(ctx, u.sk, UpgradeData, plan.Height, "v4.4.0")
		if err != nil {
			return nil, err
		}

		return fromVM, nil
	}
}

// StoreUpgrades implements upgrades.Upgrade
func (u *Upgrade) StoreUpgrades() *storetypes.StoreUpgrades {
	return &storetypes.StoreUpgrades{
		Added:   []string{},
		Deleted: []string{},
	}
}
