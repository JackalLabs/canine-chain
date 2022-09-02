package keeper

import (
	"github.com/jackal-dao/canine/x/filetree/types"
)

var _ types.QueryServer = Keeper{}
