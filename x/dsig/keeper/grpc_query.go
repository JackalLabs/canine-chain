package keeper

import (
	"github.com/jackal-dao/canine/x/dsig/types"
)

var _ types.QueryServer = Keeper{}
