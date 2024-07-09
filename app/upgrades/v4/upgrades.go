package v4

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	jklminttypes "github.com/jackalLabs/canine-chain/v4/x/jklmint/types"

	notificationsmoduletypes "github.com/jackalLabs/canine-chain/v4/x/notifications/types"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/jackalLabs/canine-chain/v4/app/upgrades"
	filetreemodulekeeper "github.com/jackalLabs/canine-chain/v4/x/filetree/keeper"
	storagekeeper "github.com/jackalLabs/canine-chain/v4/x/storage/keeper"

	storagemoduletypes "github.com/jackalLabs/canine-chain/v4/x/storage/types"
)

var _ upgrades.Upgrade = &Upgrade{}

// Upgrade represents the v4 upgrade
type Upgrade struct {
	mm           *module.Manager
	configurator module.Configurator
	sk           *storagekeeper.Keeper
	fk           *filetreemodulekeeper.Keeper
	bk           storagemoduletypes.BankKeeper
}

// NewUpgrade returns a new Upgrade instance
func NewUpgrade(mm *module.Manager, configurator module.Configurator, sk *storagekeeper.Keeper, fk *filetreemodulekeeper.Keeper, bk storagemoduletypes.BankKeeper) *Upgrade {
	return &Upgrade{
		mm:           mm,
		configurator: configurator,
		sk:           sk,
		fk:           fk,
		bk:           bk,
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
	Merkles [][]byte `json:"legacyMerkles"`
}

func UpdateFileTree(ctx sdk.Context, fk *filetreemodulekeeper.Keeper, merkleMap map[string][]byte) {
	ctx.Logger().Info("Migrating filetree formatting...")

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

		/**
		{
			"fids": ["fid1", "fid2", "fid3"]
		}

		{
			"legacyMerkles": ["merkle1", "merkle2", "merkle3"]
		}
		*/

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

func UpdatePaymentInfo(ctx sdk.Context, sk *storagekeeper.Keeper) {
	ctx.Logger().Info("Updating all user payment information...")

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

func UpdateFiles(ctx sdk.Context, sk *storagekeeper.Keeper) map[string][]byte {
	ctx.Logger().Info("Updating all files to Universal Files...")
	fidMerkle := make(map[string][]byte)

	p := sk.GetParams(ctx)

	allDeals := sk.GetAllLegacyActiveDeals(ctx)

	ctx.Logger().Info(fmt.Sprintf("There are %d active deals being migrated", len(allDeals)))

	for _, deal := range allDeals {

		merkle, err := hex.DecodeString(deal.Merkle)
		if err != nil {
			ctx.Logger().Error(fmt.Sprintf("cannot parse merkle string: '%s' | %s", deal.Merkle, err.Error()))
			continue
		}

		start, err := strconv.ParseInt(deal.Startblock, 10, 64)
		if err != nil {
			ctx.Logger().Error(fmt.Sprintf("cannot parse start block | %s", err.Error()))
			continue
		}

		end, err := strconv.ParseInt(deal.Endblock, 10, 64)
		if err != nil {
			ctx.Logger().Error(fmt.Sprintf("cannot parse end block | %s", err.Error()))
			continue
		}

		size, err := strconv.ParseInt(deal.Filesize, 10, 64)
		if err != nil {
			ctx.Logger().Error(fmt.Sprintf("cannot parse file size | %s", err.Error()))
			continue
		}

		lm := LegacyMarker{
			Fid: deal.Fid,
			Cid: deal.Cid,
		}

		fidMerkle[deal.Fid] = merkle // creating fid -> merkle mapping

		lmBytes, err := json.Marshal(lm)
		if err != nil {
			ctx.Logger().Error(fmt.Sprintf("cannot marshal legacy marker | %s", err.Error()))
			continue
		}

		uf := storagemoduletypes.UnifiedFile{
			Merkle:        merkle,
			Owner:         deal.Signee,
			Start:         start,
			Expires:       end,
			FileSize:      size,
			ProofInterval: p.ProofWindow, // TODO: Decide on default window
			ProofType:     0,
			Proofs:        make([]string, 0),
			MaxProofs:     3,
			Note:          string(lmBytes),
		}
		sk.SetFile(ctx, uf)

		_, found := sk.GetFile(ctx, merkle, deal.Signee, start)
		if !found {
			ctx.Logger().Error("Failed to migrate file")
		}
		uf.AddProver(ctx, sk, deal.Provider)

	}

	return fidMerkle
}

// Handler implements upgrades.Upgrade
func (u *Upgrade) Handler() upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		ctx.Logger().Info("\nNow updating the Jackal Protocol to:\n\n █████╗  ██████╗ █████╗  ██████╗██╗ █████╗ \n██╔══██╗██╔════╝██╔══██╗██╔════╝██║██╔══██╗\n███████║██║     ███████║██║     ██║███████║\n██╔══██║██║     ██╔══██║██║     ██║██╔══██║\n██║  ██║╚██████╗██║  ██║╚██████╗██║██║  ██║\n╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝ ╚═════╝╚═╝╚═╝  ╚═╝\n                                           \n")

		fromVM[storagemoduletypes.ModuleName] = 5
		fromVM[jklminttypes.ModuleName] = 4

		newVM, err := u.mm.RunMigrations(ctx, u.configurator, fromVM)
		if err != nil {
			return newVM, err
		}

		fidMerkleMap := UpdateFiles(ctx, u.sk)

		UpdateFileTree(ctx, u.fk, fidMerkleMap)

		UpdatePaymentInfo(ctx, u.sk) // updating payment info with values at time of upgrade

		err = u.ProvisionGauges(ctx)
		if err != nil {
			return newVM, sdkerrors.Wrapf(err, "could not provision gauges")
		}

		return newVM, err
	}
}

// ProvisionGauges creates new gauges from the already existing tokens.
// All tokens in the deposit wallet will be drip fed out to the storage providers for storage they are already providing.
func (u *Upgrade) ProvisionGauges(ctx sdk.Context) error {
	params := u.sk.GetParams(ctx)
	depAcc := params.DepositAccount
	dep, err := sdk.AccAddressFromBech32(depAcc)
	if err != nil {
		return sdkerrors.Wrapf(err, "cannot parse deposit account from params")
	}

	tok, err := storagemoduletypes.GetTokenHolderAccount()
	if err != nil {
		return sdkerrors.Wrapf(err, "cannot get token holder account")
	}

	c := u.bk.GetBalance(ctx, dep, "ujkl")
	toMove := sdk.NewCoins(c)

	err = u.bk.SendCoinsFromAccountToModule(ctx, dep, storagemoduletypes.ModuleName, toMove) // send tokens to token account
	if err != nil {
		return sdkerrors.Wrapf(err, "cannot send tokens from deposit account")
	}
	err = u.bk.SendCoinsFromModuleToAccount(ctx, storagemoduletypes.ModuleName, tok, toMove)
	if err != nil {
		return sdkerrors.Wrapf(err, "cannot send tokens to token holder account")
	}

	total := sdk.NewDecFromInt(c.Amount)
	year := total.Mul(sdk.NewDec(15).QuoInt64(100)) // 15%
	yearSend := sdk.NewCoins(sdk.NewCoin("ujkl", year.TruncateInt()))

	rest := total.Mul(sdk.NewDec(100 - 15).QuoInt64(100)) // 85%
	restSend := sdk.NewCoins(sdk.NewCoin("ujkl", rest.TruncateInt()))

	u.sk.NewGauge(ctx, yearSend, ctx.BlockTime().AddDate(1, 0, 0))   // 15% dripped over the first year
	u.sk.NewGauge(ctx, restSend, ctx.BlockTime().AddDate(200, 0, 0)) // rest dripped over the next 200 years

	return nil
}

// StoreUpgrades implements upgrades.Upgrade
func (u *Upgrade) StoreUpgrades() *storetypes.StoreUpgrades {
	return &storetypes.StoreUpgrades{
		Added: []string{
			notificationsmoduletypes.ModuleName, // swapping to brand-new notification module completely
		},
		Deleted: []string{
			"notifications",
		},
	}
}
