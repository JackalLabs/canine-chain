package keeper

import (
	"github.com/jackalLabs/canine-chain/v3/x/jklmint/types"
)

var _ types.QueryServer = Keeper{}
