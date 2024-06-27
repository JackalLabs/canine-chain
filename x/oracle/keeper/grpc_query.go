package keeper

import (
	"github.com/jackalLabs/canine-chain/v4/x/oracle/types"
)

var _ types.QueryServer = Keeper{}
