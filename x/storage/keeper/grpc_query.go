package keeper

import (
	"github.com/jackalLabs/canine-chain/v5/x/storage/types"
)

var _ types.QueryServer = Keeper{}
