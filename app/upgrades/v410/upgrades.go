package v410

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/jackalLabs/canine-chain/v4/app/upgrades"
	storageKeeper "github.com/jackalLabs/canine-chain/v4/x/storage/keeper"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
)

const AccountName = "V410_STORAGE_ESCROW"

var _ upgrades.Upgrade = &Upgrade{}

type Upgrade struct {
	mm           *module.Manager
	configurator module.Configurator
	sk           *storageKeeper.Keeper
}

// NewUpgrade returns a new Upgrade instance
func NewUpgrade(mm *module.Manager, configurator module.Configurator, sk *storageKeeper.Keeper) *Upgrade {
	return &Upgrade{
		mm:           mm,
		configurator: configurator,
		sk:           sk,
	}
}

// Name implements upgrades.Upgrade
func (u *Upgrade) Name() string {
	return "v410"
}

func RecoverFiles(ctx sdk.Context, keeper *storageKeeper.Keeper, planHeight int64) error {
	account, err := types.GetAccount(AccountName) // creating account to hold the files
	if err != nil {
		return err
	}

	for i, line := range strings.Split(strings.TrimSuffix(MerkleSize, "\n"), "\n") {
		ctx.Logger().Info(line)
		data := strings.Split(line, ",")
		if len(data) < 3 {
			continue
		}
		merkle := data[0]
		sizeString := data[1]
		fid := data[2]
		size, err := strconv.ParseInt(sizeString, 10, 64)
		if err != nil {
			ctx.Logger().Error("cannot decode %s, skipping...", merkle)
			return err
		}

		merkleBytes, err := hex.DecodeString(merkle)
		if err != nil {
			ctx.Logger().Error("cannot decode %s, skipping...", merkle)
			return err
		}
		f := types.UnifiedFile{
			Merkle:        merkleBytes,
			Owner:         account.String(),
			Start:         planHeight - int64(i/32), // 32 files per block
			Expires:       planHeight + ((200 * 365 * 24 * 60 * 60) / 6),
			FileSize:      size,
			ProofInterval: 3600,
			ProofType:     0,
			Proofs:        make([]string, 0),
			MaxProofs:     3,
			Note:          fmt.Sprintf("{\"memo\":\"Recovered during v4.1.0\", \"fid\":\"%s\"}", fid),
		}

		keeper.SetFile(ctx, f)

	}
	return nil
}

// Handler implements upgrades.Upgrade
func (u *Upgrade) Handler() upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		err := RecoverFiles(ctx, u.sk, plan.Height)
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
