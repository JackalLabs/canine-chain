package keeper

import (
	"context"
	"fmt"

	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

func (k msgServer) SignContract(_ context.Context, _ *types.MsgSignContract) (*types.MsgSignContractResponse, error) {
	return nil, fmt.Errorf("disabled new files until v4")
}
