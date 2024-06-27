package oracle

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/x/oracle/keeper"
	"github.com/jackalLabs/canine-chain/v4/x/oracle/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)

	// Set all the feeds
	for _, elem := range genState.FeedList {
		k.SetFeed(ctx, elem)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.FeedList = k.GetAllFeeds(ctx)

	return genesis
}
