package keeper

import (
	"github.com/jackal-dao/canine/x/rns/types"
)

var _ types.QueryServer = Keeper{}
