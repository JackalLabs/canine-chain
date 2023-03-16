package rns

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/rns/keeper"
	"github.com/jackalLabs/canine-chain/x/rns/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the whois
	for _, elem := range genState.WhoIsList {
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
	for _, elem := range genState.ForSaleList {
		k.SetForsale(ctx, elem)
	}
	// Set all the init
	for _, elem := range genState.InitList {
		k.SetInit(ctx, elem)
	}
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.WhoIsList = k.GetAllWhois(ctx)
	genesis.NamesList = k.GetAllNames(ctx)
	genesis.BidsList = k.GetAllBids(ctx)
	genesis.ForSaleList = k.GetAllForsale(ctx)
	genesis.InitList = k.GetAllInit(ctx)

	return genesis
}
