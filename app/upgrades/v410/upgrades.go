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

const MerkleSize = `075fe59e0539e2250ab61becec3a9b3568daf1564f8624ca45b767d176af4ca8e42dde52d00cd8432f779ac1f55bf2fa02ac032b3da318ccab8af2c9dc722de4,25,jklf1hm2dmjeax9tukul3qfjz04dgyn830t4sselzs0jxwpflvywx36wsyc46q0
14e8b5fc1353fabdd46948a7084d8a11763c7638f0a9887b3c2f116ecbcd25a9471a7fbb7beb4e627929a09d36a8a4d434a9d45edcb4e8c8c6113b1c5537706f,35,jklf1hm2dmjeax9tukul3qfjz04dgyn830t4sselzs0jxwpflvywx36wsyc46q0
32867e4d607e3d70ec6abf9f483df5c2460d0b0bec8f6c78cffd271b5b93d585b65d6813c6db3e9372f08db5df8cb46e9dd92c18cbd174a9f89c421a4624d68e,85,jklf1hm2dmjeax9tukul3qfjz04dgyn830t4sselzs0jxwpflvywx36wsyc46q0
4415a96b12f193c216f69f90145af419e48b7b008c5186228c04b1421ae8441e80c28085e8379c256ec7be91a9644ee3e58980fb4cf0f960ef247fd1077de48a,29,jklf1hm2dmjeax9tukul3qfjz04dgyn830t4sselzs0jxwpflvywx36wsyc46q0`

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

	for _, line := range strings.Split(strings.TrimSuffix(MerkleSize, "\n"), "\n") {
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
			Start:         planHeight,
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
