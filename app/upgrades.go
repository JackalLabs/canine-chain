package app

import (
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/jackalLabs/canine-chain/v5/app/upgrades"
	"github.com/jackalLabs/canine-chain/v5/app/upgrades/bouncybulldog"
	v121 "github.com/jackalLabs/canine-chain/v5/app/upgrades/testnet/121"
	"github.com/jackalLabs/canine-chain/v5/app/upgrades/testnet/alpha11"
	"github.com/jackalLabs/canine-chain/v5/app/upgrades/testnet/alpha13"
	"github.com/jackalLabs/canine-chain/v5/app/upgrades/testnet/beta6"
	"github.com/jackalLabs/canine-chain/v5/app/upgrades/testnet/beta7"
	v3 "github.com/jackalLabs/canine-chain/v5/app/upgrades/v3"
	v4 "github.com/jackalLabs/canine-chain/v5/app/upgrades/v4"
	v410 "github.com/jackalLabs/canine-chain/v5/app/upgrades/v410"
	"github.com/jackalLabs/canine-chain/v5/app/upgrades/v410beta"
	"github.com/jackalLabs/canine-chain/v5/app/upgrades/v410testnet"
	v420 "github.com/jackalLabs/canine-chain/v5/app/upgrades/v420"
	v430 "github.com/jackalLabs/canine-chain/v5/app/upgrades/v430"
	v440 "github.com/jackalLabs/canine-chain/v5/app/upgrades/v440"
	v450 "github.com/jackalLabs/canine-chain/v5/app/upgrades/v450"
	v460 "github.com/jackalLabs/canine-chain/v5/app/upgrades/v460"
	"github.com/jackalLabs/canine-chain/v5/app/upgrades/v4alpha1"
	"github.com/jackalLabs/canine-chain/v5/app/upgrades/v4alpha3"
	"github.com/jackalLabs/canine-chain/v5/app/upgrades/v4alpha5"
	v500 "github.com/jackalLabs/canine-chain/v5/app/upgrades/v500"
	"github.com/jackalLabs/canine-chain/v5/app/upgrades/v500testnet"
	v510 "github.com/jackalLabs/canine-chain/v5/app/upgrades/v510"
)

func (app *JackalApp) registerTestnetUpgradeHandlers() {
	app.registerUpgrade(alpha11.NewUpgrade(app.mm, app.configurator, app.OracleKeeper))
	app.registerUpgrade(alpha13.NewUpgrade(app.mm, app.configurator))
	app.registerUpgrade(beta6.NewUpgrade(app.mm, app.configurator, app.StorageKeeper))
	app.registerUpgrade(beta7.NewUpgrade(app.mm, app.configurator, app.NotificationsKeeper))
	app.registerUpgrade(v121.NewUpgrade(app.mm, app.configurator))

	app.registerUpgrade(v4alpha1.NewUpgrade(app.mm, app.configurator))
	app.registerUpgrade(v4alpha3.NewUpgrade(app.mm, app.configurator))
	app.registerUpgrade(v4alpha5.NewUpgrade(app.mm, app.configurator))
	app.registerUpgrade(v410beta.NewUpgrade(app.mm, app.configurator))
	app.registerUpgrade(v410testnet.NewUpgrade(app.mm, app.configurator))
	app.registerUpgrade(v500testnet.NewUpgrade(app.mm, app.configurator))
}

func (app *JackalApp) registerMainnetUpgradeHandlers() {
	app.registerUpgrade(bouncybulldog.NewUpgrade(app.mm, app.configurator, app.OracleKeeper))
	app.registerUpgrade(v3.NewUpgrade(app.mm, app.configurator, app.StorageKeeper))
	app.registerUpgrade(v4.NewUpgrade(app.mm, app.configurator, &app.StorageKeeper, &app.FileTreeKeeper, app.BankKeeper))
	app.registerUpgrade(v410.NewUpgrade(app.mm, app.configurator, &app.StorageKeeper))
	app.registerUpgrade(v420.NewUpgrade(app.mm, app.configurator))
	app.registerUpgrade(v430.NewUpgrade(app.mm, app.configurator, &app.StorageKeeper))
	app.registerUpgrade(v440.NewUpgrade(app.mm, app.configurator, &app.StorageKeeper, &app.MintKeeper))
	app.registerUpgrade(v450.NewUpgrade(app.mm, app.configurator, &app.StorageKeeper, app.BankKeeper, app.AccountKeeper))
	app.registerUpgrade(v460.NewUpgrade(app.mm, app.configurator))
	app.registerUpgrade(v500.NewUpgrade(app.mm, app.configurator))
	app.registerUpgrade(v510.NewUpgrade(app.mm, app.configurator))
}

// registerUpgrade registers the given upgrade to be supported by the app
func (app *JackalApp) registerUpgrade(upgrade upgrades.Upgrade) {
	app.upgradeKeeper.SetUpgradeHandler(upgrade.Name(), upgrade.Handler())

	upgradeInfo, err := app.upgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(err)
	}

	if upgradeInfo.Name == upgrade.Name() && !app.upgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		// Configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, upgrade.StoreUpgrades()))
	}
}
