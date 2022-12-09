package keeper

import (
	"github.com/jackalLabs/canine-chain/x/oracle/types"
)

var _ types.QueryServer = Keeper{}
