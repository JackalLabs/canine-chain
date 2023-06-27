package keeper

import (
	"github.com/jackalLabs/canine-chain/v3/x/notifications/types"
)

var _ types.QueryServer = Keeper{}
