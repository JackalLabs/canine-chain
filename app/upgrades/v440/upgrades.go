package v440

import (
	_ "embed"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/jackalLabs/canine-chain/v5/app/upgrades"
	mintkeeper "github.com/jackalLabs/canine-chain/v5/x/jklmint/keeper"
	storageKeeper "github.com/jackalLabs/canine-chain/v5/x/storage/keeper"
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

func BumpInterval(ctx sdk.Context, sk *storageKeeper.Keeper) {
	var newWindow int64 = 7200

	storageParams := sk.GetParams(ctx)
	var oldProofWindow int64 = 3600

	storageParams.CheckWindow = 300
	storageParams.ProofWindow = newWindow
	sk.SetParams(ctx, storageParams)

	files := sk.GetAllFileByMerkle(ctx)
	for _, file := range files {
		if file.ProofInterval == oldProofWindow { // updating default files to the new window
			file.ProofInterval = newWindow
			sk.SetFile(ctx, file)
		}
	}
}

// Handler implements upgrades.Upgrade
func (u *Upgrade) Handler() upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		params := u.mk.GetParams(ctx)
		params.TokensPerBlock = 3_830_000
		u.mk.SetParams(ctx, params)

		BumpInterval(ctx, u.sk)

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
