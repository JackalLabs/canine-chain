package keeper

import (
	"github.com/jackalLabs/canine-chain/v4/x/notifications/types"
)

var _ types.QueryServer = Keeper{}
