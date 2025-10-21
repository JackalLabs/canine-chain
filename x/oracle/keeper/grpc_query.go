package keeper

import (
	"github.com/jackalLabs/canine-chain/v5/x/oracle/types"
)

var _ types.QueryServer = Keeper{}
