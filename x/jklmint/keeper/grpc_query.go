package keeper

import (
	"github.com/jackalLabs/canine-chain/v5/x/jklmint/types"
)

var _ types.QueryServer = Keeper{}
