package v4

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/jackalLabs/canine-chain/v3/app/upgrades"
	filetreemodulekeeper "github.com/jackalLabs/canine-chain/v3/x/filetree/keeper"
	storagekeeper "github.com/jackalLabs/canine-chain/v3/x/storage/keeper"

	storagemoduletypes "github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

var _ upgrades.Upgrade = &Upgrade{}

// Upgrade represents the v4 upgrade
type Upgrade struct {
	mm           *module.Manager
	configurator module.Configurator
	sk           storagekeeper.Keeper
	fk           filetreemodulekeeper.Keeper
}

// NewUpgrade returns a new Upgrade instance
func NewUpgrade(mm *module.Manager, configurator module.Configurator, sk storagekeeper.Keeper, fk filetreemodulekeeper.Keeper) *Upgrade {
	return &Upgrade{
		mm:           mm,
		configurator: configurator,
		sk:           sk,
		fk:           fk,
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

type FidContents struct {
	Fid []string `json:"fids"`
}

type MerkleContents struct {
	Merkles [][]byte `json:"merkles"`
}

func UpdateFileTree(ctx sdk.Context, fk filetreemodulekeeper.Keeper, merkleMap map[string][]byte) {
	allFiles := fk.GetAllFiles(ctx)

	for _, file := range allFiles {
		contents := file.Contents

		var fidContents FidContents
		err := json.Unmarshal([]byte(contents), &fidContents)
		if err != nil {
			ctx.Logger().Debug(fmt.Errorf("cannot unmarshal %s %w", file.Address, err).Error())
			continue
		}

		merkles := make([][]byte, 0)

		for _, fid := range fidContents.Fid {
			m := merkleMap[fid]
			if m == nil {
				continue
			}

			merkles = append(merkles, m)

		}

		merkleContents := MerkleContents{Merkles: merkles}

		merkleContentBytes, err := json.Marshal(merkleContents)
		if err != nil {
			ctx.Logger().Debug(fmt.Errorf("cannot marshal merkle contents of %s %w", file.Address, err).Error())
			continue
		}

		file.Contents = string(merkleContentBytes)
		fk.SetFiles(ctx, file)
	}
}

func UpdatePaymentInfo(ctx sdk.Context, sk storagekeeper.Keeper) {
	paymentInfo := sk.GetAllStoragePaymentInfo(ctx)
	for _, info := range paymentInfo {

		planTime := info.End.Sub(info.Start)
		millis := planTime.Milliseconds()
		seconds := millis / 1000
		minutes := seconds / 60
		hours := minutes / 60

		cost := sk.GetStorageCostKbs(ctx, info.SpaceAvailable, hours)

		price := sdk.NewCoin("ujkl", cost)

		info.Coins = sdk.NewCoins(price)

		sk.SetStoragePaymentInfo(ctx, info)
	}
}

func UpdateFiles(ctx sdk.Context, u *Upgrade) map[string][]byte {
	fidMerkle := make(map[string][]byte)

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

		fidMerkle[deal.Fid] = merkle // creating fid -> merkle mapping

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

	return fidMerkle
}

// Handler implements upgrades.Upgrade
func (u *Upgrade) Handler() upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		ctx.Logger().Info("\nNow updating the Jackal Protocol to:\n\n █████╗  ██████╗ █████╗  ██████╗██╗ █████╗ \n██╔══██╗██╔════╝██╔══██╗██╔════╝██║██╔══██╗\n███████║██║     ███████║██║     ██║███████║\n██╔══██║██║     ██╔══██║██║     ██║██╔══██║\n██║  ██║╚██████╗██║  ██║╚██████╗██║██║  ██║\n╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝ ╚═════╝╚═╝╚═╝  ╚═╝\n                                           \n")

		fromVM[storagemoduletypes.ModuleName] = 5

		fidMerkleMap := UpdateFiles(ctx, u)

		UpdateFileTree(ctx, u.fk, fidMerkleMap)

		UpdatePaymentInfo(ctx, u.sk) // updating payment info with values at time of upgrade

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
