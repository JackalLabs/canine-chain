package keeper

import (
	"github.com/jackalLabs/canine-chain/x/dsig/types"
)

var _ types.QueryServer = Keeper{}
