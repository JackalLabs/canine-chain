package keeper

import (
	"github.com/jackal-dao/canine/x/notifications/types"
)

var _ types.QueryServer = Keeper{}
