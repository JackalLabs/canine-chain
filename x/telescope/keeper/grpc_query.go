package keeper

import (
	"github.com/jackal-dao/canine/x/telescope/types"
)

var _ types.QueryServer = Keeper{}
