package keeper

import (
	"github.com/jackalLabs/canine-chain/v3/x/rns/types"
)

var _ types.QueryServer = Keeper{}
