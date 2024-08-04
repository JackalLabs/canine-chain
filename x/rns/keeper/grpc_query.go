package keeper

import (
	"github.com/jackalLabs/canine-chain/v4/x/rns/types"
)

var _ types.QueryServer = Keeper{}
