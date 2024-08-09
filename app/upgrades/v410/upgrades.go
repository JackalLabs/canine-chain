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

const MerkleSize = `41c80bee949ad978cf166c653ee693675edd248dd209c14626b7c5fdad22b1285bf077167da1b53d94babeadf15ac9e22241398155141b1e9da2937e323e3d00,10843,jklf1hxx0ne28l0269fhm9xf7hn4m4ned6w2cz0u4gs64g69fjxsgm73s2zckff
bdb8770e72398e596518016e1eb1769aa8fe47e27a63f8126c23b394fecbc65875a5d5d3a0c50b26297648993efbf3b6a61742f12b71429682ba71a4cf48e5d1,1463004,jklf1hxx3q4jx5r5a5pgy4thp4q0ucccsvvry9u92pwwr57jmvx4qtj7qz0cn58
b6a34513ecea5f997fe87983b9f891384f2475adec02e0d58d3edf8875d999d04e6cdbe4161c9875f78ad4d720c77ff7a04d5a39913ada97a76ba83d371c559e,336,jklf1zzzwxkhrs9dktra54z2q693cqqf5jel4kl5d4p5szwnxw6hsa27qqvehk2`

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
			Start:         planHeight - int64(i/10), // 10 files per block
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
