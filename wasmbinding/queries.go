package wasmbinding

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/wasmbinding/bindings"
	filetreekeeper "github.com/jackalLabs/canine-chain/v3/x/filetree/keeper"
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

// GetFiles is a query to get a Files struct
// address is the full merklePath() of the file
func (qp QueryPlugin) GetFiles(ctx sdk.Context, address string, ownerAddress string) (*bindings.FilesResponse, error) {
	file, found := qp.filetreeKeeper.GetFiles(ctx, address, ownerAddress)
	if !found {
		return nil, fmt.Errorf("failed to get file: %s", file) // file is blank or empty struct if not found?
	}

	return &bindings.FilesResponse{Files: file.String()}, nil
}
