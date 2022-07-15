package keeper

import (
	"github.com/jackal-dao/canine/x/storage/types"
)

var _ types.QueryServer = Keeper{}
