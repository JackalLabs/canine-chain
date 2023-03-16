package keeper

import (
	"github.com/jackalLabs/canine-chain/x/jklmint/types"
)

var _ types.QueryServer = Keeper{}
