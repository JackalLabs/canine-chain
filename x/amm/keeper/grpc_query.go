package keeper

import (
	"github.com/jackalLabs/canine-chain/x/amm/types"
)

var _ types.QueryServer = Keeper{}
