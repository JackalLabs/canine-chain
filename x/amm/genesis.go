package lp

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/amm/keeper"
	"github.com/jackalLabs/canine-chain/x/amm/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the Pool
	for _, elem := range genState.PoolList {
		k.SetPool(ctx, elem)
	}
	// Set all the ProviderRecord
	for _, elem := range genState.ProviderRecordList {
		k.SetProviderRecord(ctx, elem)
		k.AddProviderRef(ctx, elem)
	}
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PoolList = k.GetAllPool(ctx)
	genesis.ProviderRecordList = k.GetAllProviderRecord(ctx)

	return genesis
}
