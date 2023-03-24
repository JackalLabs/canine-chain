package lp

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/amm/keeper"
	"github.com/jackalLabs/canine-chain/x/amm/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the lPool
	for _, elem := range genState.PoolList {
		k.SetPool(ctx, elem)
	}
	// Set all the lProviderRecord
	for _, elem := range genState.ProviderRecordList {
		k.SetProviderRecord(ctx, elem)
		k.AddProviderRef(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PoolList = k.GetAllPool(ctx)
	genesis.ProviderRecordList = k.GetAllProviderRecord(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
