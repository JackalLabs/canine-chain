package wasmbinding

import (
	filetreekeeper "github.com/jackalLabs/canine-chain/x/filetree/keeper"
)

type QueryPlugin struct {
	filetreeKeeper *filetreekeeper.Keeper
}

// NewQueryPlugin returns a reference to a new QueryPlugin.
func NewQueryPlugin(ftk *filetreekeeper.Keeper) *QueryPlugin {
	return &QueryPlugin{
		filetreeKeeper: ftk,
	}
}

// !!! MUST DO !!!
// Implement a query
