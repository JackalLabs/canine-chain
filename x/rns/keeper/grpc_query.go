package keeper

import (
	"github.com/jackalLabs/canine-chain/x/rns/types"
)

var _ types.QueryServer = Keeper{}
