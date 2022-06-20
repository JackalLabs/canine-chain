package keeper

import (
	"github.com/jackal-dao/canine/x/jklaccounts/types"
)

var _ types.QueryServer = Keeper{}
