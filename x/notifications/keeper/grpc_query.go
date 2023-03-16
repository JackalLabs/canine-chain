package keeper

import (
	"github.com/jackalLabs/canine-chain/x/notifications/types"
)

var _ types.QueryServer = Keeper{}
