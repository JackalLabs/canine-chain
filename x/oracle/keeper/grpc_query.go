package keeper

import (
	"github.com/jackalLabs/canine-chain/v3/x/oracle/types"
)

var _ types.QueryServer = Keeper{}
