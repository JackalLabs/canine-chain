package v4

import (
	"encoding/hex"
	"encoding/json"
	"strconv"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/jackalLabs/canine-chain/v3/app/upgrades"
	storagekeeper "github.com/jackalLabs/canine-chain/v3/x/storage/keeper"

	storagemoduletypes "github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

var _ upgrades.Upgrade = &Upgrade{}

// Upgrade represents the v4 upgrade
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
	return "v4"
}

type LegacyMarker struct {
	Fid string `json:"fid"`
	Cid string `json:"cid"`
}

// Handler implements upgrades.Upgrade
func (u *Upgrade) Handler() upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		ctx.Logger().Info("\nNow updating the Jackal Protocol to:\n\n █████╗  ██████╗ █████╗  ██████╗██╗ █████╗ \n██╔══██╗██╔════╝██╔══██╗██╔════╝██║██╔══██╗\n███████║██║     ███████║██║     ██║███████║\n██╔══██║██║     ██╔══██║██║     ██║██╔══██║\n██║  ██║╚██████╗██║  ██║╚██████╗██║██║  ██║\n╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝ ╚═════╝╚═╝╚═╝  ╚═╝\n                                           \n")

		fromVM[storagemoduletypes.ModuleName] = 5

		allDeals := u.sk.GetAllLegacyActiveDeals(ctx)

		for _, deal := range allDeals {

			merkle, err := hex.DecodeString(deal.Merkle)
			if err != nil {
				ctx.Logger().Error(err.Error())
				continue
			}

			start, err := strconv.ParseInt(deal.Startblock, 10, 64)
			if err != nil {
				ctx.Logger().Error(err.Error())
				continue
			}

			end, err := strconv.ParseInt(deal.Endblock, 10, 64)
			if err != nil {
				ctx.Logger().Error(err.Error())
				continue
			}

			size, err := strconv.ParseInt(deal.Filesize, 10, 64)
			if err != nil {
				ctx.Logger().Error(err.Error())
				continue
			}

			lm := LegacyMarker{
				Fid: deal.Fid,
				Cid: deal.Cid,
			}

			lmBytes, err := json.Marshal(lm)
			if err != nil {
				ctx.Logger().Error(err.Error())
				continue
			}

			var uf storagemoduletypes.UnifiedFile

			uf, found := u.sk.GetFile(ctx, merkle, deal.Signee, start)
			if !found {
				uf = storagemoduletypes.UnifiedFile{
					Merkle:        merkle,
					Owner:         deal.Signee,
					Start:         start,
					Expires:       end,
					FileSize:      size,
					ProofInterval: 1800, // TODO: Decide on default window
					ProofType:     0,
					Proofs:        make([]string, 0),
					MaxProofs:     3,
					Note:          string(lmBytes),
				}
			}

			u.sk.SetFile(ctx, uf)
			uf.AddProver(ctx, u.sk, deal.Provider)

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
	return &storetypes.StoreUpgrades{
		Added:   []string{},
		Deleted: []string{},
	}
}
