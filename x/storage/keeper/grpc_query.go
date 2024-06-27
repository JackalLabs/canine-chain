package keeper

import (
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
)

var _ types.QueryServer = Keeper{}
