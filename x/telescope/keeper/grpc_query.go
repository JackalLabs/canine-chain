package keeper

import (
	"github.com/themarstonconnell/telescope/x/telescope/types"
)

var _ types.QueryServer = Keeper{}
