package keeper

import (
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

var _ types.QueryServer = Keeper{}
