package telescope

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/telescope/keeper"
	"github.com/jackal-dao/canine/x/telescope/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the whois
	for _, elem := range genState.WhoisList {
		k.SetWhois(ctx, elem)
	}
	// Set all the names
	for _, elem := range genState.NamesList {
		k.SetNames(ctx, elem)
	}
	// Set all the bids
	for _, elem := range genState.BidsList {
		k.SetBids(ctx, elem)
	}
	// Set all the forsale
	for _, elem := range genState.ForsaleList {
		k.SetForsale(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.WhoisList = k.GetAllWhois(ctx)
	genesis.NamesList = k.GetAllNames(ctx)
	genesis.BidsList = k.GetAllBids(ctx)
	genesis.ForsaleList = k.GetAllForsale(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
