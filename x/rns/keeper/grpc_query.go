package keeper

import (
	"github.com/jackalLabs/canine-chain/v5/x/rns/types"
)

var _ types.QueryServer = Keeper{}
