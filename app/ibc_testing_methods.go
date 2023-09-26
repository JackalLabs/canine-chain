package app

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	ibckeeper "github.com/cosmos/ibc-go/v4/modules/core/keeper"
)

/*
We implement these methods for JackalApp so that
it satisfies the ibctesting.TestingApp interface, which
allows us to test our custom middleware

Note: osmosis implements these methods inside /app/modules.go
*/
func (app *JackalApp) GetBaseApp() *baseapp.BaseApp {
	return app.BaseApp
}

func (app *JackalApp) GetIBCKeeper() *ibckeeper.Keeper {
	return app.ibcKeeper // This is a *ibckeeper.Keeper
}

func (app *JackalApp) GetScopedIBCKeeper() capabilitykeeper.ScopedKeeper {
	return app.scopedIBCKeeper
}

// Required for ibctesting
func (app *JackalApp) GetStakingKeeper() stakingkeeper.Keeper {
	return app.stakingKeeper // Dereferencing the pointer
}

func (app *JackalApp) GetTxConfig() client.TxConfig {
	return MakeEncodingConfig().TxConfig
}
