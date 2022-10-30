package keeper

import (
	"github.com/jackalLabs/canine-chain/x/lp/types"
)

var _ types.QueryServer = Keeper{}
