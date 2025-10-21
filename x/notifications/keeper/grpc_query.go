package keeper

import (
	"github.com/jackalLabs/canine-chain/v5/x/notifications/types"
)

var _ types.QueryServer = Keeper{}
