package dsig

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/dsig/keeper"
	"github.com/jackal-dao/canine/x/dsig/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the userUploads
	for _, elem := range genState.UserUploadsList {
		k.SetUserUploads(ctx, elem)
	}
	// Set all the form
	for _, elem := range genState.FormList {
		k.SetForm(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.UserUploadsList = k.GetAllUserUploads(ctx)
	genesis.FormList = k.GetAllForm(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
